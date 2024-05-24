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

// @Summary		User Signup
// @Description	Using this handler, users can sign up.
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body		requestmodel.UserDetails	true	"User Signup details"
// @Success		200		{object}	responsemodel.SignupData
// @Failure		400		{object}	response.Response
// @Router			/signup [post]
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
	fmt.Println("333", resSignup)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "Signup failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Respond with success message and data
	response := response.Responses(http.StatusOK, "Signup successful", resSignup, nil)
	c.JSON(http.StatusOK, response)
}

// ---------------------------Send OTP------------------------------------------

// @Summary		Send OTP To Mobile
// @Description	Send OTP (One-Time Password) for verification.
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			otp	body		requestmodel.SendOtp	true	"OTP details for sending"
// @Success		200	{object}	response.Response	"OTP sent successfully"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router			/sendOTP [post]
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

// ---------------------------Verify OTP----------------------------------------

// @Summary	Verify OTP
// @Description	Verify the OTP (One-Time Password) sent to the user.
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		otp	body		requestmodel.OtpVerification	true	"OTP details for verification"
// @Success	200	{object}	response.Response		"OTP verified successfully"
// @Failure	400	{object}	response.Response		"Bad request"
// @Router		/verifyOTP [post]
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

// ---------------------------User Login----------------------------------------

// @Summary		User Login
// @Description	Using this handler, users can log in.
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body		requestmodel.UserLogin	true	"User Login details"
// @Success		200		{object}	response.Response
// @Failure		400		{object}	response.Response
// @Router			/login [post]
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

// ---------------------------Forgot Password----------------------------------------


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

// ---------------------------Get All Users----------------------------------------

// @Summary		Get All Users
// @Description	Using this handler, admin can view users.
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			page	query		int	false	"Page number"				default(1)
// @Param			limit	query		int	false	"Number of items per page"	default(5)
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/admin/users/getuser [get]
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

// ---------------------------Block User----------------------------------------

// @Summary		Block User
// @Description	Using this handler, admin can block a user.
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			userID	path	string	true	"User ID in the URL path"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/admin/users/block/{userID} [patch]
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

// ---------------------------Unblock User----------------------------------------

// @Summary		Unblock User
// @Description	Using this handler, admin can unblock a user.
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			userID	path	string	true	"User ID in the URL path"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/admin/users/unblock/{userID} [patch]
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

// ---------------------------Add Address----------------------------------------

// @Summary		Add Address
// @Description	Add a new address.
// @Tags			User Addresses
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			address	body		requestmodel.Address	true	"Address object to be added"
// @Success		201		{object}	response.Response		"Successfully added the address"
// @Failure		400		{object}	response.Response		"Bad request"
// @Router			/address/ [post]
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

// ---------------------------Get Addresses----------------------------------------

// @Summary		Get Addresses
// @Description	Retrieve a list of addresses.
// @Tags			User Addresses
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			page	query		int						false	"Page number"				default(1)
// @Param			limit	query		int						false	"Number of items per page"	default(5)
// @Success		200		{object}	[]requestmodel.Address	"Successfully retrieved addresses"
// @Failure		400		{object}	response.Response		"Bad request"
// @Router			/address/ [get]
func (u *UserHandler) GetAddress(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userAddress, err := u.userUseCase.GetAddress(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", userAddress, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ---------------------------Update Address----------------------------------------

// @Summary		Update Address
// @Description	Update an existing address.
// @Tags			User Addresses
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			address	body		requestmodel.EditAddress	true	"Updated address information"
// @Success		200		{object}	response.Response			"Successfully updated the address"
// @Failure		400		{object}	response.Response			"Bad request"
// @Router			/address/ [patch]
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

// @Summary		Delete Address
// @Description	Delete an address by ID.
// @Tags			User Addresses
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			id	path		string				true	"Address ID in the query parameter"
// @Success		200	{object}	response.Response	"Successfully deleted the address"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router			/address/{id} [delete]
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

// @Summary		Get User
// @Description	Retrieve the user's profile.
// @Tags			User Profile
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Success		200	{object}	requestmodel.UserDetails	"Successfully retrieved the user's profile"
// @Failure		400	{object}	response.Response		"Bad request"
// @Router			/profile [get]
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

// @Summary		Update User Profile
// @Description	Update the user's profile.
// @Tags			User Profile
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			profile	body		requestmodel.UserEditProfile	true	"User profile details for updating"
// @Success		200		{object}	response.Response				"Successfully updated the user's profile"
// @Failure		400		{object}	response.Response				"Bad request"
// @Router			/profile/ [patch]
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
