package transaction

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/utils"
	"sync"
	"time"
)

type ActivationCodeApi struct {
}

var getJetBrainsActivationCodeStatus sync.Map

// GetActivationCodeStatus 获取订阅状态
func (activationCodeApi *ActivationCodeApi) GetActivationCodeStatus(c *gin.Context) {
	// 先查询你的状态是否激活
	user, subErr := subscriptionPlanService.GetCurrentSubscriptionPlan(utils.GetUserID(c), 2)
	if subErr != nil {
		response.FailWithMessage("请开通您的订阅计划", c)
		return
	}
	if *user.Status == 0 {
		response.FailWithMessage("请选择您的订阅计划", c)
		return
	}
	// 获取当前时间
	currentTime := time.Now()
	// 创建一个要比较的时间
	compareTime := user.EndTime
	// 比较两个时间
	if compareTime.Before(currentTime) {
		response.FailWithMessage("您的订阅已过期，请重新选择订阅计划", c)
		return
	}
	response.OkWithMessage("请尽情享受!", c)
}

// GetJetBrainsActivationCode 获取激活码
func (activationCodeApi *ActivationCodeApi) GetJetBrainsActivationCode(c *gin.Context) {
	userId := utils.GetUserID(c)
	_, loaded := getJetBrainsActivationCodeStatus.LoadOrStore(userId, true)
	defer getJetBrainsActivationCodeStatus.Delete(userId)
	if loaded {
		response.FailStatusTooManyRequestsWithDetailed(nil, "请求过多", c)
		return
	}
	targetUrl := "https://vrg123.com"
	collyController := colly.NewCollector()
	collyController.OnRequest(func(r *colly.Request) {
		fmt.Println("Getting activation code...")
	})
	resultText := ""
	collyController.OnHTML("textarea", func(e *colly.HTMLElement) {
		resultText = e.Text
	})
	err := collyController.Visit(targetUrl)
	if err != nil {
		fmt.Println("Whoops, something went wrong!")
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(resultText, "订阅成功", c)
}
