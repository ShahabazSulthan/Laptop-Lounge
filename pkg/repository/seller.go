package repository

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
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

func (d *sellerRepository) BlockSeller(id string) error {
	query := "UPDATE sellers SET status = 'block' WHERE id = ?"
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("block seller process is not satisfied")
	}
	count := err.RowsAffected

	if count <= 0 {
		return errors.New("no seller exists with the provided ID")
	}
	return nil
}

func (d *sellerRepository) UnblockSeller(id string) error {
	query := "UPDATE sellers SET status = 'active' WHERE id = ?"
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("active seller process is not satisfied")
	}

	count := err.RowsAffected
	if count <= 0 {
		return errors.New("no seller exists with the provided ID")
	}
	return nil
}

func (d *sellerRepository) GetPendingSellers(offSet int, limit int) (*[]responsemodel.SellerDetails, error) {
	var sellers []responsemodel.SellerDetails

	query := "SELECT * FROM sellers WHERE status='pending' ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&sellers).Error
	if err != nil {
		return nil, errors.New("failed to get the pending list of sellers")
	}

	return &sellers, nil
}

func (d *sellerRepository) GetSingleSeller(id string) (*responsemodel.SellerDetails, error) {
	var seller responsemodel.SellerDetails

	query := "SELECT * FROM sellers WHERE id= ?"
	err := d.DB.Raw(query, id).Scan(&seller)
	if err.Error != nil {
		return nil, errors.New("something wrong at fetching seller details")
	}

	count := err.RowsAffected

	if count <= 0 {
		return nil, errors.New("no seller exists with the provided ID")
	}

	return &seller, nil
}

// ------------------------------------------Seller Profile------------------------------------\\

func (d *sellerRepository) GetSellerProfile(userID string) (*responsemodel.SellerProfile, error) {

	var sellerProfile responsemodel.SellerProfile

	query := "SELECT * FROM sellers WHERE id= ?"
	result := d.DB.Raw(query, userID).Scan(&sellerProfile)
	if result.Error != nil {
		return nil, errors.New("face some issue while get user profile ")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &sellerProfile, nil
}

func (d *sellerRepository) UpdateSellerProfile(editedProfile *requestmodel.SellerEditProfile) (*responsemodel.SellerProfile, error) {

	var profile responsemodel.SellerProfile

	query := "UPDATE sellers SET name=?, email=?, password=? WHERE id= ? RETURNING *;"
	result := d.DB.Raw(query, editedProfile.Name, editedProfile.Email, editedProfile.Password, editedProfile.ID).Scan(&profile)

	if result.Error != nil {
		return nil, errors.New("face some issue while update profile")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &profile, nil
}

// ------------------------------------------Seller Money Menagment------------------------------------\\

func (d *sellerRepository) UpdateSellerCredit(sellerID string, credit uint) error {
	query := "UPDATE sellers SET seller_credit = ? WHERE id=?"
	result := d.DB.Exec(query, credit, sellerID)
	if result.Error != nil {
		return errors.New("face some issue while updating seller credits")
	}
	if result.RowsAffected == 0 {

		return resCustomError.ErrNoRowAffected
	}
	return nil
}

func (d *sellerRepository) GetSellerCredit(sellerID string) (uint, error) {
	var credit uint
	query := "SELECT seller_credit FROM sellers WHERE id= ?"
	result := d.DB.Raw(query, sellerID).Scan(&credit)
	if result.Error != nil {
		return 0, errors.New("face some issue while get credit from seller table")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return credit, nil
}

func (d *sellerRepository) GetDashBordOrderCount(sellerID string, orderstatus string) (uint, error) {

	var data uint
	query := "SELECT COALESCE(COUNT(*),0) FROM order_products WHERE seller_id= $1 AND (order_status=$2 OR $2='')"
	result := d.DB.Raw(query, sellerID, orderstatus).Scan(&data)
	if result.Error != nil {
		return 0, errors.New("face some issue while get dashbord criteria")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return data, nil
}

func (d *sellerRepository) GetDashBordOrderSum(sellerID string, criteria string) (uint, error) {

	var data uint
	query := "SELECT COALESCE(SUM(" + criteria + "),0) FROM order_products WHERE seller_id= ? AND order_status= 'delivered'"
	result := d.DB.Raw(query, sellerID).Scan(&data)
	if result.Error != nil {
		return 0, errors.New("face some issue while get dashbord criteria")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return data, nil
}

// func (d *sellerRepository) GetSellerCredit(sellerID string) (uint, error) {

// 	var data uint
// 	query := "SELECT seller_credit WHERE id=?"
// 	result := d.DB.Raw(query, sellerID).Scan(data)
// 	if result.Error != nil {
// 		return nil, errors.New("face some issue while get dashbord criteria")
// 	}
// 	if result.RowsAffected == 0 {
// 		return nil, resCustomError.ErrNoRowAffected
// 	}
// 	return data, nil
// }

func (d *sellerRepository) GetLowStokesProduct(sellerID string) ([]uint, error) {

	var data []uint
	query := "SELECT COALESCE(id,0) FROM products WHERE seller_id= ? AND units<100"
	result := d.DB.Raw(query, sellerID).Scan(&data)
	if result.Error != nil {
		return nil, errors.New("face some issue while get dashbord criteria")
	}
	return data, nil
}