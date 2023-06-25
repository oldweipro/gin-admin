package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
	"strings"
)

type SecretKeyApi struct {
}

var secretKeyService = service.ServiceGroupApp.OpenfishServiceGroup.SecretKeyService

// CreateSecretKey 创建SecretKey
// @Tags SecretKey
// @Summary 创建SecretKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.SecretKey true "创建SecretKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /secretKey/createSecretKey [post]
func (secretKeyApi *SecretKeyApi) CreateSecretKey(c *gin.Context) {
	var secretKey openfish.SecretKey
	err := c.ShouldBindJSON(&secretKey)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	secretKey.CreatedBy = utils.GetUserID(c)
	secretKey.UserId = utils.GetUserID(c)
	secretKey.Sk = "sk-" + strings.ReplaceAll(uuid.NewString(), "-", "") + strings.ReplaceAll(uuid.NewString(), "-", "")[:16]
	if err := secretKeyService.CreateSecretKey(&secretKey); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSecretKey 删除SecretKey
// @Tags SecretKey
// @Summary 删除SecretKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.SecretKey true "删除SecretKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /secretKey/deleteSecretKey [delete]
func (secretKeyApi *SecretKeyApi) DeleteSecretKey(c *gin.Context) {
	var secretKey openfish.SecretKey
	err := c.ShouldBindJSON(&secretKey)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	secretKey.DeletedBy = utils.GetUserID(c)
	if err := secretKeyService.DeleteSecretKey(secretKey); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSecretKeyByIds 批量删除SecretKey
// @Tags SecretKey
// @Summary 批量删除SecretKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SecretKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /secretKey/deleteSecretKeyByIds [delete]
func (secretKeyApi *SecretKeyApi) DeleteSecretKeyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := secretKeyService.DeleteSecretKeyByIds(IDS, deletedBy); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSecretKey 更新SecretKey
// @Tags SecretKey
// @Summary 更新SecretKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.SecretKey true "更新SecretKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /secretKey/updateSecretKey [put]
func (secretKeyApi *SecretKeyApi) UpdateSecretKey(c *gin.Context) {
	var secretKey openfish.SecretKey
	err := c.ShouldBindJSON(&secretKey)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	secretKey.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Sk":     {utils.NotEmpty()},
		"SkName": {utils.NotEmpty()},
		"UserId": {utils.NotEmpty()},
		"Expire": {utils.NotEmpty()},
		"Amount": {utils.NotEmpty()},
	}
	if err := utils.Verify(secretKey, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := secretKeyService.UpdateSecretKey(secretKey); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSecretKey 用id查询SecretKey
// @Tags SecretKey
// @Summary 用id查询SecretKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfish.SecretKey true "用id查询SecretKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /secretKey/findSecretKey [get]
func (secretKeyApi *SecretKeyApi) FindSecretKey(c *gin.Context) {
	var secretKey openfish.SecretKey
	err := c.ShouldBindQuery(&secretKey)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resecretKey, err := secretKeyService.GetSecretKey(secretKey.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resecretKey": resecretKey}, c)
	}
}

// GetSecretKeyList 分页获取SecretKey列表
// @Tags SecretKey
// @Summary 分页获取SecretKey列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfishReq.SecretKeySearch true "分页获取SecretKey列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /secretKey/getSecretKeyList [get]
func (secretKeyApi *SecretKeyApi) GetSecretKeyList(c *gin.Context) {
	var pageInfo openfishReq.SecretKeySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = utils.GetUserID(c)
	if list, total, err := secretKeyService.GetSecretKeyInfoLessList(pageInfo); err != nil {
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
