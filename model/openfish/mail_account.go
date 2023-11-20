package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"time"
)

// MailAccount 结构体
type MailAccount struct {
	global.Model
	Username                 string     `json:"username" form:"username" gorm:"column:username;comment:mail账号;"`
	NickName                 string     `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:mail昵称;"`
	Remark                   string     `json:"remark" form:"remark" gorm:"column:remark;comment:mail备注;"`
	Password                 string     `json:"password" form:"password" gorm:"column:password;comment:mail密码;"`
	ClaudeSessionKey         string     `json:"claudeSessionKey" form:"claudeSessionKey" gorm:"column:claude_session_key;comment:claude SessionKey;type:longtext;"`
	ClaudeSessionKeyGetTime  *time.Time `json:"claudeSessionKeyGetTime" form:"claudeSessionKeyGetTime" gorm:"column:claude_session_key_get_time;comment:claude SessionKey 获取时间;"`
	OpenaiPassword           string     `json:"openaiPassword" form:"openaiPassword" gorm:"column:openai_password;comment:openai密码;"`
	OpenaiAccessToken        string     `json:"openaiAccessToken" form:"openaiAccessToken" gorm:"column:openai_access_token;comment:openai AccessToken;type:longtext;"`
	OpenaiPuid               string     `json:"openaiPuid" form:"openaiPuid" gorm:"column:openai_puid;comment:openai puid;type:longtext;"`
	OpenaiAccessTokenUseTime *time.Time `json:"openaiAccessTokenUseTime" form:"openaiAccessTokenUseTime" gorm:"column:openai_access_token_use_time;comment:openai AccessToken 使用时间;"`
	OpenaiAccessTokenGetTime *time.Time `json:"openaiAccessTokenGetTime" form:"openaiAccessTokenGetTime" gorm:"column:openai_access_token_get_time;comment:openai AccessToken 获取时间;"`
	OpenaiSkExpire           *time.Time `json:"openaiSkExpire" form:"openaiSkExpire" gorm:"column:openai_sk_expire;comment:openai sk 过期时间;"`
	SkUsedAt                 *time.Time `json:"skUsedAt" form:"skUsedAt" gorm:"column:sk_used_at;comment:openai sk 使用时间;"`
	OpenaiSk                 string     `json:"openaiSk" form:"openaiSk" gorm:"column:openai_sk;comment:openai密钥;"`
	OpenaiAmount             *uint      `json:"openaiAmount" form:"openaiAmount" gorm:"column:openai_amount;comment:openai余额，使用额度;"`
	OpenaiStatus             *uint      `json:"openaiStatus" form:"openaiStatus" gorm:"column:openai_status;comment:openai状态，是否1启用或0禁用2暂时不可用;"`
}

// TableName MailAccount 表名
func (MailAccount) TableName() string {
	return "mail_account"
}
