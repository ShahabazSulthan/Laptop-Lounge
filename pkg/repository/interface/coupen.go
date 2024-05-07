package interfaces

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type ICouponRepository interface {
	CreateCoupon(*requestmodel.Coupon) (*responsemodel.Coupon, error)
	CheckCouponExpired(string) (*responsemodel.Coupon, error)
	GetCoupons() (*[]responsemodel.Coupon, error)
	UpdateCouponStatus(string, string, string) (*responsemodel.Coupon, error)
}
