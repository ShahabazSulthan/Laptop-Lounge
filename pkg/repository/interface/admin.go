package interfaces

import responsemodel "Laptop_Lounge/pkg/models/responseModel"

type IAdminRepository interface {
	GetPassword(string) (string, error)

	GetSellersDetailDashBoard(string) (uint, error)
	TotalRevenue() (uint, uint, error)
	GetNetCredit() (uint, error)
	GetCouponDetails() ([]responsemodel.Coupon, error)
}
