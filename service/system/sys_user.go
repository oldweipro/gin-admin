package system

import (
	"errors"
	"fmt"
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/pkg/utils"
	"gorm.io/datatypes"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/system"
	"gorm.io/gorm"
)

//@author: [oldweipro](https://github.com/oldweipro)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter system.SysUser, err error

type UserService struct{}

var UserServiceApp = new(UserService)

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(app.DBClient.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = app.DBClient.Create(&u).Error
	return u, err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == app.DBClient {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = app.DBClient.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = app.DBClient.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = app.DBClient.Save(&user).Error
	return &user, err

}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := app.DBClient.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {

	assignErr := app.DBClient.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}

	var authority system.SysAuthority
	err = app.DBClient.Where("authority_id = ?", authorityId).First(&authority).Error
	if err != nil {
		return err
	}
	var authorityMenu []system.SysAuthorityMenu
	var authorityMenuIDs []string
	err = app.DBClient.Where("sys_authority_authority_id = ?", authorityId).Find(&authorityMenu).Error
	if err != nil {
		return err
	}

	for i := range authorityMenu {
		authorityMenuIDs = append(authorityMenuIDs, authorityMenu[i].MenuId)
	}

	var authorityMenus []system.SysBaseMenu
	err = app.DBClient.Preload("Parameters").Where("id in (?)", authorityMenuIDs).Find(&authorityMenus).Error
	if err != nil {
		return err
	}
	hasMenu := false
	for i := range authorityMenus {
		if authorityMenus[i].Name == authority.DefaultRouter {
			hasMenu = true
			break
		}
	}
	if !hasMenu {
		return errors.New("找不到默认路由,无法切换本角色")
	}

	err = app.DBClient.Model(&system.SysUser{}).Where("id = ?", id).Update("authority_id", authorityId).Error
	return err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(adminAuthorityID, id uint, authorityIds []uint) (err error) {
	return app.DBClient.Transaction(func(tx *gorm.DB) error {
		var user system.SysUser
		TxErr := tx.Where("id = ?", id).First(&user).Error
		if TxErr != nil {
			app.Logger.Debug(TxErr.Error())
			return errors.New("查询用户数据失败")
		}
		TxErr = tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			e := AuthorityServiceApp.CheckAuthorityIDAuth(adminAuthorityID, v)
			if e != nil {
				return e
			}
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Model(&user).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id int) (err error) {
	return app.DBClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&system.SysUser{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return app.DBClient.Model(&system.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"enable":     req.Enable,
		}).Error
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: SetSelfInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetSelfInfo(req system.SysUser) error {
	return app.DBClient.Model(&system.SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: SetSelfSetting
//@description: 设置用户配置
//@param: req datatypes.JSON, uid uint
//@return: err error

func (userService *UserService) SetSelfSetting(req *datatypes.JSON, uid uint) error {
	return app.DBClient.Model(&system.SysUser{}).Where("id = ?", uid).Update("origin_setting", req).Error
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = app.DBClient.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserById(id int) (user *system.SysUser, err error) {
	var u system.SysUser
	err = app.DBClient.Where("id = ?", id).First(&u).Error
	return &u, err
}

// FindUserByPhone 根据手机号查询用户信息
// @author: [oldweipro](https://github.com/oldweipro)
// @function: FindUserByPhone
// @description: 根据手机号查询用户信息
// @param: id int
// @return: err error, user *model.SysUser
func (userService *UserService) FindUserByPhone(phone string) (user *system.SysUser, err error) {
	var u system.SysUser
	err = app.DBClient.Where("phone = ?", phone).First(&u).Error
	return &u, err
}

// CheckUserPhoneExist 查看手机是否存在
// @author: [oldweipro](https://github.com/oldweipro)
// @function: FindUserByPhone
// @description: 查看手机是否存在
// @param: id int
// @return: err error, user *model.SysUser
func (userService *UserService) CheckUserPhoneExist(phone string) bool {
	var c int64
	err := app.DBClient.Where("phone = ?", phone).Count(&c).Error
	if err != nil {
		return false
	}
	return true
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
	var u system.SysUser
	if err = app.DBClient.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: ResetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = app.DBClient.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: ResetPasswordByPhone
//@description: 注册时找回密码
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPasswordByPhone(u *system.SysUser) (err error) {
	var user system.SysUser
	if err = app.DBClient.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return err
	}
	user.Password = utils.BcryptHash(u.Password)
	err = app.DBClient.Save(&user).Error
	return err
}

// VerifyUsernameWithSmsCode 校验用户名短信验证码是否正确
// @author: [oldweipro](https://github.com/oldweipro)
// @function: VerifyUsernameWithSmsCode
// @description: 校验用户名短信验证码是否正确
// @param: id int
// @return: err error, user *model.SysUser
func (userService *UserService) VerifyUsernameWithSmsCode(username string, smsCode string) bool {
	code, ok := app.BlackCache.Get(username)
	if !ok {
		return false
	}
	str, err := code.(string)
	if !err {
		return false
	}
	if str == smsCode {
		app.BlackCache.Delete(username)
		return true
	}
	return false
}
