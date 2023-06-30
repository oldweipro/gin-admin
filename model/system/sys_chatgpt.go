package system

import "github.com/oldweipro/gin-admin/global"

type ChatGpt struct {
	DBName string `json:"dbname,omitempty"`
	Chat   string `json:"chat,omitempty"`
	ChatID string `json:"chatID,omitempty"`
}

type SysChatGptOption struct {
	global.Model
	SK string `json:"sk"`
}

func (SysChatGptOption) TableName() string {
	return "sys_chat_gpt_option"
}

type ChatField struct {
	TABLE_NAME     string
	COLUMN_NAME    string
	COLUMN_COMMENT string
}

type ChatFieldNoTable struct {
	COLUMN_NAME    string
	COLUMN_COMMENT string
}
