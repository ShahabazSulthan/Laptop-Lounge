package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	useCase interfaceUseCase.IwishlistRepo
}

func NewwishlistHandler(useCase interfaceUseCase.IwishlistRepo) *WishlistHandler {
	return &WishlistHandler{useCase: useCase}
}

// @Summary Add Product to Wishlist
// @Description Add a product to the user's wishlist.
// @Tags Wishlist
// @Accept json
// @Produce json
// @Security     BearerTokenAuth
// @Security     RefreshtokenAuth
// @Param productID path string true "ID of the product to add"
// @Success 201 {object} response.Response "Product added to wishlist successfully"
// @Failure 400 {object} response.Response "Bad Request. Invalid input."
// @Failure 401 {object} response.Response "Unauthorized. User ID not found in context."
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /wishlist/{productID} [post]
func (w *WishlistHandler) AddToWishlist(c *gin.Context) {
	var wishlist *requestmodel.WishlistRequest

	if err := c.BindJSON(&wishlist); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(wishlist)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	productID := c.Param("productID")

	err = w.useCase.AddProductToWishlist(productID, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusCreated, "", wishlist, nil) // Use HTTP 201 for resource creation
		c.JSON(http.StatusCreated, finalReslt)
	}
}

// @Summary Get User Wishlist
// @Description Retrieve the user's wishlist.
// @Tags Wishlist
// @Accept json
// @Produce json
// @Security     BearerTokenAuth
// @Security     RefreshtokenAuth
// @Success 200 {object} response.Response "Wishlist retrieved successfully"
// @Failure 400 {object} response.Response "Bad Request. User ID not found in context."
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /wishlist [get]
func (w *WishlistHandler) GetWishlist(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := w.useCase.ViewUserWishlist(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Remove Product from Wishlist
// @Description Remove a product from the user's wishlist.
// @Tags Wishlist
// @Accept json
// @Produce json
// @Security     BearerTokenAuth
// @Security     RefreshtokenAuth
// @Param productID path string true "ID of the product to remove"
// @Success 200 {object} response.Response "Product removed from wishlist successfully"
// @Failure 400 {object} response.Response "Bad Request. User ID not found in context."
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /wishlist/{productID} [delete]
func (w *WishlistHandler) DeleteWishlist(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	productID := c.Param("productID")

	err := w.useCase.RemoveProductFromWishlist(productID, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Deleted", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}
