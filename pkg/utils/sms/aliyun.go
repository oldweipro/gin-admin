package sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/oldweipro/gin-admin/pkg/app"
)

func CreateClient() (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(app.Config.AliyunOSS.AccessKeyId),
		AccessKeySecret: tea.String(app.Config.AliyunOSS.AccessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func GetSmsCode(mobile string, code string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(mobile),
		SignName:      tea.String(app.Config.AliyunOSS.SignName),
		TemplateCode:  tea.String(app.Config.AliyunOSS.TemplateCode),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		str, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}
		app.Logger.Info("aliyun: " + str.String())
		return nil
	}()

	if tryErr != nil {
		var e = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			e = _t
		} else {
			e.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 e
		_, _err = util.AssertAsString(e.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
