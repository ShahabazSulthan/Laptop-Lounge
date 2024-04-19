package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase interfaceUseCase.IuserUseCase
}

func NewUserHandler(userUseCase interfaceUseCase.IuserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

// ---------------------------User Signup----------------------------------------

func (u *UserHandler) UserSignup(c *gin.Context) {
	fmt.Println("--- UserSignup called")

	// Parse request body into UserDetails struct
	var userSignupData requestmodel.UserDetails
	if err := c.BindJSON(&userSignupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate user signup data
	validationData, err := helper.Validation(userSignupData)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "Validation failed", validationData, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println("--- UserSignup data:", userSignupData)

	// Call user use case for signup
	resSignup, err := u.userUseCase.UserSignup(&userSignupData)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "Signup failed", nil, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Respond with success message and data
	response := response.Responses(http.StatusOK, "Signup successful", resSignup, nil)
	c.JSON(http.StatusOK, response)
}

// this handler function handles the process of sending an OTP by parsing the request,
// validating the data, calling the appropriate use case method,
//  and returning the response to the client based on the outcome of the OTP sending process.

func (u *UserHandler) SendOtp(c *gin.Context) {
	// Parse request body into SendOtp struct
	var sendOtp requestmodel.SendOtp
	if err := c.BindJSON(&sendOtp); err != nil {
		response := response.Responses(http.StatusBadRequest, "Invalid request body", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate SendOtp data
	data, err := helper.Validation(sendOtp)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "Validation failed", data, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println("***", sendOtp)
	// Call user use case to send OTP
	tempToken, err := u.userUseCase.SendOtp(&sendOtp)
	if err != nil {
		response := response.Responses(http.StatusInternalServerError, "Error sending OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Return success response with temporary token
	response := response.Responses(http.StatusOK, "Successfully sent OTP", tempToken, nil)
	c.JSON(http.StatusOK, response)
}

// this handler function handles the process of verifying an OTP by parsing the request,
// validating the data, calling the appropriate use case method with the token,
// and returning the response to the client based on the outcome of the verification process.

func (u *UserHandler) VerifyOTP(c *gin.Context) {
	var otpData requestmodel.OtpVerification
	token := c.Request.Header.Get("Authorization")

	if err := c.BindJSON(&otpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	data, err := helper.Validation(otpData)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "Validation failed", data, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate and process the authorization token if needed

	result, err := u.userUseCase.VerifyOtp(&otpData, token)
	if err != nil {
		response := response.Responses(http.StatusInternalServerError, "Error verifying OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Return success response with verification result
	response := response.Responses(http.StatusOK, "Successfully verified OTP", result, nil)
	c.JSON(http.StatusOK, response)
}

//  this handler function handles the process of user login by parsing the request,
// validating the data, calling the appropriate use case method,
// and returning the response to the client based on the outcome of the login process.

func (u *UserHandler) UserLogin(c *gin.Context) {

	var loginCredential requestmodel.UserLogin

	if err := c.BindJSON(&loginCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate user login data
	data, err := helper.Validation(loginCredential)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "Validation failed", data, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call user use case for login
	result, err := u.userUseCase.UserLogin(&loginCredential)
	if err != nil {
		response := response.Responses(http.StatusInternalServerError, "Error during login", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Return success response with login result
	response := response.Responses(http.StatusOK, "Successfully logged in", result, nil)
	c.JSON(http.StatusOK, response)
}

// this handler function handles the process of handling forgot password requests by parsing the request,
//  validating the data, calling the appropriate use case method with the token,
//  and returning the response to the client based on the outcome of the forgot password process.

func (u *UserHandler) ForgotPassword(c *gin.Context) {
	var forgotPassword requestmodel.ForgetPassword

	// Parse request body
	if err := c.BindJSON(&forgotPassword); err != nil {
		response := response.Responses(http.StatusBadRequest, "Invalid request body", nil, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Extract authorization token
	token := c.Request.Header.Get("Authorization")

	// Validate request data
	data, err := helper.Validation(forgotPassword)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "Validation failed", data, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call user use case for forgot password
	err = u.userUseCase.ForgetPassword(&forgotPassword, token)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "Forgot password failed", nil, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Send success response
	response := response.Responses(http.StatusOK, "Forgot password request processed successfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "1")

	users, count, err := u.userUseCase.GetAllUsers(page, limit)
	if err != nil {
		// message := fmt.Sprintf("total users  %d", *count)
		// finalReslt := response.Responses(http.StatusNotFound, message, "", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		message := fmt.Sprintf("total users  %d", *count)
		finalReslt := response.Responses(http.StatusOK, message, users, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *UserHandler) BlockUser(c *gin.Context) {
	userID := c.Param("userID")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	err := u.userUseCase.BlcokUser(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully block", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *UserHandler) UnblockUser(c *gin.Context) {
	userID := c.Param("userID")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is empty"})
		return
	}

	err := u.userUseCase.UnblockUser(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully unblock", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------user Address------------------------------------\\

func (u *UserHandler) NewAddress(c *gin.Context) {

	var Address requestmodel.Address

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	Address.Userid = userID

	if err := c.ShouldBind(&Address); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(Address)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userAddress, err := u.userUseCase.AddAddress(&Address)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", userAddress, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *UserHandler) GetAddress(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "1")

	userAddress, err := u.userUseCase.GetAddress(userID, page, limit)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", userAddress, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *UserHandler) EditAddress(c *gin.Context) {

	var Address requestmodel.EditAddress

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	Address.Userid = userID

	if err := c.ShouldBind(&Address); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userAddress, err := u.userUseCase.EditAddress(&Address)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Edited", userAddress, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *UserHandler) DeleteAddress(c *gin.Context) {

	addressID := c.Param("id")
	id := strings.TrimSpace(addressID)

	if len(id) == 0 {
		finalReslt := response.Responses(http.StatusBadRequest, "", "", resCustomError.IDParamsEmpty)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	err := u.userUseCase.DeleteAddress(id, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Deleted", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------user Profile------------------------------------\\

func (u *UserHandler) GetProfile(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	UserProfile, err := u.userUseCase.GetProfiles(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", UserProfile, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *UserHandler) EditProfile(c *gin.Context) {

	var profile requestmodel.UserEditProfile

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	profile.Id = userID

	if err := c.ShouldBind(&profile); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userProfile, err := u.userUseCase.UpdateProfile(&profile)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Edited", userProfile, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
