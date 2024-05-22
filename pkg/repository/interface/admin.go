package interfaces

type IAdminRepository interface {
	GetPassword(string) (string, error)

	GetSellersDetailDashBoard(string) (uint, error)
	TotalRevenue() (uint, uint, error)
	GetNetCredit() (uint, error)
}
