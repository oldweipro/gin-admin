package request

type SubscribeRequest struct {
	PlanId *uint `json:"planId" form:"planId"`
}
