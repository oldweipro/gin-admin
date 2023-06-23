package ladder

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/ladder"
	ladderReq "github.com/oldweipro/gin-admin/model/ladder/request"
	"gorm.io/gorm"
	"strconv"
)

type ServerNodeService struct {
}

// CreateServerNode 创建ServerNode记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverNodeService *ServerNodeService) CreateServerNode(serverNode *ladder.ServerNode) (err error) {
	err = global.DB.Create(serverNode).Error
	return err
}

// DeleteServerNode 删除ServerNode记录
// Author [piexlmax](https://github.com/piexlmax)
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
// Author [piexlmax](https://github.com/piexlmax)
func (serverNodeService *ServerNodeService) DeleteServerNodeByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ladder.ServerNode{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
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
// Author [piexlmax](https://github.com/piexlmax)
func (serverNodeService *ServerNodeService) UpdateServerNode(serverNode ladder.ServerNode) (err error) {
	err = global.DB.Save(&serverNode).Error
	return err
}

// GetServerNode 根据id获取ServerNode记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverNodeService *ServerNodeService) GetServerNode(id uint) (serverNode ladder.ServerNode, err error) {
	err = global.DB.Where("id = ?", id).First(&serverNode).Error
	return
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
	db = db.Where("server_status = 1")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&serverNodes).Error
	return serverNodes, total, err
}

// GetServerNodeLessInfoList 分页获取ServerNode记录，过滤掉敏感信息
func (serverNodeService *ServerNodeService) GetServerNodeLessInfoList(info ladderReq.ServerNodeSearch) (list []ladder.ServerNode, total int64, err error) {
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

	err = db.Select("ID", "bandwidth", "region", "describe").Limit(limit).Offset(offset).Find(&serverNodes).Error
	return serverNodes, total, err
}

// ServerNodeLogin 处理登陆获取cookie
func (serverNodeService *ServerNodeService) ServerNodeLogin(serverNode ladder.ServerNode) {
	loginFormData := make(map[string]string)
	loginFormData["username"] = serverNode.Username
	loginFormData["password"] = serverNode.Password
	client := resty.New()
	resp, err := client.R().
		SetFormData(loginFormData).
		Post("http://" + serverNode.ServerHost + ":" + strconv.Itoa(*serverNode.ServerPort) + "/login")
	if err != nil {
		fmt.Println(err)
	}
	serverNode.Cookie = resp.Cookies()[0].Value
	fmt.Println(serverNode.Cookie)
	serverNodeService.UpdateServerNode(serverNode)
}
