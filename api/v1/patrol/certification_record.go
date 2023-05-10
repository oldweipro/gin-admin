package patrol

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CertificationRecordApi struct {
}

var certificationRecordService = service.ServiceGroupApp.PatrolServiceGroup.CertificationRecordService

// CreateCertificationRecord 创建CertificationRecord
// @Tags CertificationRecord
// @Summary 创建CertificationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.CertificationRecord true "创建CertificationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /certificationRecord/createCertificationRecord [post]
func (certificationRecordApi *CertificationRecordApi) CreateCertificationRecord(c *gin.Context) {
	var certificationRecord patrol.CertificationRecord
	err := c.ShouldBindJSON(&certificationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	certificationRecord.CreatedBy = 1
	verify := utils.Rules{
		"Certification_id_card":   {utils.NotEmpty()},
		"Certification_real_name": {utils.NotEmpty()},
	}
	if err := utils.Verify(certificationRecord, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 正则校验身份证号
	idRe18 := `^([1-6][1-9]|50)\d{4}(18|19|20)\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`
	idRe15 := `^([1-6][1-9]|50)\d{4}\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\d{3}$`
	if regexp.MustCompile(idRe18).MatchString(certificationRecord.CertificationIdCard) || regexp.MustCompile(idRe15).MatchString(certificationRecord.CertificationIdCard) {
		global.GVA_LOG.Info("身份证号码正则校验通过✅")
	} else {
		global.GVA_LOG.Error("身份证号码正则校验失败")
		response.FailWithMessage("请输入正确的身份证号", c)
		return
	}
	// 应该先从数据库查询，查询没有之后才去调接口查询
	search := patrolReq.IdentitySearch{}
	search.RealName = certificationRecord.CertificationRealName
	search.IdCard = certificationRecord.CertificationIdCard
	infoList, total, _ := identityService.GetIdentityInfoList(search)
	if total > 0 {
		fmt.Println(infoList[0])
		response.OkWithMessage("身份证号码和真实姓名一致", c)
		return
	} else {
		// 如何轮换游戏的账号信息: 查出来所有的符合规则的账号（账号调用10次大约就用完了，所以做一个计数吧，每天12点清0）
		accountSearch := patrolReq.AccountSearch{}
		loginStatus := 1
		// 查询出来调用小于10的？目前大约调用10次就会出现账号锁定的情况，就得换号了
		currentCalls := 10
		accountSearch.LoginStatus = &loginStatus
		accountSearch.CurrentCalls = &currentCalls
		// 获取可用的账号
		list, total, _ := accountService.GetAccountInfoList(accountSearch)
		var account patrol.Account
		if total > 0 {
			account = list[0]
		} else {
			// 既然没有账号可用了，那我就在这个地方给他登陆一个账号
			var b bool
			gameAccount, b := certificationRecordApi.LoginNewGameAccount()
			if b {
				account = gameAccount
			} else {
				certificationRecord.CertificationMsg = "没有账号了"
				// 没有账号可用了
				global.GVA_LOG.Error("没有账号可用了!")
			}
		}
		global.GVA_LOG.Info("登陆的账号" + account.AccountName)
		// 发起认证
		m, err := certificationRecordApi.CertificationGameApi(certificationRecord.CertificationRealName, certificationRecord.CertificationIdCard, account)
		if err != nil {
			response.FailWithMessage("服务器发运行生错误", c)
			return
		}
		// 调用完接口了，所以调用次数+1吧
		currentCallsCompute := *account.CurrentCalls + 1
		account.CurrentCalls = &currentCallsCompute
		accountService.UpdateAccount(account)
		msg := m["msg"].(string)
		code := int(m["status"].(float64))
		status := m["status"].(float64)
		certificationRecord.CertificationCode = &code
		certificationRecord.CertificationMsg = msg
		jsonStr, _ := json.Marshal(m)
		certificationRecord.CertificationResult = string(jsonStr)
		global.GVA_LOG.Info(string(jsonStr))
		if status == 1 {
			var identity patrol.Identity
			identity.CreatedBy = utils.GetUserID(c)
			identity.RealName = certificationRecord.CertificationRealName
			identity.IdCard = certificationRecord.CertificationIdCard
			if err := identityService.CreateIdentity(identity); err != nil {
				global.GVA_LOG.Error("创建失败!", zap.Error(err))
			} else {
				global.GVA_LOG.Info("实名信息创建成功!")
			}
			response.OkWithMessage(msg, c)
		} else {
			response.FailWithMessage(msg, c)
		}
	}
	if err := certificationRecordService.CreateCertificationRecord(certificationRecord); err != nil {
		global.GVA_LOG.Error("数据库存储认证记录失败!", zap.Error(err))
	}
}

func (certificationRecordApi *CertificationRecordApi) CertificationGameApi(realName, idCard string, account patrol.Account) (map[string]interface{}, error) {
	// ============== 向游戏认证接口发起认证请求 ===================
	gameId := "6293"
	signKey := "a40874cdf8b253202921b25b4cd87d84"
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	sum := md5.Sum([]byte(gameId + idCard + realName + timeNow + "#" + signKey))
	sign := fmt.Sprintf("%x", sum)
	params := "realname=" + realName + "&id_card=" + idCard + "&game_id=" + gameId + "&username=" + account.AccountName + "&time=" + timeNow + "&sign=" + sign
	resp, err := http.Get("" + params)
	defer resp.Body.Close()
	if err != nil {
		global.GVA_LOG.Error("发起请求失败!", zap.Error(err))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("读取返回结果失败!", zap.Error(err))
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		global.GVA_LOG.Error("结果解析错误!", zap.Error(err))
	}
	msg := m["msg"].(string)
	if strings.Contains(msg, "已锁定") {
		// 账号锁定了，把loginStatus状态码设置为2
		loginStatusLock := 2
		account.LoginStatus = &loginStatusLock
		account.UpdatedBy = 1
		accountService.UpdateAccount(account)
		global.GVA_LOG.Error("账号锁定了，把loginStatus状态码设置为2!")
		// 当前账号被🔒锁定了，需要登陆一下新的账号了，方便下一次接口调用时可以提供服务
		gameAccount, b := certificationRecordApi.LoginNewGameAccount()
		if b {
			// TODO 应该再次发起一下请求，否则页面上就是报错了，调10下就一个报错，这个频率有点太高了，得做的再完美一些
			// 递归一下
			return certificationRecordApi.CertificationGameApi(realName, idCard, gameAccount)
		}
	} else if strings.Contains(msg, "time_expire_login") {
		// 游戏账号登陆过期了，只需要重新登陆，不需要登陆新账号
		loginStatus := 1
		account.LoginStatus = &loginStatus
		gameAccount, _ := accountService.LoginGameAccount(account)
		if gameAccount == "<script type=\"text/javascript\">top.location=\"http:\\/\\/www.9377.com\";</script>" {
			global.GVA_LOG.Info("登陆已过期的游戏账号成功！")
			// 递归一下
			return certificationRecordApi.CertificationGameApi(realName, idCard, account)
		} else {
			global.GVA_LOG.Error("登陆失败!页面返回错误")
		}
	}
	return m, err
}

func (certificationRecordApi *CertificationRecordApi) LoginNewGameAccount() (patrol.Account, bool) {
	accountUnLoginSearch := patrolReq.AccountSearch{}
	unLoginStatus := 0
	accountUnLoginSearch.LoginStatus = &unLoginStatus
	currentCalls := 10
	accountUnLoginSearch.CurrentCalls = &currentCalls
	// 获取可用的账号
	list, total, _ := accountService.GetAccountInfoList(accountUnLoginSearch)
	if total > 0 {
		account := list[0]
		//这个修改loginStatus有问题
		loginStatus := 1
		account.LoginStatus = &loginStatus
		gameAccount, _ := accountService.LoginGameAccount(account)
		if strings.Contains(gameAccount, "<script type=\"text/javascript\">top.location=\"http:") {
			global.GVA_LOG.Info("登陆原来可用的游戏账号成功！")
			return account, true
		} else {
			global.GVA_LOG.Error("登陆失败!页面返回错误")
			return account, false
		}
	} else {
		global.GVA_LOG.Error("😭😭😭是真的没有账号可用了!注册新账号吧！并且给自动登陆上")
		// 没有的话那就给他再造些
		s := time.Now().Unix()
		username := "atx" + strconv.FormatInt(s, 10)
		if certificationRecordApi.RegisterMember(username) {
			var accountRegister patrol.Account
			accountRegister.AccountName = username
			// 这个接口改变loginStatus没问题
			loginStatus := 1
			accountRegister.LoginStatus = &loginStatus
			gameAccount, _ := accountService.LoginGameAccount(accountRegister)
			if strings.Contains(gameAccount, "<script type=\"text/javascript\">top.location=\"http:") {
				global.GVA_LOG.Info("登陆新注册的游戏账号成功！")
				accountResult := patrolReq.AccountSearch{}
				accountResult.AccountName = username
				// 获取可用的账号
				listAccount, totals, _ := accountService.GetAccountInfoList(accountResult)
				if totals > 0 {
					return listAccount[0], true
				} else {
					return patrol.Account{}, false
				}
			} else {
				global.GVA_LOG.Error("登陆失败!页面返回错误")
				return patrol.Account{}, false
			}
		}
		global.GVA_LOG.Error("注册账号发生错误🙅❌!")
		return patrol.Account{}, false
	}
}

func (certificationRecordApi *CertificationRecordApi) RegisterMember(username string) bool {
	req, err := http.NewRequest("GET", ""+
		"&userid="+username+"&NAME=&password="+username+"&ID_CARD_NUMBER=&js_callback=requestCallback", nil)
	if err != nil {
		// handle err
		global.GVA_LOG.Error("请求报错了!")
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "")
	req.Header.Set("Sec-Fetch-Dest", "script")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		global.GVA_LOG.Error("注册失败了!")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	v := string(body)
	global.GVA_LOG.Info("注册游戏账号结果！" + v)
	if strings.Contains(v, "game_login.php?game=wz&server=1917&username=atx") {
		global.GVA_LOG.Info("注册成功")
		return true
	}
	return false
}

// DeleteCertificationRecord 删除CertificationRecord
// @Tags CertificationRecord
// @Summary 删除CertificationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.CertificationRecord true "删除CertificationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /certificationRecord/deleteCertificationRecord [delete]
func (certificationRecordApi *CertificationRecordApi) DeleteCertificationRecord(c *gin.Context) {
	var certificationRecord patrol.CertificationRecord
	err := c.ShouldBindJSON(&certificationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	certificationRecord.DeletedBy = utils.GetUserID(c)
	if err := certificationRecordService.DeleteCertificationRecord(certificationRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCertificationRecordByIds 批量删除CertificationRecord
// @Tags CertificationRecord
// @Summary 批量删除CertificationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CertificationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /certificationRecord/deleteCertificationRecordByIds [delete]
func (certificationRecordApi *CertificationRecordApi) DeleteCertificationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := certificationRecordService.DeleteCertificationRecordByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCertificationRecord 更新CertificationRecord
// @Tags CertificationRecord
// @Summary 更新CertificationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.CertificationRecord true "更新CertificationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /certificationRecord/updateCertificationRecord [put]
func (certificationRecordApi *CertificationRecordApi) UpdateCertificationRecord(c *gin.Context) {
	var certificationRecord patrol.CertificationRecord
	err := c.ShouldBindJSON(&certificationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	certificationRecord.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Certification_id_card":   {utils.NotEmpty()},
		"Certification_real_name": {utils.NotEmpty()},
	}
	if err := utils.Verify(certificationRecord, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := certificationRecordService.UpdateCertificationRecord(certificationRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCertificationRecord 用id查询CertificationRecord
// @Tags CertificationRecord
// @Summary 用id查询CertificationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrol.CertificationRecord true "用id查询CertificationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /certificationRecord/findCertificationRecord [get]
func (certificationRecordApi *CertificationRecordApi) FindCertificationRecord(c *gin.Context) {
	var certificationRecord patrol.CertificationRecord
	err := c.ShouldBindQuery(&certificationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recertificationRecord, err := certificationRecordService.GetCertificationRecord(certificationRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recertificationRecord": recertificationRecord}, c)
	}
}

// GetCertificationRecordList 分页获取CertificationRecord列表
// @Tags CertificationRecord
// @Summary 分页获取CertificationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrolReq.CertificationRecordSearch true "分页获取CertificationRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /certificationRecord/getCertificationRecordList [get]
func (certificationRecordApi *CertificationRecordApi) GetCertificationRecordList(c *gin.Context) {
	var pageInfo patrolReq.CertificationRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := certificationRecordService.GetCertificationRecordInfoList(pageInfo); err != nil {
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
