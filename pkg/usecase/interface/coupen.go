package interfaceUseCase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type ICouponUseCase interface {
	CreateCoupon(*requestmodel.Coupon) (*responsemodel.Coupon, error)
	GetCoupons() (*[]responsemodel.Coupon, error)
	UpdateCouponStatus(string, string) (*responsemodel.Coupon, error)
}
