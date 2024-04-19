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

func (u *CategoryHandler) FetchAllCatogry(c *gin.Context) {
	page := c.Param("page")
	limit := c.DefaultQuery("limit", "1")

	category, err := u.categoryUseCase.GetAllCategory(page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", category, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

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

func (u *CategoryHandler) FetchAllBrand(c *gin.Context) {
	page := c.Param("page")
	limit := c.DefaultQuery("limit", "1")

	brand, err := u.categoryUseCase.GetAllBrand(page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", brand, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

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
