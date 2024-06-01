package handler

import (
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
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
// @Security BearerTokenAuth
// @Security RefreshtokenAuth
// @Param productID path string true "ID of the product to add"
// @Success 201 {object} response.Response{data=map[string]string} "Product added to wishlist successfully"
// @Failure 400 {object} response.Response "Bad Request. Invalid input."
// @Failure 401 {object} response.Response "Unauthorized. User ID not found in context."
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /wishlist/{productID} [post]
func (w *WishlistHandler) AddToWishlist(c *gin.Context) {
	// Get the UserID from the context
	userID, exist := c.Get("UserID")
	if !exist {
		finalReslt := response.Responses(http.StatusUnauthorized, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusUnauthorized, finalReslt)
		return
	}

	// Get the productID from the URL path
	productID := c.Param("productID")

	// Add the product to the wishlist using the use case
	err := w.useCase.AddProductToWishlist(productID, userID.(string))
	if err != nil {
		finalReslt := response.Responses(http.StatusInternalServerError, "", nil, err.Error())
		c.JSON(http.StatusInternalServerError, finalReslt)
	} else {
		// Return a success response with user_id and product_id
		finalReslt := response.Responses(http.StatusCreated, "", map[string]string{
			"user_id":    userID.(string),
			"product_id": productID,
		}, nil)
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
