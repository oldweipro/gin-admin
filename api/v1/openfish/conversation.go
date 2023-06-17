package openfish

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
	"net/http"
	"sort"
	"strconv"
	"sync"
)

type ConversationApi struct {
}

var conversationService = service.ServiceGroupApp.OpenfishServiceGroup.ConversationService

var userRequestStatus sync.Map
var userRequestOpenAIDrawingStatus sync.Map

// OpenAIDrawing openai作图
func (conversationApi *ConversationApi) OpenAIDrawing(c *gin.Context) {
	// 获取用户ID
	userID := utils.GetUserID(c)
	// 检查用户的请求状态
	_, loaded := userRequestOpenAIDrawingStatus.LoadOrStore(userID, true)
	if loaded {
		c.JSON(429, gin.H{"msg": "太多请求了"})
		return
	}

	defer userRequestOpenAIDrawingStatus.Delete(userID) // 在处理完毕后删除用户的请求状态
	// 获取参数
	var chatReq openfishReq.ChatReq
	err := c.ShouldBindJSON(&chatReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 校验prompt是否为空
	if chatReq.Prompt == "" || chatReq.ConversationId == nil {
		response.FailWithMessage("系统错误，缺少必要参数。", c)
		global.Logger.Error("系统错误，缺少必要参数。")
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
	// AI 作画
	err = conversationService.OpenAIDrawing(chatReq, c)
	if err != nil {
		c.JSON(500, gin.H{"data": "服务器负载，请稍后重试。"})
	}
}

// ChatCompletions AI对话
func (conversationApi *ConversationApi) ChatCompletions(c *gin.Context) {
	// 获取用户ID
	userID := utils.GetUserID(c)
	// 检查用户的请求状态
	_, loaded := userRequestStatus.LoadOrStore(userID, true)
	if loaded {
		c.JSON(429, gin.H{"msg": "太多请求了"})
		return
	}

	defer userRequestStatus.Delete(userID) // 在处理完毕后删除用户的请求状态
	// 获取参数
	var chatReq openfishReq.ChatReq
	err := c.ShouldBindJSON(&chatReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 校验prompt是否为空
	if chatReq.Prompt == "" || chatReq.ConversationId == nil {
		response.FailWithMessage("系统错误，缺少必要参数。", c)
		global.Logger.Error("系统错误，缺少必要参数。")
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
	err = conversationService.ChatGPTCompletions(chatReq, c)
	if err != nil {
		c.JSON(500, gin.H{"data": "服务器负载，请稍后重试。"})
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
		global.Logger.Error("创建会话失败!", zap.Error(err))
		response.FailWithMessage("系统异常", c)
		return
	}
	// conversationId等于空则创建该信息到数据库
	conversationRecord := openfish.ConversationRecord{}

	if conversation.ConversationType == 0 {
		conversationRecord.Content = "我是由开放鱼训练的大型语言模型，请详细描述您的问题。"
	} else if conversation.ConversationType == 1 {
		conversationRecord.Content = "我是由开放鱼训练的图像生成模型，请详细描述您的图画。"
	}
	conversationRecord.Role = "system"
	conversationRecord.ConversationId = &conversation.ID
	conversationRecord.CreatedBy = utils.GetUserID(c)
	if err := conversationService.CreateConversationRecord(&conversationRecord); err != nil {
		global.Logger.Error("AI会话创建失败!", zap.Error(err))
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
		global.Logger.Error("删除失败!", zap.Error(err))
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
		global.Logger.Error("批量删除失败!", zap.Error(err))
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
		global.Logger.Error("更新失败!", zap.Error(err))
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
		global.Logger.Error("更新失败!", zap.Error(err))
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
		global.Logger.Error("查询失败!", zap.Error(err))
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
		global.Logger.Error("获取失败!", zap.Error(err))
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
	var chatReq openfishReq.ChatReq
	err := c.ShouldBindQuery(&chatReq)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if chatReq.ConversationType == nil {
		response.FailWithMessage("conversationType不能为空", c)
		return
	}
	userInfo := utils.GetUserInfo(c)
	conversationList, err := conversationService.GetConversationListByUserId(userInfo.BaseClaims.ID, *chatReq.ConversationType)

	if len(conversationList) == 0 {
		// 创建一个聊天
		var conversation openfish.Conversation
		conversation.ConversationName = "新聊天"
		conversation.ConversationType = *chatReq.ConversationType
		conversation.CreatedBy = utils.GetUserID(c)
		if err := conversationService.CreateConversation(&conversation); err != nil {
			global.Logger.Error("创建会话失败!", zap.Error(err))
			response.FailWithMessage("系统异常", c)
			return
		}
		// conversationId等于空则创建该信息到数据库
		conversationRecord := openfish.ConversationRecord{}
		if *chatReq.ConversationType == 0 {
			conversationRecord.Content = "我是由开放鱼训练的大型语言模型，请详细描述您的问题。"
		} else if *chatReq.ConversationType == 1 {
			conversationRecord.Content = "我是由开放鱼训练的图像生成模型，请详细描述您的图画。"
		}
		conversationRecord.Role = "system"
		conversationRecord.ConversationId = &conversation.ID
		conversationRecord.CreatedBy = utils.GetUserID(c)
		if err := conversationService.CreateConversationRecord(&conversationRecord); err != nil {
			global.Logger.Error("AI会话创建失败!", zap.Error(err))
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
	//conversationRecordList, err := conversationService.GetConversationRecordListByUserId(userInfo.BaseClaims.ID)
	conversationRecordList, err := conversationService.GetConversationRecordListByConversationId(conversationList[0].ID)
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
	// 按时间倒序排序
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
