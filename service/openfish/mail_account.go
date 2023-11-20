package openfish

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/ladder"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"github.com/sashabaranov/go-openai"
	"net/http"
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
	//proxyUrl := "http://127.0.0.1:7890"
	for _, account := range mailAccounts {
		//authenticator := openaiReverse.NewAuthenticator(account.Username, account.OpenaiPassword, proxyUrl)
		//authErr := authenticator.Begin()
		//if authErr != nil {
		//	global.Logger.Error(fmt.Sprintf("Location: %s, Status code: %s, Details: %s, Embedded error: %s", authErr.Location, fmt.Sprint(authErr.StatusCode), authErr.Details, authErr.Error.Error()))
		//}
		//accessToken := authenticator.GetAccessToken()
		data := map[string]interface{}{
			"username": account.Username,
			"password": account.OpenaiPassword,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("JSON编码失败:", err)
			continue
		}
		response, err := http.Post("http://localhost:8998/chatgpt/login", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("请求失败:", err)
		}
		defer response.Body.Close()

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			fmt.Println("解析失败:", err)
			continue
		}

		formattedResult := make(map[string]string)
		for key, value := range result {
			formattedResult[key] = fmt.Sprintf("%v", value)
		}

		accessToken := formattedResult["data"]
		// missing access token
		// Email is not valid.
		// captcha required
		// You do not have an account because it has been deleted or deactivated. If you believe this was an error, please contact us through our help center at help.openai.com. (error=account_deactivated)
		fmt.Println(accessToken)
		// artyom3egprod@email.com
		//arturvim2o@email.com
		//arturqdwck@email.com
		//artyom3egprod@email.com
		//arturmylip@email.com
		//arturmbv33@email.com
		//arturh3fadeev@email.com
		//arturd2d@email.com
		//arturanib5@email.com

		updateColumns := make(map[string]interface{})
		updateColumns["openai_access_token"] = accessToken
		updateColumns["openai_access_token_get_time"] = time.Now()
		if formattedResult["code"] == "7" {
			// 账号被封了
			updateColumns["openai_status"] = 0
			err = global.DB.Model(&openfish.MailAccount{}).Where("id = ?", account.ID).Updates(&updateColumns).Error
			if err != nil {
				global.Logger.Error("ID: " + strconv.Itoa(int(account.ID)) + ", 更新出错")
			}
			continue
		}
		if formattedResult["code"] == "0" {
			updateColumns["openai_status"] = 1
			err = global.DB.Model(&openfish.MailAccount{}).Where("id = ?", account.ID).Updates(&updateColumns).Error
			if err != nil {
				global.Logger.Error("ID: " + strconv.Itoa(int(account.ID)) + ", 更新出错")
			}
		}
	}
	return err
}

// SyncOpenaiInfo 同步openai的信息: sk到期时间、余额
func (mailAccountService *MailAccountService) SyncOpenaiInfo(ids request.IdsReq) (err error) {
	var mailAccounts []openfish.MailAccount
	err = global.DB.Where("id in ?", ids.Ids).Find(&mailAccounts).Error
	// TODO
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
	defer queueMutex.Unlock()
	err = global.DB.Where("openai_access_token != '' and openai_status = 1").Order("updated_at").First(&mailAccount).Error
	if err != nil {
		return
	}
	err = global.DB.Model(&openfish.MailAccount{}).Where("id = ?", mailAccount.ID).Update("updated_at", time.Now()).Error
	return
}

// GetServerNodeByUpdatedAtAsc 获取 ServerNode by updated_at asc
func (mailAccountService *MailAccountService) GetServerNodeByUpdatedAtAsc() (server string, err error) {
	queueMutex.Lock()
	defer queueMutex.Unlock()
	var serverNode ladder.ServerNode
	err = global.DB.Where("server_host != '' and is_openai_server = 1").Order("updated_at").First(&serverNode).Error
	if err != nil {
		return
	}
	server = "http://" + serverNode.ServerHost + ":9332"
	err = global.DB.Model(&ladder.ServerNode{}).Where("id = ?", serverNode.ID).Update("updated_at", time.Now()).Error
	return
}

// GetOpenaiKeyByUpdatedAtAsc 获取 AccessToken by updated_at asc
func (mailAccountService *MailAccountService) GetOpenaiKeyByUpdatedAtAsc() (mailAccount openfish.MailAccount, err error) {
	queueMutex.Lock()
	defer queueMutex.Unlock()
	err = global.DB.Where("openai_sk != ''").Order("sk_used_at").First(&mailAccount).Error
	err = global.DB.Model(&openfish.MailAccount{}).Where("id = ?", mailAccount.ID).Update("sk_used_at", time.Now()).Error
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
	err = db.Where("openai_status = ?", 2).Order("created_at desc").Find(&list).Error
	return
}

// GetMailAccountInfoList 分页获取MailAccount记录
func (mailAccountService *MailAccountService) GetMailAccountInfoList(info openfishReq.MailAccountSearch) (list []openfish.MailAccount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&openfish.MailAccount{})
	db = db.Where("openai_status > ?", 0)
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

func (mailAccountService *MailAccountService) SyncChatGPTAccessToken() {
	list, errGetMailAccountList := mailAccountService.GetMailAccountList()
	if errGetMailAccountList != nil {
		fmt.Println("同步 OpenAI ChatGPT accessToken 时，获取账户列表失败")
		return
	}
	for _, account := range list {
		// 获取当前时间
		currentTime := time.Now()
		// 计算五天前的时间
		threeDaysAgo := currentTime.Add(-5 * 24 * time.Hour)
		// 比较时间
		if account.OpenaiAccessTokenGetTime == nil || account.OpenaiAccessTokenGetTime.Before(threeDaysAgo) {
			var ids request.IdsReq
			ids.Ids = append(ids.Ids, int(account.ID))
			fmt.Println("正在同步: ", account.Username)
			errRefreshOpenaiAccessToken := mailAccountService.RefreshOpenaiAccessToken(ids)
			if errRefreshOpenaiAccessToken != nil {
				fmt.Println("同步 OpenAI ChatGPT accessToken 失败")
			}
			fmt.Println("同步完成: ", account.Username)
			time.Sleep(10 * time.Second)
		}
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " 完成所有账号AT同步")
}

func (mailAccountService *MailAccountService) SyncChatGPTAccessTokenStatus() {
	list, errGetMailAccountList := mailAccountService.GetMailAccountList()
	if errGetMailAccountList != nil {
		fmt.Println("修改所有账号AT状态时，获取账户列表失败")
		return
	}
	for _, account := range list {
		if *account.OpenaiStatus == 1 {
			continue
		}
		var ids request.IdsReq
		ids.Ids = append(ids.Ids, int(account.ID))
		errRefreshOpenaiAccessToken := mailAccountService.UpdateOpenaiStatus(ids)
		if errRefreshOpenaiAccessToken != nil {
			fmt.Println("修改账号AT状态失败")
		}
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " 修改所有账号AT状态完成")
}

func (mailAccountService *MailAccountService) UpdateOpenaiStatus(ids request.IdsReq) (err error) {
	err = global.DB.Model(&openfish.MailAccount{}).Where("id in ?", ids.Ids).Update("openai_status", 1).Error
	return
}
