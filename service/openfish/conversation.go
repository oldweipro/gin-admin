package openfish

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"github.com/oldweipro/gin-admin/service/system"
	"github.com/oldweipro/gin-admin/utils"
	"github.com/oldweipro/gin-admin/utils/upload"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ConversationService struct {
}

var chatGptService system.ChatGptService
var promptService PromptService

// CreateConversation 创建Conversation记录
func (conversationService *ConversationService) CreateConversation(conversation *openfish.Conversation) (err error) {
	err = global.DB.Create(conversation).Error
	return err
}

// CreateConversationRecord 创建ConversationRecord记录
func (conversationService *ConversationService) CreateConversationRecord(conversationRecord *openfish.ConversationRecord) (err error) {
	err = global.DB.Create(conversationRecord).Error
	return err
}

// DeleteConversation 删除Conversation记录
func (conversationService *ConversationService) DeleteConversation(conversation openfish.Conversation) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.Conversation{}).Where("id = ?", conversation.ID).Update("deleted_by", conversation.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&conversation).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteConversationByIds 批量删除Conversation记录
func (conversationService *ConversationService) DeleteConversationByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.Conversation{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&openfish.Conversation{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateConversation 更新Conversation记录
func (conversationService *ConversationService) UpdateConversation(conversation openfish.Conversation) (err error) {
	err = global.DB.Save(&conversation).Error
	return err
}

// UpdateConversationTime 更新Conversation时间
func (conversationService *ConversationService) UpdateConversationTime(id uint) (err error) {
	err = global.DB.Model(&openfish.Conversation{}).Where("id = ?", id).Update("updated_at", time.Now()).Error
	return err
}

// GetConversation 根据id获取Conversation记录
func (conversationService *ConversationService) GetConversation(id uint) (conversation openfish.Conversation, err error) {
	err = global.DB.Where("id = ?", id).First(&conversation).Error
	return
}

// GetConversationInfoList 分页获取Conversation记录
func (conversationService *ConversationService) GetConversationInfoList(info openfishReq.ConversationSearch) (list []openfish.Conversation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&openfish.Conversation{})
	var conversations []openfish.Conversation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CreatedBy != 0 {
		db = db.Where("created_by = ?", info.CreatedBy)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&conversations).Error
	return conversations, total, err
}

func (conversationService *ConversationService) GetConversationRecordList(info openfishReq.ConversationRecordSearch) (list []openfish.ConversationRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&openfish.ConversationRecord{})
	var conversationRecords []openfish.ConversationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CreatedBy != 0 {
		db = db.Where("created_by = ?", info.CreatedBy)
	}
	if info.Content != "" {
		db = db.Where("content like ?", "%"+info.Content+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&conversationRecords).Error
	return conversationRecords, total, err
}

// GetConversationRecordListWithTokenByConversationId 根据pid查询会话信息列表
// error: error, status code: 400, message: This model's maximum context length is 4097 tokens. However, your messages resulted in 6301 tokens. Please reduce the length of the messages.
func (conversationService *ConversationService) GetConversationRecordListWithTokenByConversationId(conversationId uint, tokenCount int) ([]openfish.ConversationRecord, error) {
	var conversationRecords []openfish.ConversationRecord
	query := `SELECT *
		FROM conversation_record
		WHERE id IN (
		  SELECT id
		  FROM (
			SELECT id, created_by,created_at,
				   @sum := @sum + CHAR_LENGTH(content) AS sum
			FROM conversation_record, 
				 (SELECT @sum := 0) AS vars 
				WHERE conversation_id = ?
			ORDER BY created_at desc
		  ) AS t
		  WHERE sum <= ?
		) ORDER BY created_at`
	err := global.DB.Raw(query, conversationId, 2596-tokenCount).Scan(&conversationRecords).Error
	return conversationRecords, err
}

// GetConversationRecordListByUserId 根据用户ID查询会话列表及会话信息列表
func (conversationService *ConversationService) GetConversationRecordListByUserId(userId uint) ([]openfish.ConversationRecord, error) {
	var conversationRecords []openfish.ConversationRecord
	err := global.DB.Model(&openfish.ConversationRecord{}).Where("conversation_id IN (SELECT id FROM conversation WHERE created_by = ?)", userId).Order("created_at asc").Find(&conversationRecords).Error
	return conversationRecords, err
}

// GetConversationRecordListByConversationId 根据conversationId查询会话信息列表
func (conversationService *ConversationService) GetConversationRecordListByConversationId(conversationId uint) ([]openfish.ConversationRecord, error) {
	var conversationRecords []openfish.ConversationRecord
	err := global.DB.Model(&openfish.ConversationRecord{}).Where("conversation_id = ?", conversationId).Order("created_at asc").Find(&conversationRecords).Error
	return conversationRecords, err
}

// GetConversationListByUserId 根据用户ID查询会话列表
func (conversationService *ConversationService) GetConversationListByUserId(userId uint, conversationType uint) ([]openfish.Conversation, error) {
	var conversations []openfish.Conversation
	err := global.DB.Model(&openfish.Conversation{}).Where("created_by = ?", userId).Where("conversation_type = ?", conversationType).Order("updated_at desc").First(&conversations).Error
	return conversations, err
}

// OpenAIDrawing openai作画
func (conversationService *ConversationService) OpenAIDrawing(chatReq openfishReq.ChatReq, c *gin.Context) error {
	sk, err := chatGptService.GetSK()
	if err != nil {
		global.Logger.Error("获取sk失败!", zap.Error(err))
		return err
	}
	// 更新openai sk
	sk.UpdatedAt = time.Now()
	if err := chatGptService.UpdateSK(sk); err != nil {
		global.Logger.Error("更新openai sk失败!", zap.Error(err))
	}
	config := openai.DefaultConfig(sk.SK)
	// 如果需要代理，请配置代理地址，如不需要可注释或删掉以下代码
	config.HTTPClient.Transport = &http.Transport{
		// 设置Transport字段为自定义Transport，包含代理设置
		Proxy: func(req *http.Request) (*url.URL, error) {
			// 设置代理
			proxyURL, err := url.Parse("http://127.0.0.1:7890")
			if err != nil {
				return nil, err
			}
			return proxyURL, nil
		},
	}
	client := openai.NewClientWithConfig(config)
	ctx := context.Background()
	imageRequest := openai.ImageRequest{
		Prompt:         chatReq.Prompt,
		N:              1,
		Size:           "1024x1024",
		ResponseFormat: "url",
		User:           strconv.Itoa(int(utils.GetUserID(c))),
	}
	if image, err := client.CreateImage(ctx, imageRequest); err != nil {
		return err
	} else {
		streamResponse := ""
		for _, img := range image.Data {
			fmt.Println(img.URL)
			// 图片存入本地或者OSS
			oss := upload.NewOss()
			uploadUrl, _, imgErr := oss.UploadUrl(img.URL, "")
			if imgErr != nil {
				return imgErr
			}

			server := make(map[string]string)
			server["content"] = "![](" + uploadUrl + ")"
			marshal, _ := json.Marshal(server)
			sse.Encode(c.Writer, sse.Event{
				Data: string(marshal),
			})
			c.Writer.Flush()
			if streamResponse == "" {
				streamResponse = server["content"]
			} else {
				streamResponse = streamResponse + "\n\n" + server["content"]
			}
		}
		// 数据存入数据库 预备存储新的聊天记录
		conversationRecordUser := openfish.ConversationRecord{}
		conversationRecordUser.Content = chatReq.Prompt
		conversationRecordUser.Role = "user"
		conversationRecordUser.ConversationId = chatReq.ConversationId
		conversationRecordUser.CreatedBy = utils.GetUserID(c)
		// 最后存储新的对话到数据库 提问
		if err := conversationService.CreateConversationRecord(&conversationRecordUser); err != nil {
			global.Logger.Error("用户提问数据写入异常!", zap.Error(err))
			response.FailWithMessage("系统异常", c)
			return err
		}
		// 最后存储新的对话到数据库 回答
		conversationRecordAI := openfish.ConversationRecord{}
		conversationRecordAI.Content = streamResponse
		conversationRecordAI.Role = "assistant"
		conversationRecordAI.ConversationId = chatReq.ConversationId
		conversationRecordAI.CreatedBy = utils.GetUserID(c)
		if err := conversationService.CreateConversationRecord(&conversationRecordAI); err != nil {
			global.Logger.Error("AI回答数据写入异常!", zap.Error(err))
			response.FailWithMessage("系统异常", c)
			return err
		}
		// 更新聊天室时间
		if err := conversationService.UpdateConversationTime(*chatReq.ConversationId); err != nil {
			global.Logger.Error("更新聊天室时间异常!", zap.Error(err))
			response.FailWithMessage("系统异常", c)
			return err
		}
	}
	return nil
}

// ChatGPTCompletions ChatGPT对话方法
func (conversationService *ConversationService) ChatGPTCompletions(chatReq openfishReq.ChatReq, c *gin.Context) error {
	tokenCount := conversationService.NumTokens(chatReq.Prompt)
	// 查询会话记录
	conversationRecordList, _ := conversationService.GetConversationRecordListWithTokenByConversationId(*chatReq.ConversationId, tokenCount)
	var messages []openai.ChatCompletionMessage
	// 判断是否提示词问答
	if chatReq.PromptId > 0 {
		// 查询提示词信息
		prompt, err := promptService.GetPrompt(chatReq.PromptId)
		if err != nil {
			fmt.Println("查询 GetPrompt 出错", err)
		}
		// 组装openai消息参数
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    "user",
			Content: prompt.Content,
		})
		// 组装openai消息参数
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    "assistant",
			Content: "让我们开始吧。",
		})
		// 前端会携带，后端也暂时写死：使用提示词只能是单机模式
		chatReq.StandardAlone = 1
	}
	// 判断是否开启单机模式 0关闭 1开启
	if chatReq.StandardAlone == 1 {
	} else {
		// 如果没有开启单机模式，则查询上下文
		for _, cr := range conversationRecordList {
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    cr.Role,
				Content: cr.Content,
			})
		}
	}
	// 预备存储新的聊天记录
	conversationRecordUser := openfish.ConversationRecord{}
	conversationRecordUser.Content = chatReq.Prompt
	conversationRecordUser.Role = "user"
	conversationRecordUser.ConversationId = chatReq.ConversationId
	conversationRecordUser.CreatedBy = utils.GetUserID(c)
	// 组装新消息参数
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    conversationRecordUser.Role,
		Content: conversationRecordUser.Content,
	})
	//var completionReq openai.ChatCompletionRequest
	//numTokens(completionReq.Prompt.(string)) * completionReq.N
	//inputTokens := numTokens(completionReq.Messages[0].Content) * completionReq.N
	//completionTokens := completionReq.MaxTokens * completionReq.N
	//res.Usage = Usage{
	//	PromptTokens:     inputTokens,
	//	CompletionTokens: completionTokens,
	//	TotalTokens:      inputTokens + completionTokens,
	//}
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo0613,
		MaxTokens: 1000,
		Messages:  messages,
		Stream:    true,
	}
	fmt.Printf("%v", messages)
	if err := conversationService.ChatOpenAIReverse(&conversationRecordUser, req, c, chatReq); err != nil {
		global.Logger.Error("逆向工程调用错误: ", zap.Error(err))
		if err = conversationService.ChatOpenAIApiKey(&conversationRecordUser, req, c, chatReq); err != nil {
			global.Logger.Error("OpenAI调用错误: ", zap.Error(err))
			return err
		}
	}
	return nil
}

