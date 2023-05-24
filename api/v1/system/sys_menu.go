package system

import (
	"encoding/json"
	"fmt"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/system"
	systemReq "github.com/oldweipro/gin-admin/model/system/request"
	systemRes "github.com/oldweipro/gin-admin/model/system/response"
	"github.com/oldweipro/gin-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct{}

// GetMenu
// @Tags      AuthorityMenu
// @Summary   获取用户动态路由
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      request.Empty                                                  true  "空"
// @Success   200   {object}  response.Response{data=systemRes.SysMenusResponse,msg=string}  "获取用户动态路由,返回包括系统菜单详情列表"
// @Router    /menu/getMenu [post]
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if menus == nil {
		menus = []system.SysMenu{}
	}
	response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
}

func (a *AuthorityMenuApi) GetMenus(c *gin.Context) {
	menusStr := `
[
  {
    "path": "/dashboard",
    "name": "Dashboard",
    "component": "LAYOUT",
    "redirect": "/dashboard/console",
    "meta": {
      "icon": "DashboardOutlined",
      "title": "仪表盘",
      "sort": 2
    },
    "children": [
      {
        "path": "console",
        "name": "dashboard_console",
        "component": "/dashboard/console/console",
        "meta": {
          "title": "主控台"
        }
      },
      {
        "path": "monitor",
        "name": "dashboard_monitor",
        "component": "/dashboard/monitor/monitor",
        "meta": {
          "title": "监控页"
        }
      },
      {
        "path": "workplace",
        "name": "dashboard_workplace",
        "component": "/dashboard/workplace/workplace",
        "meta": {
          "hidden": true,
          "title": "工作台"
        }
      }
    ]
  }, {
    "path": "/system",
    "name": "System",
    "component": "LAYOUT",
    "redirect": "/system/menu",
    "meta": {
      "title": "系统设置",
      "icon": "DashboardOutlined",
      "sort": 1
    },
    "children": [
      {
        "path": "menu",
        "name": "system_menu",
        "component": "/system/menu/menu",
        "meta": {
          "title": "菜单权限管理"
        }
      },
      {
        "path": "role",
        "name": "system_role",
        "component": "/system/role/role",
        "meta": {
          "title": "角色权限管理"
        }
      }
    ]
  },{
    "path": "/chat",
    "name": "Chat",
    "component": "LAYOUT",
    "redirect": "/chat/chatgpt",
    "meta": {
      "icon": "DashboardOutlined",
      "title": "ChatGPT",
      "isRoot": true,
      "sort": 2
    },
    "children": [
      {
        "path": "chatgpt",
        "name": "chat_chatgpt",
        "component": "/chat/chatgpt/chatgpt",
        "meta": {
          "title": "ChatGPT"
        }
      }
    ]
  }
]`

	var menus []map[string]interface{}
	err := json.Unmarshal([]byte(menusStr), &menus)
	fmt.Println(err)
	response.OkWithDetailed(menus, "获取成功", c)
}

func (a *AuthorityMenuApi) GetMenusList(c *gin.Context) {
	menusStr := `[
    {
      "label": "Dashboard",
      "key": "dashboard",
      "type": 1,
      "subtitle": "dashboard",
      "openType": 1,
      "auth": "dashboard",
      "path": "/dashboard",
      "children": [
        {
          "label": "主控台",
          "key": "console",
          "type": 1,
          "subtitle": "console",
          "openType": 1,
          "auth": "console",
          "path": "/dashboard/console"
        },
        {
          "label": "工作台",
          "key": "workplace",
          "type": 1,
          "subtitle": "workplace",
          "openType": 1,
          "auth": "workplace",
          "path": "/dashboard/workplace"
        }
      ]
    },
    {
      "label": "表单管理",
      "key": "form",
      "type": 1,
      "subtitle": "form",
      "openType": 1,
      "auth": "form",
      "path": "/form",
      "children": [
        {
          "label": "基础表单",
          "key": "basic-form",
          "type": 1,
          "subtitle": "basic-form",
          "openType": 1,
          "auth": "basic-form",
          "path": "/form/basic-form"
        },
        {
          "label": "分步表单",
          "key": "step-form",
          "type": 1,
          "subtitle": "step-form",
          "openType": 1,
          "auth": "step-form",
          "path": "/form/step-form"
        },
        {
          "label": "表单详情",
          "key": "detail",
          "type": 1,
          "subtitle": "detail",
          "openType": 1,
          "auth": "detail",
          "path": "/form/detail"
        }
      ]
    }
  ]`
	var menus []map[string]interface{}
	err := json.Unmarshal([]byte(menusStr), &menus)
	fmt.Println(err)
	response.OkWithDetailed(menus, "获取成功", c)
}

