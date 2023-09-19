package transaction

type RouterGroup struct {
	ChatTicketRouter
	WalletsRouter
	TransactionHistoryRouter
	ProductRouter
	SubscriptionPlanRouter
	RedeemRouter
	ActivationCodeRouter
}
