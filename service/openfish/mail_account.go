package openfish

import (
	"errors"
	"fmt"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"github.com/oldweipro/gin-admin/utils/openai"
	"strconv"
	"time"
)

type MailAccountService struct {
}

// RefreshClaudeChat 产生一次Claude对话
func (mailAccountService *MailAccountService) RefreshClaudeChat(ids request.IdsReq) (err error) {
	var mailAccounts []openfish.MailAccount
	err = global.DB.Where("id in ?", ids.Ids).Find(&mailAccounts).Error
	//
	return err
}

// RefreshOpenaiAccessToken 刷新 openai AccessToken
func (mailAccountService *MailAccountService) RefreshOpenaiAccessToken(ids request.IdsReq) (err error) {
	var mailAccounts []openfish.MailAccount
	err = global.DB.Where("id in ?", ids.Ids).Find(&mailAccounts).Error
	// 循环遍历
	proxyUrl := "http://127.0.0.1:7890"
	for _, account := range mailAccounts {
		authenticator := openai.NewAuthenticator(account.Username, account.OpenaiPassword, proxyUrl)
		authErr := authenticator.Begin()
		if authErr != nil {
			return errors.New(fmt.Sprintf("Location: %s, Status code: %s, Details: %s, Embedded error: %s", authErr.Location, fmt.Sprint(authErr.StatusCode), authErr.Details, authErr.Error.Error()))
		}
		accessToken := authenticator.GetAccessToken()
		updateColumns := make(map[string]interface{})
		updateColumns["openai_access_token"] = accessToken
		updateColumns["openai_access_token_get_time"] = time.Now()
		err = global.DB.Model(&openfish.MailAccount{}).Where("id = ?", account.ID).Updates(&updateColumns).Error
		if err != nil {
			global.Logger.Error("ID: " + strconv.Itoa(int(account.ID)) + ", 更新出错")
		}
	}
	return err
}

// SyncOpenaiInfo 同步openai的信息: sk到期时间、余额
func (mailAccountService *MailAccountService) SyncOpenaiInfo(ids request.IdsReq) (err error) {
	var mailAccounts []openfish.MailAccount
	err = global.DB.Where("id in ?", ids.Ids).Find(&mailAccounts).Error
	//
	return err
}

// CreateMailAccount 创建MailAccount记录
func (mailAccountService *MailAccountService) CreateMailAccount(mailAccount *openfish.MailAccount) (err error) {
	err = global.DB.Create(mailAccount).Error
	return err
}

// DeleteMailAccount 删除MailAccount记录
func (mailAccountService *MailAccountService) DeleteMailAccount(mailAccount openfish.MailAccount) (err error) {
	err = global.DB.Delete(&mailAccount).Error
	return err
}

// DeleteMailAccountByIds 批量删除MailAccount记录
func (mailAccountService *MailAccountService) DeleteMailAccountByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]openfish.MailAccount{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMailAccount 更新MailAccount记录
func (mailAccountService *MailAccountService) UpdateMailAccount(mailAccount openfish.MailAccount) (err error) {
	err = global.DB.Save(&mailAccount).Error
	return err
}

// GetMailAccount 根据id获取MailAccount记录
func (mailAccountService *MailAccountService) GetMailAccount(id uint) (mailAccount openfish.MailAccount, err error) {
	err = global.DB.Where("id = ?", id).First(&mailAccount).Error
	return
}

// GetAccessTokenByUpdatedAtAsc 获取 AccessToken by updated_at asc
func (mailAccountService *MailAccountService) GetAccessTokenByUpdatedAtAsc() (mailAccount openfish.MailAccount, err error) {
	err = global.DB.Where("openai_access_token != ''").Order("updated_at").First(&mailAccount).Error
	return
}

// UpdateAccessTokenWithUpdatedAt 更新 AccessToken's UpdatedAt
func (mailAccountService *MailAccountService) UpdateAccessTokenWithUpdatedAt(id uint) (err error) {
	err = global.DB.Model(&openfish.MailAccount{}).Where("id = ?", id).Update("updated_at", time.Now()).Error
	return
}

// GetMailAccountInfoList 分页获取MailAccount记录
func (mailAccountService *MailAccountService) GetMailAccountInfoList(info openfishReq.MailAccountSearch) (list []openfish.MailAccount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&openfish.MailAccount{})
	var mailAccounts []openfish.MailAccount
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&mailAccounts).Error
	return mailAccounts, total, err
}
