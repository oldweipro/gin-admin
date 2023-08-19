package transaction

type ServiceGroup struct {
	ChatTicketService
	WalletsService
	HistoryService
	ProductService
	SubscriptionPlanService
	SubscribeService
}
