package interfaces

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IUserRepo interface {
	UserSignUp(*requestmodel.UserDetails) (*responsemodel.SignupData, error)
	IsUserExist(string) int
	CheckReferalCodeExist(string) (uint, string, error)
	ChangeUserStatusActive(string) error
	FetchUserID(string) (string, error)
	FetchPasswordUsingPhone(string) (string, error)
	UpdatePassword(string, string) error

	AllUsers(int, int) (*[]responsemodel.UserDetails, error)
	UserCount(chan int)
	BlockUser(string) error
	UnblockUser(string) error

	CreateAddress(*requestmodel.Address) (*requestmodel.Address, error)
	GetAddress(string) (*[]requestmodel.Address, error)
	UpdateAddress(*requestmodel.EditAddress) (*requestmodel.EditAddress, error)
	GetAAddress(string) (*requestmodel.Address, error)
	DeleteAddress(string, string) error

	GetProfile(string) (*requestmodel.UserDetails, error)
	UpdateProfile(*requestmodel.UserDetails) (*requestmodel.UserDetails, error)
}
