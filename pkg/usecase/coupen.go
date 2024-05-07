package usecase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
)

type couponUseCase struct {
	repo interfaces.ICouponRepository
}

func NewCouponUseCase(repository interfaces.ICouponRepository) interfaceUseCase.ICouponUseCase {
	return &couponUseCase{repo: repository}
}

func (r *couponUseCase) CreateCoupon(newCoupon *requestmodel.Coupon) (*responsemodel.Coupon, error) {
	coupon, err := r.repo.CreateCoupon(newCoupon)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (r *couponUseCase) GetCoupons() (*[]responsemodel.Coupon, error) {
	coupons, err := r.repo.GetCoupons()
	if err != nil {
		return nil, err
	}
	return coupons, nil
}

func (r *couponUseCase) UpdateCouponStatus(couponID, status string) (*responsemodel.Coupon, error) {
	var coupon *responsemodel.Coupon
	var err error
	if status == "active" {
		coupon, err = r.repo.UpdateCouponStatus(couponID, status, "")
		if err != nil {
			return nil, err
		}
	} else {
		coupon, err = r.repo.UpdateCouponStatus(couponID, "", status)
		if err != nil {
			return nil, err
		}
	}
	return coupon, nil
}