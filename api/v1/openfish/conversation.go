package openfish

import (
	"context"
	"crypto/md5"
	"encoding/hex"
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
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"io"
	"net/http"
	"sort"
	"strconv"
)

type ConversationApi struct {
}

var conversationService = service.ServiceGroupApp.OpenfishServiceGroup.ConversationService

// 并发数为n的信号量
var semaphore = make(chan struct{}, 4)

// ChatCompletions 使用第三方库新接口
func (conversationApi *ConversationApi) ChatCompletions(c *gin.Context) {
	semaphore <- struct{}{} // 获取信号量
	defer func() { <-semaphore }()
	// 获取参数
	var chatReq openfishReq.ChatReq
	err := c.ShouldBindJSON(&chatReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	if chatReq.Prompt == "" || chatReq.ConversationId == nil {
		response.FailWithMessage("系统错误，缺少必要参数。", c)
		global.GVA_LOG.Error("系统错误，缺少必要参数。")
		return
	}
	// md5校验参数
	str := chatReq.Prompt + "-" + strconv.FormatUint(uint64(*chatReq.ConversationId), 10) + "5eb63bbbe01eeed093cb22bb8f5acdc3"
	hash := md5.Sum([]byte(str))
	hexHash := hex.EncodeToString(hash[:])
	if hexHash != chatReq.Sign {
		response.FailWithMessage("系统错误，缺少必要参数。", c)
		return
	}
	c.Status(http.StatusOK)
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	tokenCount := len(chatReq.Prompt)
	// 查询会话记录
	conversationRecordList, _ := conversationService.GetConversationRecordListByConversationId(*chatReq.ConversationId, tokenCount)
	var messages []openai.ChatCompletionMessage
	for _, cr := range conversationRecordList {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    cr.Role,
			Content: cr.Content,
		})
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
	// ==================OpenAI调用开始==================
	ctx := context.Background()

	// =======start 官方接口：更换TOKEN，使用代理=======
	config := openai.DefaultConfig("TOKEN")
	// 如果需要代理，请配置代理地址，如不需要可注释或删掉以下代码
	//config.HTTPClient.Transport = &http.Transport{
	//	// 设置Transport字段为自定义Transport，包含代理设置
	//	Proxy: func(req *http.Request) (*url.URL, error) {
	//		// 设置代理
	//		proxyURL, err := url.Parse("http://127.0.0.1:7890")
	//		if err != nil {
	//			return nil, err
	//		}
	//		return proxyURL, nil
	//	},
	//}
	// =======end 官方接口：更换TOKEN，使用代理=======

	// =======start 逆向官网接口：使用逆向工程=======
	config.BaseURL = "http://127.0.0.1:8080/v1"
	// =======end 逆向官网接口：使用逆向工程=======

	client := openai.NewClientWithConfig(config)
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo0301,
		MaxTokens: 1000,
		Messages:  messages,
		Stream:    true,
	}
	defer config.HTTPClient.CloseIdleConnections()
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("聊天对话异常 error: %v\n", err)
		return
	}
	defer stream.Close()

	var streamResponse string
	for {
		respResult, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println()
			fmt.Println("程序调用结束")
			if streamResponse != "" {
				// 最后存储新的对话到数据库 提问
				if err := conversationService.CreateConversationRecord(&conversationRecordUser); err != nil {
					global.GVA_LOG.Error("用户提问数据写入异常!", zap.Error(err))
					response.FailWithMessage("系统异常", c)
					return
				}
				// 最后存储新的对话到数据库 回答
				conversationRecordAI := openfish.ConversationRecord{}
				conversationRecordAI.Content = streamResponse
				conversationRecordAI.Role = "assistant"
				conversationRecordAI.ConversationId = chatReq.ConversationId
				conversationRecordAI.CreatedBy = utils.GetUserID(c)
				if err := conversationService.CreateConversationRecord(&conversationRecordAI); err != nil {
					global.GVA_LOG.Error("AI回答数据写入异常!", zap.Error(err))
					response.FailWithMessage("系统异常", c)
					return
				}
				// 更新聊天室时间
				if err := conversationService.UpdateConversationTime(*chatReq.ConversationId); err != nil {
					global.GVA_LOG.Error("更新聊天室时间异常!", zap.Error(err))
					response.FailWithMessage("系统异常", c)
					return
				}
			}
			return
		}

		if err != nil {
			return
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

// CreateConversation 创建Conversation
// @Tags Conversation
// @Summary 创建Conversation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Conversation true "创建Conversation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /conversation/createConversation [post]
func (conversationApi *ConversationApi) CreateConversation(c *gin.Context) {
	var conversation openfish.Conversation
	err := c.ShouldBindJSON(&conversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	conversation.CreatedBy = utils.GetUserID(c)
	if err := conversationService.CreateConversation(&conversation); err != nil {
		global.GVA_LOG.Error("创建会话失败!", zap.Error(err))
		response.FailWithMessage("系统异常", c)
		return
	}
	// conversationId等于空则创建该信息到数据库
	conversationRecord := openfish.ConversationRecord{}
	conversationRecord.Content = "我是由开放鱼训练的大型语言模型，请详细描述您的问题。"
	conversationRecord.Role = "system"
	conversationRecord.ConversationId = &conversation.ID
	conversationRecord.CreatedBy = utils.GetUserID(c)
	if err := conversationService.CreateConversationRecord(&conversationRecord); err != nil {
		global.GVA_LOG.Error("AI会话创建失败!", zap.Error(err))
		response.FailWithMessage("系统异常", c)
	} else {
		response.OkWithData(gin.H{"conversationId": conversation.ID}, c)
	}
}

// DeleteConversation 删除Conversation
// @Tags Conversation
// @Summary 删除Conversation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Conversation true "删除Conversation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /conversation/deleteConversation [delete]
func (conversationApi *ConversationApi) DeleteConversation(c *gin.Context) {
	var conversation openfish.Conversation
	err := c.ShouldBindJSON(&conversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	conversation.DeletedBy = utils.GetUserID(c)
	if err := conversationService.DeleteConversation(conversation); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteConversationByIds 批量删除Conversation
// @Tags Conversation
// @Summary 批量删除Conversation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Conversation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /conversation/deleteConversationByIds [delete]
func (conversationApi *ConversationApi) DeleteConversationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := conversationService.DeleteConversationByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateConversation 更新Conversation
// @Tags Conversation
// @Summary 更新Conversation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Conversation true "更新Conversation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /conversation/updateConversation [put]
func (conversationApi *ConversationApi) UpdateConversation(c *gin.Context) {
	var conversation openfish.Conversation
	err := c.ShouldBindJSON(&conversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	conversation.UpdatedBy = utils.GetUserID(c)
	if err := conversationService.UpdateConversation(conversation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// UpdateConversationName 更新聊天室名字
func (conversationApi *ConversationApi) UpdateConversationName(c *gin.Context) {
	var conversation openfish.Conversation
	err := c.ShouldBindJSON(&conversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	getConversation, err := conversationService.GetConversation(conversation.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	getConversation.ConversationName = conversation.ConversationName
	getConversation.UpdatedBy = utils.GetUserID(c)
	if err := conversationService.UpdateConversation(getConversation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindConversation 用id查询Conversation
// @Tags Conversation
// @Summary 用id查询Conversation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfish.Conversation true "用id查询Conversation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /conversation/findConversation [get]
func (conversationApi *ConversationApi) FindConversation(c *gin.Context) {
	var conversation openfish.Conversation
	err := c.ShouldBindQuery(&conversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reconversation, err := conversationService.GetConversation(conversation.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reconversation": reconversation}, c)
	}
}

// GetConversationList 分页获取Conversation列表
// @Tags Conversation
// @Summary 分页获取Conversation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfishReq.ConversationSearch true "分页获取Conversation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /conversation/getConversationList [get]
func (conversationApi *ConversationApi) GetConversationList(c *gin.Context) {
	var pageInfo openfishReq.ConversationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := conversationService.GetConversationInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetCurrentUserConversationList 分页获取当前用户对话列表
// @Tags Conversation
// @Summary 分页获取当前用户对话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfishReq.ConversationSearch true "分页获取当前用户对话列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /conversation/getConversationList [get]
func (conversationApi *ConversationApi) GetCurrentUserConversationList(c *gin.Context) {
	userInfo := utils.GetUserInfo(c)
	conversationList, err := conversationService.GetConversationListByUserId(userInfo.BaseClaims.ID)

	if len(conversationList) == 0 {
		// 创建一个聊天
		var conversation openfish.Conversation
		conversation.ConversationName = "新聊天"
		conversation.CreatedBy = utils.GetUserID(c)
		if err := conversationService.CreateConversation(&conversation); err != nil {
			global.GVA_LOG.Error("创建会话失败!", zap.Error(err))
			response.FailWithMessage("系统异常", c)
			return
		}
		// conversationId等于空则创建该信息到数据库
		conversationRecord := openfish.ConversationRecord{}
		conversationRecord.Content = "我是由开放鱼训练的大型语言模型，请详细描述您的问题。"
		conversationRecord.Role = "system"
		conversationRecord.ConversationId = &conversation.ID
		conversationRecord.CreatedBy = utils.GetUserID(c)
		if err := conversationService.CreateConversationRecord(&conversationRecord); err != nil {
			global.GVA_LOG.Error("AI会话创建失败!", zap.Error(err))
			response.FailWithMessage("系统异常", c)
			return
		}
		conversationList = append(conversationList, conversation)
	}
	/**
	遍历成前端所需要的
	{
	    "title": "New Chat",
	    "uuid": 1683529241906,
	    "isEdit": false,
	    "conversationId": 579
	}
	*/
	var reConversationList []map[string]interface{}
	for _, conversation := range conversationList {
		reConversation := make(map[string]interface{})
		reConversation["title"] = conversation.ConversationName
		reConversation["uuid"] = conversation.ID
		reConversation["isEdit"] = false
		reConversation["conversationId"] = conversation.ID
		reConversationList = append(reConversationList, reConversation)
	}
	var chatList []map[string]interface{}
	for _, conversation := range conversationList {
		reConversation := make(map[string]interface{})
		reConversation["label"] = conversation.ConversationName
		reConversation["key"] = conversation.ID
		reConversation["icon"] = "BookOutline"
		chatList = append(chatList, reConversation)
	}
	conversationRecordList, err := conversationService.GetConversationRecordListByUserId(userInfo.BaseClaims.ID)
	if err != nil {
		return
	}
	// 创建conversation对话对象，两个属性 data和uuid
	var conversationDataGroups []*openfishReq.ConversationDataGroups
	for _, record := range conversationRecordList {
		found := false
		for _, group := range conversationDataGroups {
			if group.Uuid == *record.ConversationId {
				requestOption := make(map[string]interface{})
				requestOption["prompt"] = record.Content
				requestOption["options"] = nil
				conversationData := map[string]interface{}{
					"text":                record.Content,
					"dateTime":            record.CreatedAt.Format("2006-01-02 15:04:05"),
					"conversationOptions": nil,
					"error":               false,
					"inversion":           false,
					"loading":             false,
					"requestOptions":      requestOption,
				}
				if record.Role == "user" {
					conversationData["inversion"] = true
					requestOption["prompt"] = record.Content
				} else if record.Role == "assistant" {
					// TODO oldwei 计划是requestOption的prompt是不一样的，应该是user的text
					requestOption["prompt"] = record.Content
					conversationData["inversion"] = false
				}
				conversationData["requestOptions"] = requestOption
				group.Data = append(group.Data, conversationData)
				found = true
				break
			}
		}
		if !found {
			requestOption := make(map[string]interface{})
			requestOption["prompt"] = record.Content
			requestOption["options"] = nil
			conversationData := map[string]interface{}{
				"text":                record.Content,
				"dateTime":            record.CreatedAt.Format("2006-01-02 15:04:05"),
				"conversationOptions": nil,
				"error":               false,
				"inversion":           false,
				"loading":             false,
				"requestOptions":      requestOption,
			}
			if record.Role == "user" {
				requestOption["prompt"] = record.Content
			} else if record.Role == "assistant" {
				// TODO oldwei 计划是requestOption的prompt是不一样的，应该是user的text
				requestOption["prompt"] = record.Content
			}
			conversationData["requestOptions"] = requestOption
			conversationDataGroups = append(conversationDataGroups, &openfishReq.ConversationDataGroups{
				Data: []map[string]interface{}{
					conversationData,
				},
				Uuid: *record.ConversationId,
			})
		}
	}
	// 按年龄倒序排序
	sort.Slice(conversationDataGroups, func(i, j int) bool {
		return conversationDataGroups[i].Uuid > conversationDataGroups[j].Uuid
	})
	resultData := make(map[string]interface{})
	resultData["chat"] = conversationDataGroups
	resultData["history"] = reConversationList
	resultData["chatList"] = chatList
	resultData["usingContext"] = true
	resultData["active"] = reConversationList[0]["uuid"]
	resultData["activeConversationId"] = reConversationList[0]["uuid"]
	response.OkWithDetailed(resultData, "获取成功", c)
}