func (conversationService *ConversationService) NumTokens(s string) int {
	return int(float32(len(s)) / 4)
}

// ChatOpenAIReverse https://chat.openai.com逆向接口: 需要自己搭建服务 https://github.com/acheong08/ChatGPT-to-API
func (conversationService *ConversationService) ChatOpenAIReverse(conversationRecordUser *openfish.ConversationRecord, req openai.ChatCompletionRequest, c *gin.Context, chatReq openfishReq.ChatReq) error {
	config := openai.DefaultConfig("OpenAIReverse")
	config.BaseURL = "http://127.0.0.1:8080/v1"
	client := openai.NewClientWithConfig(config)
	ctx := context.Background()
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return err
	}
	defer config.HTTPClient.CloseIdleConnections()
	defer stream.Close()
	return conversationService.ChatStream(stream, conversationRecordUser, c, chatReq)
}

// ChatOpenAIApiKey 官方接口：更换TOKEN，使用代理
func (conversationService *ConversationService) ChatOpenAIApiKey(conversationRecordUser *openfish.ConversationRecord, req openai.ChatCompletionRequest, c *gin.Context, chatReq openfishReq.ChatReq) error {
	sk, err := chatGptService.GetSK()
	if err != nil {
		global.Logger.Error("获取sk失败!", zap.Error(err))
		return err
	}
	// 更新openai sk
	sk.UpdatedAt = time.Now()
	if err := chatGptService.UpdateSK(sk); err != nil {
		global.Logger.Error("更新openai sk失败!", zap.Error(err))
	}
	config := openai.DefaultConfig(sk.SK)
	// 如果需要代理，请配置代理地址，如不需要可注释或删掉以下代码
	config.HTTPClient.Transport = &http.Transport{
		// 设置Transport字段为自定义Transport，包含代理设置
		Proxy: func(req *http.Request) (*url.URL, error) {
			// 设置代理
			proxyURL, err := url.Parse("http://127.0.0.1:7890")
			if err != nil {
				return nil, err
			}
			return proxyURL, nil
		},
	}
	client := openai.NewClientWithConfig(config)
	ctx := context.Background()
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return err
	}
	defer config.HTTPClient.CloseIdleConnections()
	defer stream.Close()
	return conversationService.ChatStream(stream, conversationRecordUser, c, chatReq)
}

