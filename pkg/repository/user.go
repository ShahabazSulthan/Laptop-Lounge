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

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.IUserRepo {
	return &userRepository{DB: DB}
}

// --------------------User Sign Up----------------------------------------\\

func (d *userRepository) UserSignUp(UserDetails *requestmodel.UserDetails) (*responsemodel.SignupData, error) {

	var userData responsemodel.SignupData

	query := "INSERT INTO users (name,email,phone,password,referal_code) values($1,$2,$3,$4,$5) RETURNING *"

	if UserDetails == nil {
		return nil, errors.New("userdetails cannot be nil")
	}

	if UserDetails.Email == "" || UserDetails.Password == "" {
		return nil, errors.New("email and password are required")
	}

	result := d.DB.Raw(query, UserDetails.Name, UserDetails.Email, UserDetails.Phone, UserDetails.Password, UserDetails.ReferalCode).Scan(&userData)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userData, nil
}

// ----------------Check User Availability-----------------------\\
// it checks whether a user with a specific phone number exists in the database.

func (d *userRepository) IsUserExist(phone string) int {
	var userCount int

	query := "SELECT COUNT(*) FROM users WHERE phone=$1 AND status!=$2"
	err := d.DB.Raw(query, phone, "delete").Row().Scan(&userCount)
	if err != nil {
		fmt.Println("Error for user exist, using same phone in signup")
	}
	return userCount
}

// ----------------Check Referal Code Exist-----------------------\\
//  it checks whether a referral code exists in the database and is associated with an active user.

func (d *userRepository) CheckReferalCodeExist(referalCode string) (uint, string, error) {
	var isExist uint
	var userId string

	query := "SELECT COUNT(*), id FROM users WHERE referal_code = ? AND status = 'active' GROUP BY id"
	result := d.DB.Raw(query, referalCode)
	result.Row().Scan(&isExist, &userId)
	if result.Error != nil {
		return 0, "", result.Error
	}
	return isExist, userId, nil
}

// ----------------Update Password-----------------------\\

func (d *userRepository) UpdatePassword(phone string, password string) error {

	query := "UPDATE users SET password=? WHERE phone=? AND status= 'active'"
	result := d.DB.Exec(query, password, phone)

	if result.Error != nil {
		return errors.New("faced some issue while updating password")
	}

	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}

	return nil
}

// ----------------Change Status Active-----------------------\\

func (d *userRepository) ChangeUserStatusActive(phone string) error {

	fmt.Println(phone)
	query := "UPDATE users SET status = 'active' WHERE phone = ?"
	result := d.DB.Exec(query, phone)

	if result.Error != nil {

		return errors.New("failed to change user status to active")
	} else {
		return nil
	}

}

// ----------------Fetch User ID Using Phone Number-----------------------\\

func (d *userRepository) FetchUserID(phone string) (string, error) {
	var userID string

	query := "SELECT id FROM users WHERE phone=? AND status='active'"
	data := d.DB.Raw(query, phone).Row()

	if err := data.Scan(&userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found or inactive")
		}
		return "", fmt.Errorf("failed to fetch user ID: %w", err)
	}

	return userID, nil
}

// ----------------Fetch Password Using Phone Number-----------------------\\

func (d *userRepository) FetchPasswordUsingPhone(phone string) (string, error) {
	var password string

	query := "SELECT password FROM users WHERE phone=? AND status='active'"
	row := d.DB.Raw(query, phone).Row()

	err := row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user does not exist or is blocked")
		}
		return "", fmt.Errorf("error scanning row: %w", err)
	}

	return password, nil
}
