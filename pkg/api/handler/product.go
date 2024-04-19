package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	userCase interfaceUseCase.IProductUseCase
}

func NewProductHandler(usercase interfaceUseCase.IProductUseCase) *ProductHandler {
	return &ProductHandler{userCase: usercase}
}

// POST endpoint for adding a product for a specific seller
func (ph *ProductHandler) AddProduct(c *gin.Context) {
	var productDetails requestmodel.ProductReq

	// Extract SellerID from the URL parameter
	sellerID := c.Param("SellerID")
	fmt.Println("sellerid", sellerID)

	// Check if SellerID exists in the request
	if sellerID == "" {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	// Bind request body to productDetails
	if err := c.ShouldBind(&productDetails); err != nil {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	// Retrieve productImage from form data
	image, err := c.FormFile("image")
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	productDetails.Image = image

	// Convert SellerID to integer
	sellerIDInt, err := strconv.Atoi(sellerID)
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}
	productDetails.SellerID = uint(sellerIDInt)

	// Validate productDetails
	data, err := helper.Validation(productDetails)
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	// Add product using the use case
	product, err := ph.userCase.AddProduct(&productDetails)

	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, "refine request", "", err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	// Respond with success message and product data
	finalResult := response.Responses(http.StatusOK, "Successfully Uploaded", product, nil)
	c.JSON(http.StatusOK, finalResult)
}

func (u *ProductHandler) BlockProduct(c *gin.Context) {
	sellerid, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	productID := c.Param("productid")
	err := u.userCase.BlockProduct(sellerid, productID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully product blocked", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *ProductHandler) UnblockProduct(c *gin.Context) {
	sellerid, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	productID := c.Param("productid")
	err := u.userCase.UnblockProduct(sellerid, productID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully product unblocked", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *ProductHandler) DeleteProduct(c *gin.Context) {
	sellerid, exist := c.MustGet("Sellerid").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	productID := c.Param("productid")
	err := u.userCase.DeleteProduct(sellerid, productID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully product deleted", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *ProductHandler) GetProduct(c *gin.Context) {
	page := c.DefaultQuery("limit", "1")
	limit := c.DefaultQuery("limit", "1")

	product, err := u.userCase.GetAllProducts(page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", product, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *ProductHandler) GetAProduct(c *gin.Context) {
    id := c.Param("productid")

    product, err := u.userCase.GetAProduct(id)
    if err != nil {
        finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
        c.JSON(http.StatusBadRequest, finalResult)
        return
    }

    finalResult := response.Responses(http.StatusOK, "", product, nil)
    c.JSON(http.StatusOK, finalResult)
}

func (u *ProductHandler) GetSellerIProduct(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "1")
	sellerID := c.Param("SellerID") // Assuming the seller ID is in the URL path

	Products, err := u.userCase.GetSellerProducts(page, limit, sellerID)
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
	} else {
		finalResult := response.Responses(http.StatusOK, "", Products, nil)
		c.JSON(http.StatusOK, finalResult)
	}
}

func (u *ProductHandler) EditProduct(c *gin.Context) {

	var edittedProduct requestmodel.EditProduct

	edittedProduct.SellerID = c.MustGet("SellerID").(string)

	if err := c.BindJSON(&edittedProduct); err != nil {
		fmt.Println(err)
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	updatedProduct, err := u.userCase.EditProduct(&edittedProduct)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", updatedProduct, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *ProductHandler) FilterProduct(c *gin.Context) {
	var criterial requestmodel.FilterCriterion

	criterial.Category = c.Query("category")
	criterial.Brand = c.Query("brand")
	criterial.Product = c.Query("product")
	criterial.MinPrice, _ = helper.StringToUintConvertion(c.Query("minprice"))
	criterial.MaxPrice, _ = helper.StringToUintConvertion(c.Query("maxprice"))

	filteredProduct, err := u.userCase.GetProductFilter(&criterial)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", filteredProduct, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