// ChatStream AI对话流处理
func (conversationService *ConversationService) ChatStream(stream *openai.ChatCompletionStream, conversationRecordUser *openfish.ConversationRecord, c *gin.Context, chatReq openfishReq.ChatReq) error {

	var streamResponse string
	for {
		respResult, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			if streamResponse != "" {
				// 最后存储新的对话到数据库 提问
				if err := conversationService.CreateConversationRecord(conversationRecordUser); err != nil {
					global.Logger.Error("用户提问数据写入异常!", zap.Error(err))
					response.FailWithMessage("系统异常", c)
					return err
				}
				// 最后存储新的对话到数据库 回答 TODO 增加token计数
				conversationRecordAI := openfish.ConversationRecord{}
				conversationRecordAI.Content = streamResponse
				conversationRecordAI.Role = "assistant"
				conversationRecordAI.ConversationId = chatReq.ConversationId
				conversationRecordAI.CreatedBy = utils.GetUserID(c)
				if err := conversationService.CreateConversationRecord(&conversationRecordAI); err != nil {
					global.Logger.Error("AI回答数据写入异常!", zap.Error(err))
					response.FailWithMessage("系统异常", c)
					return err
				}
				// 更新聊天室时间
				if err := conversationService.UpdateConversationTime(*chatReq.ConversationId); err != nil {
					global.Logger.Error("更新聊天室时间异常!", zap.Error(err))
					response.FailWithMessage("系统异常", c)
					return err
				}
			}
			return nil
		}

		if err != nil {
			return err
		}
		streamResponse += respResult.Choices[0].Delta.Content
		server := make(map[string]string)
		server["content"] = respResult.Choices[0].Delta.Content
		marshal, _ := json.Marshal(server)
		sse.Encode(c.Writer, sse.Event{
			Data: string(marshal),
		})
		c.Writer.Flush()
	}
}
