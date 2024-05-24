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

type CategoryHandler struct {
	categoryUseCase interfaceUseCase.ICategoryUseCase
}

func NewCategoryHandler(useCase interfaceUseCase.ICategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase: useCase}
}

// @Summary		Add Category
// @Description	Using this handler, admin can add a new category
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			Category	Details		body	requestmodel.Category	true	"Details of the category"
// @Success		200			{object}	response.Response{}
// @Failure		400			{object}	response.Response{}
// @Router			/admin/category/ [post]
func (u *CategoryHandler) NewCategory(c *gin.Context) {

	var categoryDetails requestmodel.Category

	err := c.BindJSON(&categoryDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, resCustomError.BindingConflict)
	}

	data, err := helper.Validation(categoryDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.categoryUseCase.NewCategory(&categoryDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", result, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Category succesfully added", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary      Get All Categories
// @Description  Using this handler, admin can get a list of all categories
// @Tags         Category
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Param        page  query  int  true  "Page number for pagination (default 1)"  default(1)
// @Param        limit  query  int  true  "Number of items to return per page (default 5)"  default(5)
// @Success      200    {object}  response.Response  "Paginated list of categories"
// @Failure      400    {object}  response.Response  "Bad request"
// @Router       /admin/category/ [get]
func (u *CategoryHandler) FetchAllCatogry(c *gin.Context) {

	category, err := u.categoryUseCase.GetAllCategory()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", category, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

// @Summary      Edit a Category by ID
// @Description  Edit an existing category using this handler.
// @Tags         Category
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Param        category  body  requestmodel.CategoryDetails  true  "Updated category"
// @Success      200       {object}  response.Response  "Category edited successfully"
// @Failure      400       {object}  response.Response  "Invalid input or validation error"
// @Router       /admin/category/ [patch]
func (u *CategoryHandler) UpdateCategory(c *gin.Context) {
	var categoryData requestmodel.CategoryDetails

	if err := c.BindJSON(&categoryData); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(categoryData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	categoryRes, err := u.categoryUseCase.EditCategory(&categoryData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "refine request", categoryRes, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully Updated", categoryRes, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary      Delete a Category by ID
// @Description  Delete an existing category using this handler.
// @Tags         Category
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Param        id  path  int  true  "ID of the category to delete"
// @Success      204  "Category deleted successfully"
// @Failure      400  {object}  response.Response  "Invalid input or validation error"
// @Router       /admin/category/{id} [delete]
func (u *CategoryHandler) DeleteCategory(c *gin.Context) {

	id := c.Param("id")

	err := u.categoryUseCase.DeleteCategory(id)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully category deleted", nil, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//----------------------------------Brand-----------------------------------------------

// @Summary      Create a Brand
// @Description  Create a new brand using this handler.
// @Tags         Brand
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Param        Brand  body  requestmodel.Brand  true  "Name of the brand"
// @Success      201    {object}  response.Response  "Brand created successfully"
// @Failure      400    {object}  response.Response  "Invalid input or validation error"
// @Router       /admin/brand/ [post]
func (u *CategoryHandler) CreateBrand(c *gin.Context) {
	var BrandDetails requestmodel.Brand

	err := c.ShouldBindJSON(&BrandDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, resCustomError.BindingConflict)
		return
	}

	data, err := helper.Validation(BrandDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.categoryUseCase.CreateBrand(&BrandDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", result, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Brand succesfully added", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary      Get List of Brands
// @Description  Get a paginated list of brands using this handler.
// @Tags         Brand
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Param        page  query  int  true  "Page number for pagination (default 1)"  default(1)
// @Param        limit  query  int  true  "Number of items to return per page (default 5)"  default(5)
// @Success      200    {object}  response.Response  "Paginated list of brands"
// @Failure      400    {object}  response.Response  "Invalid input or validation error"
// @Router       /admin/brand/ [get]
func (u *CategoryHandler) FetchAllBrand(c *gin.Context) {

	brand, err := u.categoryUseCase.GetAllBrand()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", brand, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

// @Summary      Edit a Brand by ID
// @Description  Edit an existing brand using this handler.
// @Tags         Brand
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Param        name  body  requestmodel.BrandDetails  true  "Updated name of the brand"
// @Success      200   {object}  response.Response  "Brand edited successfully"
// @Failure      400   {object}  response.Response  "Invalid input or validation error"
// @Router       /admin/brand/ [patch]
func (u *CategoryHandler) UpdateBrand(c *gin.Context) {
	var brandData requestmodel.BrandDetails

	if err := c.BindJSON(&brandData); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(brandData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	brandRes, err := u.categoryUseCase.EditBrand(&brandData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "refine request", brandRes, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully acomplish", brandRes, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Delete a Brand by ID
// @Description Delete an existing brand by its ID.
// @Tags Brand
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param id path int true "ID of the brand to delete"
// @Success 204 "Brand deleted successfully"
// @Failure 400 {object} response.Response{} "Invalid input or validation error"
// @Router /admin/brand/{id} [delete]
func (u *CategoryHandler) DeleteBrand(c *gin.Context) {

	id := c.Param("id")

	err := u.categoryUseCase.DeleteBrand(id)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully Brand deleted", nil, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//--------------------------Category Offer------------------------------------------------

//---------------------------Create Category Offer--------------------------------------------------------//

// CreateCategoryOffer creates a new offer for a category by the seller.
// @Summary      Create Category Offer
// @Description  Create a new offer for a category by the seller.
// @Tags         Seller category offers
// @Accept       json
// @Produce      json
// @Param        categoryOffer body requestmodel.CategoryOffer true "Details for creating a category offer"
// @Success      201 {object} response.Response "Category offer created successfully"
// @Failure      400 {object} response.Response "Bad request. Please provide valid details for creating a category offer."
// @Router       /seller/categoryoffer/{seller_id} [post]
func (u *CategoryHandler) CreateCategoryOffer(c *gin.Context) {
	var categoryOffer requestmodel.CategoryOffer
	if err := c.BindJSON(&categoryOffer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}

	// Set SellerID from the request context if available, otherwise handle as needed
	if sellerID, ok := c.Get("seller_id"); ok {
		categoryOffer.SellerID = sellerID.(string)
	}

	if data, err := helper.Validation(categoryOffer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": data})
		return
	}

	result, err := u.categoryUseCase.CategoryOffer(&categoryOffer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create category offer", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category offer created successfully", "data": result})
}

//---------------------------Block Category Offer--------------------------------------------------------//

// BlockCategoryOffer blocks or disables a category offer by the seller.
// @Summary      Block Category Offer
// @Description  Block or disable a category offer by the seller.
// @Tags         Seller category offers
// @Accept       json
// @Produce      json
// @Param        categoryOfferID path int true "ID of the category offer to be blocked"
// @Success      200 {object} response.Response "Category offer blocked successfully"
// @Failure      400 {object} response.Response "Bad request. Please provide a valid category offer ID."
// @Router       /seller/categoryoffer/block/{categoryOfferID} [patch]
func (u *CategoryHandler) BlockCategoryOffer(c *gin.Context) {

	categoryOfferID := c.Param("categoryOfferID")
	result, err := u.categoryUseCase.ChangeStatusOfCategoryOffer("block", categoryOfferID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully offer block ", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------UnBlock of Category Offer--------------------------------------------------------//

// UnBlockCategoryOffer unblocks or enables a previously blocked category offer by the seller.
// @Summary      Unblock Category Offer
// @Description  Unblock or enable a previously blocked category offer by the seller.
// @Tags         Seller category offers
// @Accept       json
// @Produce      json
// @Param        categoryOfferID path int true "ID of the category offer to be unblocked"
// @Success      200 {object} response.Response "Category offer unblocked successfully"
// @Failure      400 {object} response.Response "Bad request. Please provide a valid category offer ID."
// @Router       /seller/categoryoffer/unblock/{categoryOfferID} [patch]
func (u *CategoryHandler) UnBlockCategoryOffer(c *gin.Context) {

	categoryOfferID := c.Param("categoryOfferID")
	result, err := u.categoryUseCase.ChangeStatusOfCategoryOffer("active", categoryOfferID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully offer unblock ", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------Delete Category Offer--------------------------------------------------------//

// DeleteCategoryOffer deletes a category offer by the seller.
// @Summary      Delete Category Offer
// @Description  Delete a category offer by the seller.
// @Tags         Seller category offers
// @Accept       json
// @Produce      json
// @Param        categoryOfferID path int true "ID of the category offer to be deleted"
// @Success      200 {object} response.Response "Category offer deleted successfully"
// @Failure      400 {object} response.Response "Bad request. Please provide a valid category offer ID."
// @Router       /seller/categoryoffer/delete/{categoryOfferID} [patch]
func (u *CategoryHandler) DeleteCategoryOffer(c *gin.Context) {

	categoryOfferID := c.Param("categoryOfferID")
	result, err := u.categoryUseCase.ChangeStatusOfCategoryOffer("delete", categoryOfferID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully offer deleted ", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------Get ALL Category Offer--------------------------------------------------------//

// GetAllCategoryOffer retrieves all category offers by the seller.
// @Summary      Get Seller Category Offers
// @Description  Retrieve all category offers by the seller.
// @Tags         Seller category offers
// @Accept       json
// @Produce      json
// @Success      200 {object} response.Response "Category offers retrieved successfully"
// @Failure      400 {object} response.Response "Bad request. Unable to retrieve category offers."
// @Router       /seller/categoryoffer/{SellerID} [get]
func (u *CategoryHandler) GetAllCategoryOffer(c *gin.Context) {
	sellerID := c.Param("SellerID")

	result, err := u.categoryUseCase.GetAllCategoryOffer(sellerID)
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
	} else {
		finalResult := response.Responses(http.StatusOK, "Category offers", result, nil)
		c.JSON(http.StatusOK, finalResult)
	}
}

// ---------------------------Edit Category Offer--------------------------------------------------------//
// EditCategoryOffer edits details of a category offer by the seller.
// @Summary      Edit Category Offer
// @Description  Edit details of a category offer by the seller.
// @Tags         Seller category offers
// @Accept       json
// @Produce      json
// @Param        editDetails body requestmodel.EditCategoryOffer true "Details for editing a category offer"
// @Success      200 {object} response.Response "Category offer edited successfully"
// @Failure      400 {object} response.Response "Bad request. Please provide valid edit details."
// @Router       /seller/categoryoffer/{seller_id} [patch]
func (u *CategoryHandler) EditCategoryOffer(c *gin.Context) {
	// Retrieve the seller ID from the URL path parameter
	sellerID := c.Param("seller_id")

	var categoryOffer requestmodel.EditCategoryOffer
	if err := c.BindJSON(&categoryOffer); err != nil {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	categoryOffer.SellerID = sellerID
	data, err := helper.Validation(categoryOffer)
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	result, err := u.categoryUseCase.UpdateCategoryOffer(&categoryOffer)
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
	} else {
		finalResult := response.Responses(http.StatusOK, "Successfully offer updated", result, nil)
		c.JSON(http.StatusOK, finalResult)
	}
}
