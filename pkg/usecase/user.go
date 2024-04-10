package usecase

import (
	"Laptop_Lounge/pkg/config"
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"Laptop_Lounge/pkg/service"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"errors"
	"fmt"
	"log"
)

type userUseCase struct {
	repo  interfaces.IUserRepo
	token config.Token
}

func NewUserUseCase(userRepository interfaces.IUserRepo, token *config.Token) interfaceUseCase.IuserUseCase {
	return &userUseCase{
		repo:  userRepository,
		token: *token,
	}
}

// -------------------------------USER SIGNUP --------------------------------------------------------

func (u *userUseCase) UserSignup(userData *requestmodel.UserDetails) (*responsemodel.SignupData, error) {

	// Initialize response data
	var signupData *responsemodel.SignupData

	// Check if the user already exists
	if isExist := u.repo.IsUserExist(userData.Phone); isExist >= 1 {
		return nil, errors.New("user with this phone number already exists, please try again with another phone number")
	}

	// Set up Twilio and send OTP
	if err := service.TwilioSetup(); err != nil {
		return nil, fmt.Errorf("error setting up Twilio: %v", err)
	}
	_, err := service.SendOtp(userData.Phone)
	if err != nil {
		return nil, fmt.Errorf("error sending OTP: %v", err)
	}

	// isExist, referalOFUserID, err := u.repo.CheckReferalCodeExist(userData.ReferalCode)
	// if err != nil {
	// 	return nil, err
	// }

	// Hash password and generate referral code

	userData.Password = helper.HashPassword(userData.Password)
	userData.ReferalCode, _ = helper.GenerateReferalCode()

	// Save user data
	signupData, err = u.repo.UserSignUp(userData)
	if err != nil {
		return nil, fmt.Errorf("error signing up user: %v", err)
	}

	// Generate token for OTP verification

	token, err := service.TemperveryTokenForOtpVerification(u.token.TemperveryKey, userData.Phone)
	if err != nil {
		fmt.Println("errr2", err)
		return nil, fmt.Errorf("error generating verification token: %v", err)
	}
	signupData.Token = token

	return signupData, nil
}

// SendOtp sends an OTP to the specified phone number and returns the OTP along with any errors encountered.
func (r *userUseCase) SendOtp(phone *requestmodel.SendOtp) (string, error) {
	// Setup Twilio services
	if err := service.TwilioSetup(); err != nil {
		return "", fmt.Errorf("error setting up Twilio: %v", err)
	}

	// Log phone number for debugging
	fmt.Println("Sending OTP to phone:", phone.Phone)

	// Send OTP
	otpCode, err := service.SendOtp(phone.Phone)
	if err != nil {
		return "", fmt.Errorf("error sending OTP: %v", err)
	}

	// Generate token for OTP verification
	token, err := service.TemperveryTokenForOtpVerification(r.token.TemperveryKey, phone.Phone)
	if err != nil {
		return "", fmt.Errorf("error generating verification token: %v", err)
	}

	// Log success and return token
	log.Printf("OTP sent successfully to %s with code: %s", phone.Phone, otpCode)

	return token, nil
}

//   verifies the OTP, changes the user status, fetches the user ID, generates access and refresh tokens, and prepares the response accordingly

func (u *userUseCase) VerifyOtp(otpConstrain *requestmodel.OtpVerification, token string) (responsemodel.OtpValidation, error) {
	var otpResponse responsemodel.OtpValidation

	// Extract phone number from token
	phone, err := service.FetchPhoneFromToken(token, u.token.TemperveryKey)
	if err != nil {
		otpResponse.Token = "Invalid token"
		return otpResponse, fmt.Errorf("error extracting token: %v", err)
	}

	// Verify OTP
	service.TwilioSetup()
	if err := service.VerifyOtp(phone, otpConstrain.Otp); err != nil {
		otpResponse.Otp = "OTP verification failed"
		return otpResponse, fmt.Errorf("OTP verification error: %v", err)
	}

	// Change user status
	if err := u.repo.ChangeUserStatusActive(phone); err != nil {
		otpResponse.Phone = "User not found, verify the phone number"
		return otpResponse, fmt.Errorf("error changing user status: %v", err)
	}

	// Fetch user ID
	userID, err := u.repo.FetchUserID(phone)
	if err != nil {
		otpResponse.Result = "Error fetching user ID"
		return otpResponse, fmt.Errorf("error fetching user ID: %v", err)
	}

	// Generate access and refresh tokens
	accessToken, err := service.GenerateAcessToken(u.token.UsersSecurityKey, userID)
	if err != nil {
		otpResponse.Result = "Error generating access token"
		return otpResponse, fmt.Errorf("error generating access token: %v", err)
	}

	refreshToken, err := service.GenerateRefreshToken(u.token.UsersSecurityKey)
	if err != nil {
		otpResponse.Result = "Error generating refresh token"
		return otpResponse, fmt.Errorf("error generating refresh token: %v", err)
	}

	// Prepare response
	otpResponse.AccessToken = accessToken
	otpResponse.RefreshToken = refreshToken
	otpResponse.Result = "success"

	return otpResponse, nil
}

// this method handles the user login process by fetching the password, comparing passwords, fetching the user ID, generating tokens, and preparing the response object.

func (u *userUseCase) UserLogin(loginCredential *requestmodel.UserLogin) (responsemodel.UserLogin, error) {
	var resUserLogin responsemodel.UserLogin
	fmt.Println("---", loginCredential)

	// Validate input
	if loginCredential == nil || loginCredential.Phone == "" || loginCredential.Password == "" {
		return resUserLogin, errors.New("invalid login credentials")
	}

	// Fetch password from repository
	password, err := u.repo.FetchPasswordUsingPhone(loginCredential.Phone)
	if err != nil {
		return resUserLogin, fmt.Errorf("error fetching password: %v", err)
	}

	// Compare passwords
	if err := helper.CompairPassword(password, loginCredential.Password); err != nil {
		return resUserLogin, fmt.Errorf("incorrect password: %v", err)
	}

	// Fetch user ID
	userID, err := u.repo.FetchUserID(loginCredential.Phone)
	if err != nil {
		return resUserLogin, fmt.Errorf("error fetching user ID: %v", err)
	}

	// Generate access token
	accessToken, err := service.GenerateAcessToken(u.token.UsersSecurityKey, userID)
	if err != nil {
		return resUserLogin, fmt.Errorf("error generating access token: %v", err)
	}

	// Generate refresh token
	refreshToken, err := service.GenerateRefreshToken(u.token.UsersSecurityKey)
	if err != nil {
		return resUserLogin, fmt.Errorf("error generating refresh token: %v", err)
	}

	// Populate response object
	resUserLogin.AccessToken = accessToken
	resUserLogin.RefreshToken = refreshToken

	return resUserLogin, nil
}

//  the method follows the workflow of verifying the OTP, hashing the new password, and updating the password in the repository.

func (r *userUseCase) ForgetPassword(newPassword *requestmodel.ForgetPassword, token string) error {
	// Fetch phone number from token
	phone, err := service.FetchPhoneFromToken(token, r.token.TemperveryKey)
	if err != nil {
		return err
	}

	// Verify OTP
	if err := service.VerifyOtp(phone, newPassword.Otp); err != nil {
		return err
	}

	// Hash new password
	hashPassword := helper.HashPassword(newPassword.Password)

	// Update password in repository
	if err := r.repo.UpdatePassword(phone, hashPassword); err != nil {
		return err
	}

	// Return nil if all steps are successful
	return nil
}
