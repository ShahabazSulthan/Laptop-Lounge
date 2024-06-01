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

type SellerHandler struct {
	usecase interfaceUseCase.ISellerUseCase
}

func NewSellerHandler(sellerUseCase interfaceUseCase.ISellerUseCase) *SellerHandler {
	return &SellerHandler{usecase: sellerUseCase}
}

// @Summary		Seller Signup
// @Description	Using this handler, a seller can sign up.
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Param			Seller	body		requestmodel.SellerSignup	true	"Seller signup details"
// @Success		200		{object}	response.Response	"Successfully signed up"
// @Failure		400		{object}	response.Response	"Bad request"
// @Router			/seller/signup [post]
func (u *SellerHandler) SellerSignup(c *gin.Context) {
	var sellerDetails requestmodel.SellerSignup

	if err := c.BindJSON(&sellerDetails); err != nil {
		c.JSON(http.StatusBadRequest, "can't bind json with struct")
	}

	data, err := helper.Validation(sellerDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.usecase.SellerSignup(&sellerDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", result, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully signup", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Seller Login
// @Description	Using this handler, a seller can log in.
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Param			Seller	body		requestmodel.SellerLogin	true	"Seller login details"
// @Success		200		{object}	response.Response	"Successfully logged in"
// @Failure		400		{object}	response.Response	"Bad request"
// @Failure		500		{object}	response.Response	"Internal server error"
// @Router			/seller/login [post]
func (u *SellerHandler) SellerLogin(c *gin.Context) {
	var loginData requestmodel.SellerLogin
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	data, err := helper.Validation(loginData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.usecase.SellerLogin(&loginData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", result, err.Error())

		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully login", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Sellers
// @Description	Using this handler, admin can get a list of sellers.
// @Tags			Admin Seller Control
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			page	query		int	true	"Page number for pagination (default 1)"	default(1)
// @Param			limit	query		int	true	"Number of items to return per page (default 5)"	default(5)
// @Success		200		{object}	response.Response	"Successfully retrieved list of sellers"
// @Failure		400		{object}	response.Response	"Bad request"
// @Router			/admin/sellers/getsellers [get]
func (u *SellerHandler) GetSellers(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")

	sellers, count, err := u.usecase.GetAllSellers(page, limit)
	if err != nil {
		// message := fmt.Sprintf("total sellers  %d", *count)
		// finalReslt := response.Responses(http.StatusNotFound, message, "", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		message := fmt.Sprintf("total sellers  %d", *count)
		finalReslt := response.Responses(http.StatusOK, message, sellers, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Block Seller
// @Description	Using this handler, admin can block a seller.
// @Tags			Admin Seller Control
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			sellerID	path		string	true	"Seller ID in the URL path"
// @Success		200		{object}	response.Response	"Successfully blocked seller"
// @Failure		400		{object}	response.Response	"Bad request"
// @Router			/admin/sellers/block/{sellerID} [patch]
func (u *SellerHandler) BlockSeller(c *gin.Context) {

	sellerId := c.Param("sellerID")
	id := strings.TrimSpace(sellerId)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	err := u.usecase.BlockSeller(sellerId)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully block", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Block Seller
// @Description	Using this handler, admin can block a seller
// @Tags			Admin Seller Control
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	query		string	true	"Seller ID in the URL path"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/sellers/unblock/:sellerID  [patch]
func (u *SellerHandler) UnblockSeller(c *gin.Context) {
	sellerId := c.Param("sellerID")
	id := strings.TrimSpace(sellerId)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	err := u.usecase.ActiveSeller(sellerId)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully unblock", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Pending Sellers
// @Description	Using this handler, admin can get a list of pending sellers.
// @Tags			Admin Seller Control
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			page	query		int	true	"Page number for pagination (default 1)"	default(1)
// @Param			limit	query		int	true	"Number of items to return per page (default 5)"	default(5)
// @Success		200		{object}	response.Response	"Successfully retrieved list of pending sellers"
// @Failure		400		{object}	response.Response	"Bad request"
// @Router			/admin/sellers/pending [get]
func (u *SellerHandler) GetPendingSellers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "1")

	sellers, err := u.usecase.GetAllPendingSellers(page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", sellers, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Single Seller Details
// @Description	Using this handler, admin can get details of a single seller.
// @Tags			Admin Seller Control
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			sellerID	path		string	true	"Seller ID in the URL path"
// @Success		200		{object}	response.Response	"Successfully retrieved seller details"
// @Failure		400		{object}	response.Response	"Bad request"
// @Router			/admin/sellers/singleview/{sellerID} [get]
func (u *SellerHandler) FetchSingleSeller(c *gin.Context) {
	sellerID := c.Param("sellerID")
	id := strings.TrimSpace(sellerID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	sellerDetails, err := u.usecase.FetchSingleSeller(sellerID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", sellerDetails, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

// @Summary		Verify Seller
// @Description	Using this handler, admin can verify a seller.
// @Tags			Admin Seller Control
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			sellerID	path		string	true	"Seller ID in the URL path"
// @Success		200	{object}	response.Response	"Successfully verified the seller"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router			/admin/sellers/verify/{sellerID} [patch]
func (u *SellerHandler) VerifySeller(c *gin.Context) {
	sellerID := c.Param("sellerID") // Update parameter name here
	id := strings.TrimSpace(sellerID)

	fmt.Println("--", sellerID)
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is empty"}) // Use a string directly here
		return
	}

	err := u.usecase.ActiveSeller(sellerID) // Pass sellerID here
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Verification Success", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------Seller Profile------------------------------------\\

// @Summary		Get Seller Profile
// @Description	Retrieve details of the seller's profile.
// @Tags			Seller Profile
// @Accept			json
// @Produce		json
// @Param			SellerID	path		string							true	"Seller ID in the URL path"
// @Success		200	{object}	response.Response	"Successfully retrieved the seller's profile"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router			/seller/profile/{SellerID} [get]
func (u *SellerHandler) GetSellerProfile(c *gin.Context) {
	sellerID := c.Param("SellerID")

	sellerProfile, err := u.usecase.GetSellerProfile(sellerID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := response.Responses(http.StatusOK, "", sellerProfile, nil)
	c.JSON(http.StatusOK, finalReslt)
}

// @Summary		Update Seller Profile
// @Description	Update the seller's profile.
// @Tags			Seller Profile
// @Accept			json
// @Produce		json
// @Param			SellerID	path		string							true	"Seller ID in the URL path"
// @Param			profile		body		requestmodel.SellerEditProfile	true	"Seller profile details for updating"
// @Success		200			{object}	response.Response				"Successfully updated the seller's profile"
// @Failure		400			{object}	response.Response				"Bad request"
// @Router			/seller/profile/{SellerID} [patch]
func (u *SellerHandler) EditSellerProfile(c *gin.Context) {
	var profile requestmodel.SellerEditProfile

	sellerID := c.Param("SellerID")

	profile.ID = sellerID

	if err := c.ShouldBind(&profile); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userProfile, err := u.usecase.UpdateSellerProfile(&profile)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Successfully Edited", userProfile, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Seller Dashboard
// @Description	Retrieve details for the seller sales.
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			SellerID	path		string	true	"Seller ID in the URL path"
// @Success		200	{object}	response.Response	"Details retrieved successfully"
// @Failure		401	{object}	response.Response	"Unauthorized. Authentication required."
// @Router			/seller/{SellerID} [get]
func (u *SellerHandler) SellerDashbord(c *gin.Context) {

	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	dashBord, err := u.usecase.GetSellerDashbord(sellerID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", dashBord, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
