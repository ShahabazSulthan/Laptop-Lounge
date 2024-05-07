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
	if err := c.ShouldBindJSON(&productDetails); err != nil {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

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
	// Extract SellerID and productID from URL parameters
	sellerID := c.Param("SellerID")
	productID := c.Param("productid")

	// Check if SellerID exists in the request
	if sellerID == "" {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	// Block the product using the use case
	err := u.userCase.BlockProduct(sellerID, productID)
	if err != nil {
		finalResult := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalResult)
	} else {
		finalResult := response.Responses(http.StatusOK, "Successfully product blocked", "", nil)
		c.JSON(http.StatusOK, finalResult)
	}
}

func (u *ProductHandler) UnblockProduct(c *gin.Context) {
	sellerID := c.Param("SellerID")
	productID := c.Param("productid")

	// Check if SellerID exists in the request
	if sellerID == "" {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	err := u.userCase.UnblockProduct(sellerID, productID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully product unblocked", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *ProductHandler) DeleteProduct(c *gin.Context) {
	sellerID := c.Param("SellerID")
	productID := c.Param("productid")

	// Check if SellerID exists in the request
	if sellerID == "" {
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	err := u.userCase.DeleteProduct(sellerID, productID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully product deleted", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *ProductHandler) GetProduct(c *gin.Context) {
	
	product, err := u.userCase.GetAllProducts()
	if err != nil {
		// If there's an error, respond with a JSON error message and status 404
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		// If successful, construct the final response using a custom Responses function
		finalResult := response.Responses(http.StatusOK, "", product, nil)

		// Return the final response with status 200
		c.JSON(http.StatusOK, finalResult)
	}
}

func (u *ProductHandler) GetAProduct(c *gin.Context) {
	// Extract the 'productid' parameter from the URL path
	id := c.Param("productid")

	// Call the GetAProduct method from the userCase with the extracted id
	product, err := u.userCase.GetAProduct(id)
	if err != nil {
		// If there's an error, construct a response with a JSON error message and status 400 (Bad Request)
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return // Exit the function to avoid further processing
	}

	// If successful, construct the final response with the product data and status 200 (OK)
	finalResult := response.Responses(http.StatusOK, "", product, nil)
	c.JSON(http.StatusOK, finalResult)
}

func (u *ProductHandler) GetAProductHightoLow(c *gin.Context) {

	product, err := u.userCase.GetAProductHightoLow()
	if err != nil {
		// If there's an error, construct a response with a JSON error message and status 400 (Bad Request)
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return // Exit the function to avoid further processing
	}

	// If successful, construct the final response with the product data and status 200 (OK)
	finalResult := response.Responses(http.StatusOK, "", product, nil)
	c.JSON(http.StatusOK, finalResult)
}

func (u *ProductHandler) GetAProductLowtoHigh(c *gin.Context) {

	product, err := u.userCase.GetAProductLowtoHigh()
	if err != nil {
		// If there's an error, construct a response with a JSON error message and status 400 (Bad Request)
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return // Exit the function to avoid further processing
	}

	// If successful, construct the final response with the product data and status 200 (OK)
	finalResult := response.Responses(http.StatusOK, "", product, nil)
	c.JSON(http.StatusOK, finalResult)
}

func (u *ProductHandler) GetSellerIProduct(c *gin.Context) {
	page := c.Param("page")
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
	var editedProduct requestmodel.EditProduct

	// Get SellerID from URL params
	editedProduct.SellerID = c.Param("SellerID")

	// Bind JSON data to editedProduct
	if err := c.ShouldBindJSON(&editedProduct); err != nil {
		fmt.Println(err)
		finalResult := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
		return
	}

	// Call the use case method to edit the product
	updatedProduct, err := u.userCase.EditProduct(&editedProduct)
	if err != nil {
		finalResult := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalResult)
	} else {
		finalResult := response.Responses(http.StatusOK, "", updatedProduct, nil)
		c.JSON(http.StatusOK, finalResult)
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
