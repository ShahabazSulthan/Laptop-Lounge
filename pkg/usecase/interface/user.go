package interfaceUseCase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IuserUseCase interface {
	UserSignup(*requestmodel.UserDetails) (*responsemodel.SignupData, error)
	VerifyOtp(*requestmodel.OtpVerification, string) (responsemodel.OtpValidation, error)
	SendOtp(*requestmodel.SendOtp) (string, error)
	UserLogin(*requestmodel.UserLogin) (responsemodel.UserLogin, error)
	ForgetPassword(*requestmodel.ForgetPassword, string) error

	GetAllUsers(string, string) (*[]responsemodel.UserDetails, *int, error)
	BlcokUser(string) error
	UnblockUser(string) error

	AddAddress(*requestmodel.Address) (*requestmodel.Address, error)
	GetAddress(string) (*[]requestmodel.Address, error)
	EditAddress(*requestmodel.EditAddress) (*requestmodel.EditAddress, error)
	DeleteAddress(string, string) error

	GetProfiles(string) (*requestmodel.UserDetails, error)
	UpdateProfile(*requestmodel.UserEditProfile) (*requestmodel.UserDetails, error)
}
