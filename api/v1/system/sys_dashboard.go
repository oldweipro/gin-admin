package system

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/model/common/response"
)

type DashboardApi struct{}

func (d *DashboardApi) Console(c *gin.Context) {
	consoleStr := `
{
  "visits": {
    "dayVisits": 29902,
    "rise": 82,
    "decline": 49,
    "amount": 658774
  },
  "saleroom": {
    "weekSaleroom": 30402,
    "amount": 944394,
    "degree": 84.1428
  },
  "orderLarge": {
    "weekLarge": 87599,
    "rise": 68,
    "decline": 23,
    "amount": 574423
  },
  "volume": {
    "weekLarge": 69973,
    "rise": 23,
    "decline": 43,
    "amount": 519204
  }
}`
	var console map[string]interface{}
	err := json.Unmarshal([]byte(consoleStr), &console)
	fmt.Println(err)
	response.OkWithDetailed(console, "获取成功", c)
}
