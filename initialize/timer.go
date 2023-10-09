package initialize

import (
	"fmt"
	"github.com/oldweipro/gin-admin/global"
	//v1 "github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/service"
)

func Timer() {
	//var certificationRecordApi = v1.ApiGroupApp.PatrolApiGroup.CertificationRecordApi
	//var personnelService = service.ServiceGroupApp.PatrolServiceGroup.PersonnelService
	var serverNodeService = service.ServiceGroupApp.LadderServiceGroup.ServerNodeService
	//var mailAccountService = service.ServiceGroupApp.OpenfishServiceGroup.MailAccountService
	if global.ConfigServer.Timer.Start {
		//for i := range global.ConfigServer.Timer.Detail {
		//	go func(detail config.Detail) {
		//		var option []cron.Option
		//		if global.ConfigServer.Timer.WithSeconds {
		//			option = append(option, cron.WithSeconds())
		//		}
		//		_, err := global.Timer.AddTaskByFunc("ClearDB", global.ConfigServer.Timer.Spec, func() {
		//			err := utils.ClearTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
		//			if err != nil {
		//				fmt.Println("timer error:", err)
		//			}
		//		}, option...)
		//		if err != nil {
		//			fmt.Println("add timer error:", err)
		//		}
		//	}(global.ConfigServer.Timer.Detail[i])
		//}
		// 每天重置被锁定的账号：account表中loginStatus重置为0
		//_, err := global.GVA_Timer.AddTaskByFunc("ResetGameAccount", "0 3 * * *", func() {
		//	// 重置所有account表中login_status=0，current_calls=0
		//	global.DB.Exec("UPDATE account SET login_status=0, current_calls=0 WHERE deleted_by=0")
		//	// 登陆其中一个账号
		//	//certificationRecordApi.LoginNewGameAccount()
		//	fmt.Println("每天重置被锁定的账号,定时任务执行完毕")
		//})
		//if err != nil {
		//	fmt.Println("添加每天重置被锁定的账号定时任务 error:", err)
		//}
		//
		//// 每天凌晨2点开始同步
		//_, err = global.GVA_Timer.AddTaskByFunc("SyncPersonnel", "0 2 * * *", func() {
		//	err = personnelService.SyncPersonnel()
		//	if err != nil {
		//		fmt.Println("定时任务同步人员数据失败")
		//	}
		//})
		//if err != nil {
		//	fmt.Println("添加每天同步人员数据定时任务 error:", err)
		//}
		//
		//// 每天凌晨3点开始同步
		//_, err = global.GVA_Timer.AddTaskByFunc("SyncPersonnel", "0 3 * * *", func() {
		//	err = personnelService.SyncPersonnelImg()
		//	if err != nil {
		//		fmt.Println("定时任务同步人员数据失败")
		//	}
		//})
		//if err != nil {
		//	fmt.Println("添加每天同步人员数据定时任务 error:", err)
		//}

		// 定时任务【同步梯子cookie】
		_, err := global.Timer.AddTaskByFunc("SyncLadderCookie", "0 0 1,5,10,15,20,25,28 * *", func() {
			err := serverNodeService.SyncLadderCookie()
			if err != nil {
				fmt.Println("定时任务【同步梯子cookie】失败")
			}
		})
		if err != nil {
			fmt.Println("添加每天【同步梯子cookie】定时任务 error:", err)
		}

		// 同步 OpenAI ChatGPT accessToken
		//_, err = global.Timer.AddTaskByFunc("SyncChatGPTAccessToken", global.ConfigServer.Timer.Spec, func() {
		//	go mailAccountService.SyncChatGPTAccessToken()
		//})
		//if err != nil {
		//	fmt.Println("添加同步 OpenAI ChatGPT accessToken 定时任务 error:", err)
		//}

		// 每小时修改所有账号AT状态
		//_, err = global.Timer.AddTaskByFunc("SyncChatGPTAccessTokenStatus", "@hourly", func() {
		//	go mailAccountService.SyncChatGPTAccessTokenStatus()
		//})
		//if err != nil {
		//	fmt.Println("添加 每小时修改所有账号AT状态 定时任务 error:", err)
		//}

		// 同步 Claude SK
		//_, err = global.Timer.AddTaskByFunc("SyncChatGPTAccessToken", "0 0 1,5,10,15,20,25,28 * *", func() {
		//	list, errGetMailAccountList := mailAccountService.GetMailAccountList()
		//	if errGetMailAccountList != nil {
		//		fmt.Println("同步 Claude SK 时，获取账户列表失败")
		//	}
		//	for _, account := range list {
		//		if account.ClaudeSessionKey != "" {
		//			var ids request.IdsReq
		//			ids.Ids = append(ids.Ids, int(account.ID))
		//			errRefreshOpenaiAccessToken := mailAccountService.RefreshClaudeChat(ids)
		//			if errRefreshOpenaiAccessToken != nil {
		//				fmt.Println("同步 Claude SK 失败")
		//			}
		//			time.Sleep(30 * time.Second)
		//		}
		//	}
		//})
		//if err != nil {
		//	fmt.Println("添加同步 Claude SK 定时任务 error:", err)
		//}
	}
}
