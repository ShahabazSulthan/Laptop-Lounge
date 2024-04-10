package repository

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type sellerRepository struct {
	DB *gorm.DB
}

func NewSellerRepository(db *gorm.DB) interfaces.ISellerRepo {
	return &sellerRepository{DB: db}
}

func (d *sellerRepository) IsSellerExist(email string) (int, error) {
	const (
		StatusDelete = "delete"
	)

	var sellerCount int

	query := "SELECT COUNT(*) FROM sellers WHERE email=$1 AND status!=$2"
	err := d.DB.Raw(query, email, StatusDelete).Row().Scan(&sellerCount)
	if err == sql.ErrNoRows {
		return 0, nil // No seller found
	} else if err != nil {
		return 0, fmt.Errorf("error fetching seller count using email: %w", err)
	}

	return sellerCount, nil
}

func (d *sellerRepository) CreateSeller(SellerData *requestmodel.SellerSignup) error {
	query := `
    INSERT INTO sellers (name, email, password, gst_no, descriptioin)
    VALUES($1, $2, $3, $4, $5)
`

	err := d.DB.Exec(query, SellerData.Name, SellerData.Email, SellerData.Password, SellerData.GST_NO, SellerData.Description).Error
	if err != nil {
		return fmt.Errorf("error creating seller: %w", err)
	}

	return nil
}

func (d *sellerRepository) GetHashPassAndStatus(email string) (string, string, string, error) {
	var password, status, sellerID string
	query := "SELECT password, id, status FROM sellers WHERE email=? AND status!='delete'"
	err := d.DB.Raw(query, email).Row().Scan(&password, &sellerID, &status)
	if err != nil {
		return "", "", "", errors.New("feching password and status, can't make action in database")
	}
	return password, sellerID, status, nil
}

func (d *sellerRepository) GetPasswordByMail(email string) string {
	var hashedPassword string
	query := "SELECT password FROM sellers WHERE email=? AND status='active'"
	err := d.DB.Raw(query, email).Row().Scan(&hashedPassword)
	if err != nil {
		return ""
	}
	return hashedPassword
}

func (d *sellerRepository) AllSellers(offSet int, limit int) (*[]responsemodel.SellerDetails, error) {
	var sellers []responsemodel.SellerDetails

	query := "SELECT * FROM sellers ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&sellers).Error
	if err != nil {
		return nil, errors.New("can't get seller data from db")
	}

	return &sellers, nil
}

func (d *sellerRepository) SellerCount(ch chan int) {
	var count int

	query := "SELECT COUNT(email) FROM sellers WHERE status = 'block' WHERE id=? "
	d.DB.Raw(query).Scan(&count)
	ch <- count

}
