package ladder

import (
	"github.com/go-resty/resty/v2"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/ladder"
	ladderReq "github.com/oldweipro/gin-admin/model/ladder/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type ServerNodeService struct {
}

var inboundsService InboundsService

// CreateServerNode 创建ServerNode记录
func (serverNodeService *ServerNodeService) CreateServerNode(serverNode *ladder.ServerNode) (err error) {
	err = global.DB.Create(serverNode).Error
	return err
}

// DeleteServerNode 删除ServerNode记录
func (serverNodeService *ServerNodeService) DeleteServerNode(serverNode ladder.ServerNode) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ladder.ServerNode{}).Where("id = ?", serverNode.ID).Update("deleted_by", serverNode.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&serverNode).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteServerNodeByIds 批量删除ServerNode记录
func (serverNodeService *ServerNodeService) DeleteServerNodeByIds(ids request.IdsReq, deletedBy uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ladder.ServerNode{}).Where("id in ?", ids.Ids).Update("deleted_by", deletedBy).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&ladder.ServerNode{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateServerNode 更新ServerNode记录
func (serverNodeService *ServerNodeService) UpdateServerNode(serverNode ladder.ServerNode) (err error) {
	err = global.DB.Save(&serverNode).Error
	return err
}

// GetServerNode 根据id获取ServerNode记录
func (serverNodeService *ServerNodeService) GetServerNode(id uint) (serverNode ladder.ServerNode, err error) {
	err = global.DB.Where("id = ?", id).First(&serverNode).Error
	return
}

// GetServerNodeList 获取ServerNode所有记录
func (serverNodeService *ServerNodeService) GetServerNodeList() (list []ladder.ServerNode, err error) {
	// 创建db
	db := global.DB.Model(&ladder.ServerNode{})
	var serverNodes []ladder.ServerNode
	err = db.Find(&serverNodes).Error
	return serverNodes, err
}

// GetServerNodeInfoList 分页获取ServerNode记录
func (serverNodeService *ServerNodeService) GetServerNodeInfoList(info ladderReq.ServerNodeSearch) (list []ladder.ServerNode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&ladder.ServerNode{})
	var serverNodes []ladder.ServerNode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&serverNodes).Error
	return serverNodes, total, err
}

func (serverNodeService *ServerNodeService) GetUserServerNodeList() (serverNodes []ladder.ServerNode, err error) {
	err = global.DB.Where("server_status = 1").Find(&serverNodes).Error
	return
}

// GetServerNodeLessInfoList 分页获取ServerNode记录，过滤掉敏感信息
func (serverNodeService *ServerNodeService) GetServerNodeLessInfoList(info ladderReq.ServerNodeSearch, userID uint) (list []ladderReq.ServerNodeResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&ladder.ServerNode{})
	var serverNodes []ladder.ServerNode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	db = db.Where("server_status = 1")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Select("id", "bandwidth", "region").Limit(limit).Offset(offset).Find(&serverNodes).Error
	if err != nil {
		return nil, 0, err
	}
	var inbounds []ladder.Inbounds
	err = global.DB.Where("uid = ?", userID).Find(&inbounds).Error
	if err != nil {
		return nil, 0, err
	}
	for _, node := range serverNodes {
		var info ladderReq.ServerNodeResponse
		info.Id = node.ID
		info.Bandwidth = node.Bandwidth
		info.Region = node.Region
		// 过滤数据
		for _, inbound := range inbounds {
			if *inbound.Sid == node.ID {
				obj, err := inboundsService.GetInboundsInfo(*inbound.Uid, *inbound.Sid)
				if err != nil {
					break
				}
				info.ExpiryTime = obj.ExpiryTime
				info.Up = obj.Up
				info.Down = obj.Down
				break
			}
		}
		list = append(list, info)
	}
	return list, total, err
}

// ServerNodeLogin 处理登陆获取cookie
func (serverNodeService *ServerNodeService) ServerNodeLogin(serverNode ladder.ServerNode) (err error) {
	loginFormData := make(map[string]string)
	loginFormData["username"] = serverNode.Username
	loginFormData["password"] = serverNode.Password
	client := resty.New()
	resp, err := client.R().
		SetFormData(loginFormData).
		Post("http://" + serverNode.ServerHost + ":" + strconv.Itoa(*serverNode.ServerPort) + "/login")
	if err != nil {
		global.Logger.Error(serverNode.ServerHost+"登录梯子失败:", zap.Error(err))
		return
	}
	global.Logger.Info("返回的结果:", zap.String("response", resp.String()))
	serverNode.Cookie = resp.Cookies()[0].Value
	serverNode.UpdatedAt = time.Now()
	err = serverNodeService.UpdateServerNode(serverNode)
	return
}

func (serverNodeService *ServerNodeService) SyncLadderCookie() (err error) {
	list, err := serverNodeService.GetServerNodeList()
	for _, node := range list {
		err = serverNodeService.ServerNodeLogin(node)
		if err != nil {
			return
		}
	}
	return
}
