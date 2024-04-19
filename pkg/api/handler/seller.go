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

func (u *SellerHandler) GetSellers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
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
