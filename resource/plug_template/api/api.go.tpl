package api

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/response"
{{ if .NeedModel }}	"github.com/oldweipro/gin-admin/plugin/{{ .Snake}}/model" {{ end }}
	"github.com/oldweipro/gin-admin/plugin/{{ .Snake}}/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type {{ .PlugName}}Api struct{}

// @Tags {{ .PlugName}}
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /{{ .RouterGroup}}/routerName [post]
func (p *{{ .PlugName}}Api) ApiName(c *gin.Context) {
    {{ if .HasRequest}}
        var plug model.Request
        _ = c.ShouldBindJSON(&plug)
    {{ end }}
        if {{ if .HasResponse }} res, {{ end }} err:= service.ServiceGroupApp.PlugService({{ if .HasRequest }}plug{{ end -}}); err != nil {
		global.Logger.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
	{{if .HasResponse }}
	    response.OkWithDetailed(res,"成功",c)
	{{else}}
	    response.OkWithData("成功", c)
	{{ end -}}

	}
}
