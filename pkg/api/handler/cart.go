package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
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

// @Summary      Create User Cart
// @Description  Create a user cart.
// @Tags         UserCart
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Security     RefreshtokenAuth
// @Param        cart  body  requestmodel.Cart  true  "Cart details for creating"
// @Success      200   {object}  response.Response  "User cart created successfully"
// @Failure      400   {object}  response.Response  "Bad request"
// @Router       /cart/ [post]
func (u *CartHandler) CreateCart(c *gin.Context) {
	var cart requestmodel.Cart

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	cart.UserID = userID

	if err := c.ShouldBind(&cart); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	// Check if quantity exceeds the limit of 5
	if cart.Quantity > 5 {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, "Maximum cart quantity exceeded (max 5 per user).")
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.CreateCart(&cart)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Successfully added", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary     Delete Item from User Cart
// @Description Delete a product from the user's cart.
// @Tags        UserCart
// @Accept      json
// @Produce     json
// @Security    BearerTokenAuth
// @Security    Refreshtoken
// @Param       productID path string true "Product ID to delete from the cart"
// @Success     200 {object} response.Response "Product deleted from the cart successfully"
// @Failure     400 {object} response.Response "Bad request"
// @Router      /cart/{productID} [delete]
func (u *CartHandler) DeleteProductFromCart(c *gin.Context) {

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

// @Summary      Get User Cart
// @Description  Retrieve all items in the user's cart.
// @Tags         UserCart
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Security     RefreshtokenAuth
// @Success      200  {object}  response.Response  "Successfully retrieved user cart items"
// @Failure      400  {object}  response.Response  "Bad request"
// @Router       /cart/ [get]
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
