package ladder

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/ladder"
	ladderReq "github.com/oldweipro/gin-admin/model/ladder/request"
	systemReq "github.com/oldweipro/gin-admin/model/system/request"
	systemService "github.com/oldweipro/gin-admin/service/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type InboundsService struct {
}

var userService systemService.UserService

// CreateInbounds 创建Inbounds记录
func (inboundsService *InboundsService) CreateInbounds(inbounds *ladder.Inbounds) (err error) {
	err = global.DB.Create(inbounds).Error
	return err
}

// DeleteInbounds 删除Inbounds记录
func (inboundsService *InboundsService) DeleteInbounds(inbounds ladder.Inbounds) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ladder.Inbounds{}).Where("id = ?", inbounds.ID).Update("deleted_by", inbounds.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&inbounds).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteInboundsByIds 批量删除Inbounds记录
func (inboundsService *InboundsService) DeleteInboundsByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ladder.Inbounds{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&ladder.Inbounds{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateInbounds 更新Inbounds记录
func (inboundsService *InboundsService) UpdateInbounds(inbounds ladder.Inbounds) (err error) {
	err = global.DB.Save(&inbounds).Error
	return err
}

// GetInbounds 根据id获取Inbounds记录
func (inboundsService *InboundsService) GetInbounds(id uint) (inbounds ladder.Inbounds, err error) {
	err = global.DB.Where("id = ?", id).First(&inbounds).Error
	return
}

func (inboundsService *InboundsService) GetInboundsLink(userInfo systemReq.CustomClaims, sid uint) (inbounds ladder.Inbounds, err error) {
	if err = global.DB.Where("uid = ? and sid = ?", userInfo.BaseClaims.ID, sid).First(&inbounds).Error; err != nil {
		// 如果入站链接不存在，则创建链接
		inbounds.Sid = &sid
		if err = inboundsService.CreateServerNodeInboundsLink(userInfo, &inbounds); err != nil {
			global.Logger.Error("远程服务器创建链接失败: ", zap.Error(err))
		} else {
			err = inboundsService.CreateInbounds(&inbounds)
		}
	}
	return
}

func (inboundsService *InboundsService) SetInboundsLink(userInfo systemReq.CustomClaims, inbounds ladder.Inbounds) (err error) {
	if err = global.DB.Where("uid = ? and sid = ?", userInfo.BaseClaims.ID, *inbounds.Sid).First(&inbounds).Error; err != nil {
		// 如果入站链接不存在，则创建链接，防止还没生成链接有人就点重置按钮导致数据库查不出来而报错
		if err := inboundsService.CreateServerNodeInboundsLink(userInfo, &inbounds); err != nil {
			global.Logger.Error("远程服务器创建链接失败: ", zap.Error(err))
		} else {
			err = inboundsService.CreateInbounds(&inbounds)
		}
	} else {
		if err := inboundsService.CreateServerNodeInboundsLink(userInfo, &inbounds); err != nil {
			global.Logger.Error("远程服务器创建链接失败: ", zap.Error(err))
		} else {
			err = inboundsService.UpdateInbounds(inbounds)
		}
	}
	return
}

// CreateServerNodeInboundsLink 向节点服务器添加节点链接
func (inboundsService *InboundsService) CreateServerNodeInboundsLink(userInfo systemReq.CustomClaims, inbounds *ladder.Inbounds) (err error) {
	// 查询服务器信息
	var serverNode ladder.ServerNode
	global.DB.Where("id = ?", *inbounds.Sid).First(&serverNode)
	user, err := userService.FindUserById(userInfo.BaseClaims.ID)
	if err != nil {
		return
	}
	var up int64 = 0
	var down int64 = 0
	var total int64 = 0
	enable := true
	expiryTime := user.LadderExpire
	// 判断一下，如果有数据就不重置了
	if inbounds.Up == nil {
		inbounds.Up = &up
	}
	if inbounds.Down == nil {
		inbounds.Down = &down
	}
	inbounds.Total = &total
	inbounds.Remark = userInfo.NickName
	inbounds.Enable = &enable
	inbounds.ExpiryTime = &expiryTime
	inbounds.Protocol = "vmess"
	inbounds.Uid = &userInfo.BaseClaims.ID
	inbounds.ClientId = uuid.NewString()

	// 👇组装请求参数
	queryParams := make(map[string]string)
	queryParams["id"] = strconv.Itoa(int(*inbounds.Uid))
	queryParams["up"] = strconv.FormatInt(*inbounds.Up, 10)
	queryParams["down"] = strconv.FormatInt(*inbounds.Down, 10)
	queryParams["total"] = strconv.FormatInt(*inbounds.Total, 10)
	queryParams["remark"] = inbounds.Remark
	queryParams["enable"] = "true"
	queryParams["expiryTime"] = strconv.FormatInt(*inbounds.ExpiryTime, 10)
	queryParams["listen"] = ""
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(40000) + 20000
	inbounds.Port = strconv.Itoa(r)
	queryParams["port"] = inbounds.Port
	queryParams["protocol"] = inbounds.Protocol
	settings := ladder.Settings{
		DisableInsecureEncryption: false,
		Clients: []ladder.Clients{
			{
				Id:      inbounds.ClientId,
				AlterId: 0,
			},
		},
	}
	settingsJson, _ := json.MarshalIndent(settings, "", "  ")
	inbounds.Settings = string(settingsJson)
	queryParams["settings"] = inbounds.Settings
	streamSettings := ladder.StreamSettings{
		Network:  "tcp",
		Security: "tls",
		TlsSettings: ladder.TlsSettings{
			ServerName: serverNode.Domain,
			Certificates: []ladder.Certificates{
				{
					// TODO 暂且先都是以文件准，应该兼容以字符串的方式
					CertificateFile: serverNode.PemFile,
					KeyFile:         serverNode.KeyFile,
				},
			},
		},
		TcpSettings: ladder.TcpSettings{
			Header: ladder.Header{
				Type: "none",
			},
		},
	}
	streamSettingssJson, _ := json.MarshalIndent(streamSettings, "", "  ")
	inbounds.StreamSettings = string(streamSettingssJson)
	queryParams["streamSettings"] = inbounds.StreamSettings
	sniffing := ladder.Sniffing{
		Enabled: true,
		DestOverride: []string{
			"http",
			"tls",
		},
	}
	sniffingJson, _ := json.MarshalIndent(sniffing, "", "  ")
	inbounds.Sniffing = string(sniffingJson)
	queryParams["sniffing"] = inbounds.Sniffing
	// 👇发起请求
	reqUrl := "http://" + serverNode.ServerHost + ":" + strconv.Itoa(*serverNode.ServerPort) + "/xui/inbound/add"
	cookie := &http.Cookie{
		Name:  "session",
		Value: serverNode.Cookie,
	}
	client := resty.New()
	resp, err := client.R().
		SetCookies([]*http.Cookie{
			cookie,
		}).
		SetFormData(queryParams).
		Post(reqUrl)
	if err != nil {
		fmt.Println(err)
	}
	// 这里判断一下resp返回的内容
	//fmt.Println(resp.StatusCode())
	//fmt.Println(resp)
	if resp.StatusCode() != 200 {
		err = errors.New("请求节点服务器错误")
	}
	// TODO 这个vmess链接应该是动态生成的
	vMessLink := make(map[string]interface{})
	vMessLink["v"] = "2"
	vMessLink["ps"] = serverNode.Region
	vMessLink["add"] = serverNode.Domain
	vMessLink["port"] = inbounds.Port
	vMessLink["id"] = inbounds.ClientId
	vMessLink["aid"] = 0
	vMessLink["net"] = "tcp"
	vMessLink["type"] = "none"
	vMessLink["host"] = ""
	vMessLink["path"] = ""
	vMessLink["tls"] = "tls"
	vMessLinkJson, _ := json.MarshalIndent(vMessLink, "", "  ")
	inbounds.Link = string(vMessLinkJson)
	vMessLinkJsonBase64 := base64.StdEncoding.EncodeToString(vMessLinkJson)
	inbounds.Link64 = "vmess://" + vMessLinkJsonBase64
	return
}

// GetInboundsInfoList 分页获取Inbounds记录
func (inboundsService *InboundsService) GetInboundsInfoList(info ladderReq.InboundsSearch) (list []ladder.Inbounds, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&ladder.Inbounds{})
	var inboundss []ladder.Inbounds
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&inboundss).Error
	return inboundss, total, err
}
