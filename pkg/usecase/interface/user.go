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
	ForgetPassword(*requestmodel.ForgetPassword,string) error
}
