package openfish

import (
	"context"
	"fmt"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"github.com/oldweipro/gin-admin/utils/openai_reverse"
	"github.com/sashabaranov/go-openai"
	"strconv"
	"sync"
	"time"
)

type MailAccountService struct {
}

var queueMutex sync.Mutex

// RefreshClaudeChat 产生一次Claude对话
func (mailAccountService *MailAccountService) RefreshClaudeChat(ids request.IdsReq) (err error) {
	var mailAccounts []openfish.MailAccount
	err = global.DB.Where("claude_session_key != '' and id in ?", ids.Ids).Find(&mailAccounts).Error
	if err != nil {
		return err
	}
	// 循环遍历
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo0613,
		MaxTokens: 1000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "你好啊",
			},
		},
		Stream: false,
	}
	go func() {
		for _, account := range mailAccounts {
			config := openai.DefaultConfig(account.ClaudeSessionKey)
			config.BaseURL = "http://127.0.0.1:8787/v1"
			client := openai.NewClientWithConfig(config)
			ctx := context.Background()
			response, err := client.CreateChatCompletion(ctx, req)
			config.HTTPClient.CloseIdleConnections()
			if err != nil {
				global.Logger.Error(err.Error())
			} else {
				global.Logger.Info(response.Choices[0].Message.Content)
				// 更新数据库时间
				updateColumns := make(map[string]interface{})
				updateColumns["claude_session_key_get_time"] = time.Now()
				err = global.DB.Model(&openfish.MailAccount{}).Where("id = ?", account.ID).Updates(&updateColumns).Error
				if err != nil {
					global.Logger.Error("ID: " + strconv.Itoa(int(account.ID)) + ", 更新出错")
				}
			}
		}
	}()
	return err
}

// RefreshOpenaiAccessToken 刷新 openai AccessToken
func (mailAccountService *MailAccountService) RefreshOpenaiAccessToken(ids request.IdsReq) (err error) {
	var mailAccounts []openfish.MailAccount
	err = global.DB.Where("id in ?", ids.Ids).Find(&mailAccounts).Error
	// 循环遍历
	proxyUrl := "http://127.0.0.1:7890"
	go func() {
		for _, account := range mailAccounts {
			authenticator := openaiReverse.NewAuthenticator(account.Username, account.OpenaiPassword, proxyUrl)
			authErr := authenticator.Begin()
			if authErr != nil {
				global.Logger.Error(fmt.Sprintf("Location: %s, Status code: %s, Details: %s, Embedded error: %s", authErr.Location, fmt.Sprint(authErr.StatusCode), authErr.Details, authErr.Error.Error()))
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
	}()
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
	queueMutex.Lock()
	err = global.DB.Where("openai_access_token != ''").Order("updated_at").First(&mailAccount).Error
	err = global.DB.Model(&openfish.MailAccount{}).Where("id = ?", mailAccount.ID).Update("updated_at", time.Now()).Error
	defer queueMutex.Unlock()
	return
}

// UpdateAccessTokenWithUpdatedAt 更新 AccessToken's UpdatedAt
func (mailAccountService *MailAccountService) UpdateAccessTokenWithUpdatedAt(id uint) (err error) {
	err = global.DB.Model(&openfish.MailAccount{}).Where("id = ?", id).Update("updated_at", time.Now()).Error
	return
}

// GetMailAccountList 获取MailAccount列表记录
func (mailAccountService *MailAccountService) GetMailAccountList() (list []openfish.MailAccount, err error) {
	db := global.DB.Model(&openfish.MailAccount{})
	err = db.Order("created_at desc").Find(&list).Error
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
