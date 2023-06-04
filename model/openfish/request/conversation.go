package request

import (
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	"time"
)

type ConversationSearch struct {
	openfish.Conversation
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// ChatReq 请求参数列表
type ChatReq struct {
	Prompt           string `form:"prompt" json:"prompt"`
	ConversationId   *uint  `form:"conversationId" json:"conversationId"`
	ConversationType *uint  `form:"conversationType" json:"conversationType"`
	Sign             string `form:"sign" json:"sign"`
}

type ConversationDataGroups struct {
	Data []map[string]interface{} `json:"data"`
	Uuid uint                     `json:"uuid"`
}
