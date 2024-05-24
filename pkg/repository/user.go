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

//---retrieves a list of users from the database based on the provided offset and limit.

func (d *userRepository) AllUsers(offSet int, limit int) (*[]responsemodel.UserDetails, error) {
	var users []responsemodel.UserDetails

	query := "SELECT * FROM users ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&users).Error
	if err != nil {
		return nil, errors.New("can't get user data from db")
	}

	return &users, nil
}

//--counts the number of users in the database whose status is not "delete".

func (d *userRepository) UserCount(ch chan int) {
	var count int
	query := "SELECT COUNT(phone) FROM users WHERE status!='delete'"
	d.DB.Raw(query).Scan(&count)
	ch <- count
}

//--blocks a user in the database based on the provided user ID

func (d *userRepository) BlockUser(id string) error {
	query := "UPDATE users SET status = 'block' WHERE id=? "
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("block user process , is not satisfied")
	}
	count := err.RowsAffected
	if count <= 0 {
		return errors.New("no user exist by id ")
	}
	return nil
}

//--unblocks a previously blocked user in the database based on the provided user ID.

func (d *userRepository) UnblockUser(id string) error {
	query := "UPDATE users SET status = 'active' WHERE id=?"
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("active user process , is not satisfied")
	}

	if err.RowsAffected <= 0 {
		return errors.New("no user exist by id ")
	}
	return nil
}

//-------------------------Address --------------------------------------------
//--creates a new address record in the database.

func (d *userRepository) CreateAddress(address *requestmodel.Address) (*requestmodel.Address, error) {

	query := `INSERT INTO addresses ( userid, first_name, last_name, street, city, state, pincode, land_mark, phone_number) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;`

	result := d.DB.Raw(query,
		address.Userid, address.FirstName, address.LastName,
		address.Street, address.City, address.State, address.Pincode,
		address.LandMark, address.PhoneNumber,
	).Scan(&address)

	if result.Error != nil {
		return nil, errors.New("face some issue while address insertion ")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}

	return address, nil
}

//---Retrieves a list of addresses belonging to a specific user ID from the database.

func (d *userRepository) GetAddress(userID string) (*[]requestmodel.Address, error) {
	if userID == "" {
		return nil, errors.New("userID is empty")
	}

	var address []requestmodel.Address

	query := "SELECT * FROM addresses WHERE userid=? AND status='active' ORDER BY id"
	result := d.DB.Raw(query, userID).Scan(&address)
	if result.Error != nil {
		return nil, errors.New("error fetching address")
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &address, nil
}

//--Updates an existing address record in the database based on the provided address struct.

func (d *userRepository) UpdateAddress(address *requestmodel.EditAddress) (*requestmodel.EditAddress, error) {

	query := "UPDATE addresses SET first_name=?, last_name=?, street=?, city=?, state=?, pincode=?, land_mark=?, phone_number=? WHERE id=? AND userid= ? RETURNING *;"
	result := d.DB.Raw(query,
		address.FirstName, address.LastName,
		address.Street, address.City, address.State, address.Pincode,
		address.LandMark, address.PhoneNumber,
		address.ID, address.Userid,
	).Scan(&address)

	if result.Error != nil {
		return nil, errors.New("face some issue while update address ")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return address, nil
}

//--Retrieves a single address record from the database based on the provided address ID.

func (d *userRepository) GetAAddress(addressID string) (*requestmodel.Address, error) {

	var address requestmodel.Address

	query := "SELECT * FROM addresses WHERE id=?"
	result := d.DB.Raw(query, addressID).Scan(&address)
	if result.Error != nil {
		return nil, errors.New("face some issue while address fetch")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &address, nil
}

//--Deletes an address record from the database based on the provided address ID and user ID.

func (d *userRepository) DeleteAddress(addressID string, userID string) error {

	query := "DELETE FROM addresses WHERE id= ? AND userid= ?"
	result := d.DB.Exec(query, addressID, userID)
	if result.Error != nil {
		return errors.New("face some issue while deleting address ")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}

//-----------------------------User Profile----------------------------------------//

//--retrieves a user profile based on the provided user ID from the database.

func (d *userRepository) GetProfile(userID string) (*requestmodel.UserDetails, error) {

	var userDetails requestmodel.UserDetails

	query := "SELECT id, name , email, phone, referal_code FROM users WHERE id= ?"
	result := d.DB.Raw(query, userID).Scan(&userDetails)
	if result.Error != nil {
		return nil, errors.New("face some issue while get user profile ")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &userDetails, nil
}

//--updates a user's profile information in the database based on the provided edited profile details.

func (d *userRepository) UpdateProfile(editedProfile *requestmodel.UserDetails) (*requestmodel.UserDetails, error) {

	var profile requestmodel.UserDetails

	query := "UPDATE users SET name=?, email=? WHERE id= ? RETURNING *;"
	result := d.DB.Raw(query, editedProfile.Name, editedProfile.Email, editedProfile.Id).Scan(&profile)
	if result.Error != nil {
		return nil, errors.New("face some issue while update profile")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &profile, nil
}
