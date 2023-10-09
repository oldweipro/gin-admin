package memo_nexus

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/memo_nexus/bilibili/request"
	"github.com/oldweipro/gin-admin/service"
	"time"
)

type MemoNexusApi struct {
}

var bilibiliService = service.ServiceGroupApp.MemoNexusService.BilibiliService

func (memoNexusApi *MemoNexusApi) GetLoginQrcodeGenerate(c *gin.Context) {
	qrcode := bilibiliService.GetLoginQrcodeGenerate()
	if qrcode.Url == "" {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(qrcode, c)
	}

}

func (memoNexusApi *MemoNexusApi) LoginQrcodePoll(c *gin.Context) {
	var bodyJson request.LoginQrcode
	err := c.ShouldBindJSON(&bodyJson)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 每秒钟请求一次接口 LoginQrcodePoll
	count := 0
	for {
		count++
		builtinMember := bilibiliService.LoginQrcodePoll(bodyJson.QrcodeKey)
		if builtinMember.Mid == 86038 || count > 60 {
			response.FailWithMessage("验证码已过期", c)
			return
		}
		if builtinMember.SessData != "" {
			// 存入数据库中
			bilibiliService.GetBuiltinMemberDetail(&builtinMember)
			if builtinMember.Mid == 0 {
				response.FailWithMessage("获取用户信息失败", c)
				return
			}
			err = bilibiliService.SaveBuiltinMember(builtinMember)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
			response.OkWithMessage("登陆成功", c)
			return
		}
		time.Sleep(time.Second * 1)
	}
}
