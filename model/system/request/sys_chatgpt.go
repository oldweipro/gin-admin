package request

import (
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
