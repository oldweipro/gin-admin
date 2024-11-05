package request

import (
	"github.com/oldweipro/gin-admin/model/system"
)

// SendSmsCodeReq 发送短信验证码结构体
type SendSmsCodeReq struct {
	Username  string `json:"username"`  //用户名，可以是账号，也可以是手机号，后期还可能是邮箱
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// RegisterOrResetPassword 用户注册或重置密码结构体
type RegisterOrResetPassword struct {
	Username   string `json:"userName" example:"用户名"`
	Password   string `json:"passWord" example:"密码"`
	Captcha    string `json:"captcha"`    // 验证码
	CaptchaId  string `json:"captchaId"`  // 验证码ID
	SmsCaptcha string `json:"smsCaptcha"` // 短信验证码
}

// Register User register structure
type Register struct {
	Username     string `json:"userName" example:"用户名"`
	Password     string `json:"passWord" example:"密码"`
	NickName     string `json:"nickName" example:"昵称"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
}

// Login User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// ChangePasswordReq Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// SetUserAuth Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// SetUserAuthorities Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                              // 主键ID
	NickName     string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                               // 用户昵称
	Phone        string                `json:"phone"  gorm:"comment:用户手机号"`                                                             // 用户手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                   // 角色ID
	Email        string                `json:"email"  gorm:"comment:用户邮箱"`                                                              // 用户邮箱
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/oldwei_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string                `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                         // 用户侧边主题
	Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                                                              //冻结用户
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
