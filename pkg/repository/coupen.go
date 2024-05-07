package repository

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"errors"
	"time"

	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

func NewCoupenRepository(db *gorm.DB) interfaces.ICouponRepository {
	return &couponRepository{DB: db}
}

func (d *couponRepository) CreateCoupon(newCoupon *requestmodel.Coupon) (*responsemodel.Coupon, error) {
	var createdCoupon *responsemodel.Coupon

	// Calculate expiration time 3 days from now
	couponExpireTime := time.Now().Add(3 * 24 * time.Hour)

	query := "INSERT INTO coupons (name, type, discount, minimum_required, maximum_allowed, start_date, end_date) VALUES (?, ?, ?, ?, ?, now(), ?) RETURNING *"
	result := d.DB.Raw(query, newCoupon.Name, newCoupon.Type, newCoupon.Discount, newCoupon.MinimumRequired, newCoupon.MaximumAllowed, couponExpireTime).Scan(&createdCoupon)

	if result.Error != nil {
		return nil, errors.New("encountered an issue while creating a new coupon")
	}

	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}

	return createdCoupon, nil
}


func (d *couponRepository) CheckCouponExpired(couponCode string) (*responsemodel.Coupon, error) {

	var couponData responsemodel.Coupon
	query := "SELECT * FROM coupons WHERE id= ? AND status= 'active'"
	result := d.DB.Raw(query, couponCode).Scan(&couponData)
	if result.Error != nil {
		return nil, errors.New("face some issue while check coupon exist")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("not a valid coupon, better luck next time")
	}
	return &couponData, nil

}

func (d *couponRepository) GetCoupons() (*[]responsemodel.Coupon, error) {

	var coupons []responsemodel.Coupon
	query := "SELECT * FROM coupons"
	result := d.DB.Raw(query).Scan(&coupons)
	if result.Error != nil {
		return nil, errors.New("face some issue while get coupons")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &coupons, nil
}

func (d *couponRepository) UpdateCouponStatus(couponID, active, block string) (*responsemodel.Coupon, error) {

	var coupon responsemodel.Coupon
	var result *gorm.DB

	query := "UPDATE coupons SET status= ? WHERE id=? RETURNING*"
	if active != "" {
		result = d.DB.Raw(query, active, couponID).Scan(&coupon)
	}
	if block != "" {
		result = d.DB.Raw(query, block, couponID).Scan(&coupon)
	}
	if result.Error != nil {
		return nil, errors.New("face some issue while update coupons status")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &coupon, nil
}