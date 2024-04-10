package interfaces

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IUserRepo interface {
	UserSignUp(*requestmodel.UserDetails) (*responsemodel.SignupData,error)
	IsUserExist(string) int
	CheckReferalCodeExist(string) (uint, string, error)
	ChangeUserStatusActive(string) error
	FetchUserID(string) (string, error)
	FetchPasswordUsingPhone(string) (string, error)
	UpdatePassword(string, string) error
}