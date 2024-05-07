package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	useCase interfaceUseCase.ICartUseCase
}

func NewCartHandler(carUseCase interfaceUseCase.ICartUseCase) *CartHandler {
	return &CartHandler{useCase: carUseCase}
}

func (u *CartHandler) CreateCart(c *gin.Context) {

	var cart requestmodel.Cart

	fmt.Println("Context:", c.Keys)

	userID := c.Param("UserID") // Extract UserID from URL parameter
    fmt.Println("Extracted UserID:", userID)
	

	cart.UserID = userID

	if err := c.ShouldBind(&cart); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.CreateCart(&cart)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully added", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *CartHandler) DeleteProductFromCart(c *gin.Context) {

	ProductID := c.Param("productID")
	id := strings.TrimSpace(ProductID)

	if len(id) == 0 {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.IDParamsEmpty)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userID := c.Param("UserID")

	err := u.useCase.DeleteProductFromCart(id, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Deleted", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

func (u *CartHandler) IncrementQuantityCart(c *gin.Context) {

	ProductID := c.Param("productID")
	id := strings.TrimSpace(ProductID)

	if len(id) == 0 {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.IDParamsEmpty)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.QuantityIncriment(id, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *CartHandler) DecrementQuantityCart(c *gin.Context) {

	id := c.Param("productID")

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.QuantityDecrease(id, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *CartHandler) ShowCart(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.ShowCart(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