// GetBaseMenuTree
// @Tags      AuthorityMenu
// @Summary   获取用户动态路由
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      request.Empty                                                      true  "空"
// @Success   200   {object}  response.Response{data=systemRes.SysBaseMenusResponse,msg=string}  "获取用户动态路由,返回包括系统菜单列表"
// @Router    /menu/getBaseMenuTree [post]
func (a *AuthorityMenuApi) GetBaseMenuTree(c *gin.Context) {
	menus, err := menuService.GetBaseMenuTree()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
}

// AddMenuAuthority
// @Tags      AuthorityMenu
// @Summary   增加menu和角色关联关系
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.AddMenuAuthorityInfo  true  "角色ID"
// @Success   200   {object}  response.Response{msg=string}   "增加menu和角色关联关系"
// @Router    /menu/addMenuAuthority [post]
func (a *AuthorityMenuApi) AddMenuAuthority(c *gin.Context) {
	var authorityMenu systemReq.AddMenuAuthorityInfo
	err := c.ShouldBindJSON(&authorityMenu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// GetMenuAuthority
// @Tags      AuthorityMenu
// @Summary   获取指定角色menu
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetAuthorityId                                     true  "角色ID"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "获取指定角色menu"
// @Router    /menu/getMenuAuthority [post]
func (a *AuthorityMenuApi) GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityId
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(param, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menus, err := menuService.GetMenuAuthority(&param)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
}

// AddBaseMenu
// @Tags      Menu
// @Summary   新增菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysBaseMenu             true  "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success   200   {object}  response.Response{msg=string}  "新增菜单"
// @Router    /menu/addBaseMenu [post]
func (a *AuthorityMenuApi) AddBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu.Meta, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = menuService.AddBaseMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// DeleteBaseMenu
// @Tags      Menu
// @Summary   删除菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "菜单id"
// @Success   200   {object}  response.Response{msg=string}  "删除菜单"
// @Router    /menu/deleteBaseMenu [post]
func (a *AuthorityMenuApi) DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = baseMenuService.DeleteBaseMenu(menu.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateBaseMenu
// @Tags      Menu
// @Summary   更新菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysBaseMenu             true  "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success   200   {object}  response.Response{msg=string}  "更新菜单"
// @Router    /menu/updateBaseMenu [post]
func (a *AuthorityMenuApi) UpdateBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu.Meta, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = baseMenuService.UpdateBaseMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetBaseMenuById
// @Tags      Menu
// @Summary   根据id获取菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                                                   true  "菜单id"
// @Success   200   {object}  response.Response{data=systemRes.SysBaseMenuResponse,msg=string}  "根据id获取菜单,返回包括系统菜单列表"
// @Router    /menu/getBaseMenuById [post]
func (a *AuthorityMenuApi) GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menu, err := baseMenuService.GetBaseMenuById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
}

// GetMenuList
// @Tags      Menu
// @Summary   分页获取基础menu列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取基础menu列表,返回包括列表,总数,页码,每页数量"
// @Router    /menu/getMenuList [post]
func (a *AuthorityMenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menuList, total, err := menuService.GetInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menuList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
