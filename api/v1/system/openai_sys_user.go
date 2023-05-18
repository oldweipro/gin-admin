package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/system"
	"github.com/oldweipro/gin-admin/model/system/request"
	systemRes "github.com/oldweipro/gin-admin/model/system/response"
	"github.com/oldweipro/gin-admin/utils"
	"github.com/oldweipro/gin-admin/utils/aliyun"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
	"time"
)

// OpenFishLogin
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/openfish_login [post]
func (b *BaseApi) OpenFishLogin(c *gin.Context) {
	var l request.OpenFishLogin
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || verify(l.Phone, l.SmsCode) {
		//查询是否存在 注册/登录
		sysUser, err := userService.FindUserByPhone(l.Phone)
		if err != nil {
			//注册
			var authorities []system.SysAuthority
			authorities = append(authorities, system.SysAuthority{
				AuthorityId: 9953,
			})
			u := &system.SysUser{Username: l.Phone, NickName: l.Phone, Password: "123456", HeaderImg: "", AuthorityId: 9953, Authorities: authorities, Enable: 1, Phone: l.Phone}
			*sysUser, err = userService.Register(*u)
			if err != nil {
				global.GVA_LOG.Error("注册失败!", zap.Error(err))
				response.FailWithDetailed(systemRes.SysUserResponse{User: *sysUser}, "注册失败", c)
				return
			}
		}

		if sysUser.Enable != 1 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		b.TokenNext(c, *sysUser)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// SmsCode 获取验证码
func (b *BaseApi) SmsCode(c *gin.Context) {
	var s request.SmsCode
	err := c.ShouldBindJSON(&s)

	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	rand.Seed(uint64(time.Now().UnixNano()))
	verificationCode := fmt.Sprintf("%06d", rand.Intn(1000000))
	global.Cache.Set(s.Mobile, verificationCode, time.Second*time.Duration(openCaptchaTimeOut))
	global.GVA_LOG.Info("打印verificationCode：" + verificationCode)
	// 发短信 阿里云依赖
	err = aliyun.GetSmsCode(s.Mobile, verificationCode)
	if err != nil {
		response.FailWithMessage("验证码获取失败", c)
		return
	}

	if err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysSmsCodeResponse{OpenCaptcha: oc}, "验证码获取成功", c)
}

func verify(username string, smsCode string) bool {
	code, ok := global.Cache.Get(username)
	if !ok {
		return false
	}
	str, err := code.(string)
	if !err {
		return false
	}
	if str == smsCode {
		global.Cache.Delete(username)
		return true
	}
	return false
}
