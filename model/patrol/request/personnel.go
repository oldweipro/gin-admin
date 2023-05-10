package request

import (
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/patrol"
	"time"
)

type PersonnelSearch struct {
	patrol.Personnel
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type PersonnelQueryResult struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    PersonnelQueryResultData `json:"data"`
}
type PersonnelQueryResultData struct {
	EndRow            int                `json:"endRow"`
	HasNextPage       bool               `json:"hasNextPage"`
	HasPreviousPage   bool               `json:"hasPreviousPage"`
	IsFirstPage       bool               `json:"isFirstPage"`
	IsLastPage        bool               `json:"isLastPage"`
	List              []patrol.Personnel `json:"list"`
	NavigateFirstPage int                `json:"navigateFirstPage"`
	NavigateLastPage  int                `json:"navigateLastPage"`
	NavigatePages     int                `json:"navigatePages"`
	NavigatepageNums  []int              `json:"navigatepageNums"`
	NextPage          int                `json:"nextPage"`
	PageNum           int                `json:"pageNum"`
	PageSize          int                `json:"pageSize"`
	Pages             int                `json:"pages"`
	PrePage           int                `json:"prePage"`
	Size              int                `json:"size"`
	StartRow          int                `json:"startRow"`
	Total             int                `json:"total"`
}
