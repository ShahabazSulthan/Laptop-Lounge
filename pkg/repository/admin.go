package repository

import (
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"errors"
	"log"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.IAdminRepository {
	return &adminRepository{DB: db}
}

func (d *adminRepository) GetPassword(email string) (string, error) {
	var hashedPassword string

	query := "SELECT password FROM admins WHERE email =?"
	err := d.DB.Raw(query, email).Row().Scan(&hashedPassword)
	if err != nil {
		// Log the actual error for debugging purposes
		log.Println("Error fetching admin password:", err)
		return "", errors.New("error fetching admin password")
	}

	return hashedPassword, nil
}

func (d *adminRepository) GetSellersDetailDashBoard(criteria string) (uint, error) {
	var data uint

	query := "SELECT COUNT(*) FROM sellers WHERE status = $1 OR $1 = ''"
	result := d.DB.Raw(query, criteria).Scan(&data)

	if result.Error != nil {
		// Log the actual error for debugging purposes
		log.Println("Error executing query:", result.Error)
		return 0, result.Error
	}

	return data, nil
}

func (d *adminRepository) TotalRevenue() (uint, uint, error) {
	var count, sum uint
	query := "SELECT COALESCE(COUNT(*), 0), COALESCE(SUM(price), 0) FROM order_products WHERE order_status='delivered'"
	result := d.DB.Raw(query).Row().Scan(&count, &sum)
	if result != nil {
		return 0, 0, resCustomError.ErrAdminDashbord
	}

	return count, sum, nil
}

func (d *adminRepository) GetNetCredit() (uint, error) {
	var credit uint
	query := "SELECT COALESCE(SUM(seller_credit),0) FROM sellers"
	result := d.DB.Raw(query).Scan(&credit)
	if result.Error != nil {
		return 0, resCustomError.ErrAdminDashbord
	}
	return credit, nil
}

func (d *adminRepository) GetCouponDetails() ([]responsemodel.Coupon, error) {
	var coupons []responsemodel.Coupon
	query := "SELECT id, name, type, discount, minimum_required, maximum_allowed, start_date, end_date FROM coupons WHERE status = 'active'"
	result := d.DB.Raw(query).Scan(&coupons)

	if result.Error != nil {
		return nil, result.Error
	}

	return coupons, nil
}
