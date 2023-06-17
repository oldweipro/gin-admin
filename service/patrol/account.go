package patrol

import (
	"fmt"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
)

type AccountService struct {
}

// CreateAccount 创建Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) CreateAccount(account patrol.Account) (err error) {
	err = global.DB.Create(&account).Error
	return err
}

// DeleteAccount 删除Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) DeleteAccount(account patrol.Account) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.Account{}).Where("id = ?", account.ID).Update("deleted_by", account.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&account).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAccountByIds 批量删除Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) DeleteAccountByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.Account{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&patrol.Account{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAccount 更新Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) UpdateAccount(account patrol.Account) (err error) {
	err = global.DB.Save(&account).Error
	return err
}

// GetAccount 根据id获取Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) GetAccount(id uint) (account patrol.Account, err error) {
	err = global.DB.Where("id = ?", id).First(&account).Error
	return
}

// GetAccountInfoList 分页获取Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) GetAccountInfoList(info patrolReq.AccountSearch) (list []patrol.Account, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&patrol.Account{})
	var accounts []patrol.Account
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AccountName != "" {
		db = db.Where("account_name LIKE ?", "%"+info.AccountName+"%")
	}
	if info.LoginStatus != nil {
		db = db.Where("login_status = ?", info.LoginStatus)
	}
	if info.CurrentCalls != nil {
		db = db.Where("current_calls < ?", info.CurrentCalls)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&accounts).Error
	return accounts, total, err
}

func (accountService *AccountService) LoginGameAccount(account patrol.Account) (string, error) {
	params := url.Values{}
	params.Add("do", "login")
	params.Add("gourl", "")
	params.Add("login_save", "1")
	params.Add("username", account.AccountName)
	params.Add("password", account.AccountName)
	resp, err := http.PostForm("",
		params)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	bodyResult := string(body)
	global.Logger.Info("登陆游戏账号返回结果: " + bodyResult)
	if err != nil {
		// handle error
		global.Logger.Error("登陆失败!", zap.Error(err))
	}
	if err := accountService.UpdateAccount(account); err != nil {
		global.Logger.Error("登陆失败,数据库更新失败!", zap.Error(err))
		return bodyResult, err
	}
	return bodyResult, err
}
