package transaction

type RouterGroup struct {
	ChatTicketRouter
	WalletsRouter
	TransactionHistoryRouter
	ProductRouter
}
