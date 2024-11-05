package system

import (
	"errors"
	"github.com/oldweipro/gin-admin/model/system"
	"github.com/oldweipro/gin-admin/model/system/request"
	"github.com/oldweipro/gin-admin/model/system/response"
	"github.com/oldweipro/gin-admin/pkg/app"
	"gorm.io/gorm"
)

type AuthorityBtnService struct{}

var AuthorityBtnServiceApp = new(AuthorityBtnService)

func (a *AuthorityBtnService) GetAuthorityBtn(req request.SysAuthorityBtnReq) (res response.SysAuthorityBtnRes, err error) {
	var authorityBtn []system.SysAuthorityBtn
	err = app.DBClient.Find(&authorityBtn, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
	if err != nil {
		return
	}
	var selected []uint
	for _, v := range authorityBtn {
		selected = append(selected, v.SysBaseMenuBtnID)
	}
	res.Selected = selected
	return res, err
}

func (a *AuthorityBtnService) SetAuthorityBtn(req request.SysAuthorityBtnReq) (err error) {
	return app.DBClient.Transaction(func(tx *gorm.DB) error {
		var authorityBtn []system.SysAuthorityBtn
		err = tx.Delete(&[]system.SysAuthorityBtn{}, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
		if err != nil {
			return err
		}
		for _, v := range req.Selected {
			authorityBtn = append(authorityBtn, system.SysAuthorityBtn{
				AuthorityId:      req.AuthorityId,
				SysMenuID:        req.MenuID,
				SysBaseMenuBtnID: v,
			})
		}
		if len(authorityBtn) > 0 {
			err = tx.Create(&authorityBtn).Error
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (a *AuthorityBtnService) CanRemoveAuthorityBtn(ID string) (err error) {
	fErr := app.DBClient.First(&system.SysAuthorityBtn{}, "sys_base_menu_btn_id = ?", ID).Error
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("此按钮正在被使用无法删除")
}
