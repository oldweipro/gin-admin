package memo_nexus

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/memo_nexus/bilibili"
	"github.com/oldweipro/gin-admin/model/memo_nexus/bilibili/response"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"time"
)

type BilibiliService struct {
}

func (bilibiliService *BilibiliService) GetLoginQrcodeGenerate() response.LoginQrcodeGenerate {
	qrcodeGenerate := response.LoginQrcodeGenerate{}
	client := req.C().SetTimeout(5 * time.Second)
	var respData response.CommonResponse[response.LoginQrcodeGenerate]
	resp, err := client.R().
		SetSuccessResult(&respData).
		SetErrorResult(&respData).
		EnableDump().Get("https://passport.bilibili.com/x/passport-login/web/qrcode/generate")
	if err != nil {
		log.Println("error:", err)
		log.Println("raw content:")
		log.Println(resp.Dump())
		return qrcodeGenerate
	}

	if resp.IsErrorState() {
		fmt.Println(respData.Message)
		return qrcodeGenerate
	}

	if resp.IsSuccessState() {
		if respData.Code == 0 {
			return respData.Data
		}
	}

	log.Println("unknown status", resp.Status)
	log.Println("raw content:")
	log.Println(resp.Dump())
	return qrcodeGenerate
}

func (bilibiliService *BilibiliService) LoginQrcodePoll(qrcodeKey string) bilibili.BuiltinMember {
	builtinMember := bilibili.BuiltinMember{}
	client := req.C().SetTimeout(5 * time.Second)
	params := make(map[string]interface{})
	params["qrcode_key"] = qrcodeKey
	params["source"] = "main-fe-header"
	var respData response.CommonResponse[response.QrcodePoll]
	resp, err := client.R().
		SetQueryParamsAnyType(params).
		SetSuccessResult(&respData).
		SetErrorResult(&respData).
		EnableDump().Get("https://passport.bilibili.com/x/passport-login/web/qrcode/poll")
	if err != nil {
		log.Println("error:", err)
		log.Println("raw content:")
		log.Println(resp.Dump())
		return builtinMember
	}

	if resp.IsErrorState() {
		fmt.Println(respData.Message)
		return builtinMember
	}

	if resp.IsSuccessState() {
		if respData.Data.Code == 86038 {
			builtinMember.Mid = 86038
			return builtinMember
		}
		if respData.Code == 0 && respData.Data.Code == 0 {
			for _, cookie := range resp.Cookies() {
				switch cookie.Name {
				case "SESSDATA":
					builtinMember.SessData = cookie.Value
				case "DedeUserID":
					builtinMember.DedeUserID = cookie.Value
				case "DedeUserID__ckMd5":
					builtinMember.DedeUserIDCkMd5 = cookie.Value
				case "sid":
					builtinMember.Sid = cookie.Value
				case "bili_jct":
					builtinMember.BiliJct = cookie.Value
				}
			}
			builtinMember.RefreshToken = respData.Data.RefreshToken
		}
		return builtinMember
	}

	log.Println("unknown status", resp.Status)
	log.Println("raw content:")
	log.Println(resp.Dump())
	return builtinMember
}

func (bilibiliService *BilibiliService) SaveBuiltinMember(builtinMember bilibili.BuiltinMember) (err error) {
	err = global.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "mid"}},
		UpdateAll: true,
	}).Create(&builtinMember).Error
	return
}

func (bilibiliService *BilibiliService) GetBuiltinMemberDetail(builtinMember *bilibili.BuiltinMember) {
	client := req.C().SetTimeout(5 * time.Second)

	cookie := http.Cookie{Name: "SESSDATA", Value: builtinMember.SessData}
	client.SetCommonCookies(&cookie)

	var respData response.CommonResponse[response.ProfileData]
	resp, err := client.R().
		SetSuccessResult(&respData).
		SetErrorResult(&respData).
		EnableDump().Get("https://api.bilibili.com/x/space/v2/myinfo")
	if err != nil {
		log.Println("error:", err)
		log.Println("raw content:")
		log.Println(resp.Dump())
		return
	}

	if resp.IsErrorState() {
		fmt.Println(respData.Message)
		return
	}

	if resp.IsSuccessState() {
		builtinMember.Name = respData.Data.Profile.Name
		builtinMember.Mid = respData.Data.Profile.Mid
		return
	}

	log.Println("unknown status", resp.Status)
	log.Println("raw content:")
	log.Println(resp.Dump())
	return
}
