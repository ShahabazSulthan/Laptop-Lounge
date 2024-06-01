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

// AddProduct adds a new product for a specific seller.
// @Summary      Add Product
// @Description  Add a new product from the seller.
// @Tags         Seller Products
// @Accept       json
// @Produce      json
// @Param        SellerID      path      string                 true  "ID of the seller"
// @Param        product       body      requestmodel.ProductReq true  "Product details for adding"
// @Success      200           {object}  response.Response        "Successfully added the product"
// @Failure      400           {object}  response.Response        "Bad request"
// @Router       /seller/products/{SellerID} [post]
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

// BlockProduct blocks a product from being displayed.
// @Summary      Block Product
// @Description  Block a product from being displayed.
// @Tags         Seller Products
// @Accept       json
// @Produce      json
// @Param        SellerID   path     string                true  "Seller ID in the URL path"
// @Param        productid  path     string                true  "Product ID in the URL path"
// @Success      200        {object} response.Response     "Successfully blocked the product"
// @Failure      400        {object} response.Response     "Bad request"
// @Router       /seller/products/{SellerID}/{productid}/block [patch]
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

// UnblockProduct unblocks a product for display.
// @Summary      Unblock Product
// @Description  Unblock a product for display.
// @Tags         Seller Products
// @Accept       json
// @Produce      json
// @Param        SellerID   path     string                true  "Seller ID in the URL path"
// @Param        productid  path     string                true  "Product ID in the URL path"
// @Success      200        {object} response.Response     "Successfully unblocked the product"
// @Failure      400        {object} response.Response     "Bad request"
// @Router       /seller/products/{SellerID}/{productid}/unblock [patch]
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

// DeleteProduct deletes a product by ID.
// @Summary      Delete Product
// @Description  Delete a product by ID.
// @Tags         Seller Products
// @Accept       json
// @Produce      json
// @Param        SellerID   path     string                true  "Seller ID in the URL path"
// @Param        productid  path     string                true  "Product ID in the URL path"
// @Success      200        {object} response.Response     "Successfully deleted the product"
// @Failure      400        {object} response.Response     "Bad request"
// @Router       /seller/products/{SellerID}/{productid} [delete]
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

// GetProduct retrieves a list of available products.
// @Summary      Get Available Laptops
// @Description  Retrieve a list of products.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        page   query    int     false  "Page number"                 default(1)
// @Param        limit  query    int     false  "Number of items per page"    default(5)
// @Success      200    {object} response.Response "Successfully retrieved products"
// @Failure      400    {object} response.Response "Bad request"
// @Router       /products [get]
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

// /	@Summary		Get Seller Product
//
//	@Description	Retrieve details of a single seller product.
//	@Tags			Seller Products
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Security		Refreshtoken
//	@Param			productid	path		string				true	"Product ID in the URL path"
//	@Success		200	{object}	response.Response	"Successfully retrieved the seller product"
//	@Failure		400	{object}	response.Response	"Bad request"
//	@Router			/product/{productid} [get]
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

// @Summary		Get High To Low Price
// @Description	Retrieve a list of products sorted by price from high to low.
// @Tags			Search
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Response	"Successfully retrieved products"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router		    /product/HighToLow [get]
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

// @Summary		Get Low To High Price
// @Description	Retrieve a list of products sorted by price from low to high.
// @Tags			Search
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Response	"Successfully retrieved products"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router		    /product/LowToHigh [get]
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

// @Summary		Get A To Z Product Name
// @Description	Retrieve a list of products sorted by name from A to Z.
// @Tags			Search
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Response	"Successfully retrieved products"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router		    /product/AtoZ [get]
func (u *ProductHandler) GetAProductAtoZ(c *gin.Context) {

	product, err := u.userCase.GetAProductAtoZ()
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

// @Summary		Get Z To A Product Name
// @Description	Retrieve a list of products sorted by name from Z to A.
// @Tags			Search
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Response	"Successfully retrieved products"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router		    /product/ZtoA [get]
func (u *ProductHandler) GetAProductZtoA(c *gin.Context) {

	product, err := u.userCase.GetAProductZtoA()
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

// @Summary		Get Seller Products
// @Description	Retrieve a list of seller products with pagination.
// @Tags			Seller Products
// @Accept			json
// @Produce		json
// @Param			SellerID	path		string	true	"Seller ID"
// @Param			page		path		int		true	"Page number"
// @Param			limit		query		int		false	"Number of items per page"	default(5)
// @Success		200		{object}	response.Response	"Successfully retrieved seller products"
// @Failure		400		{object}	response.Response	"Bad request"
// @Router			/seller/products/seller/{SellerID}/{page} [get]
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

// @Summary		Edit Seller Product
// @Description	Edit details of a seller product.
// @Tags			Seller Products
// @Accept			json
// @Produce		json
// @Param			SellerID	path		string						true	"Seller ID"
// @Param			product		body		requestmodel.EditProduct	true	"Updated product details"
// @Success		200		{object}	response.Response			"Successfully edited the seller product"
// @Failure		400		{object}	response.Response			"Bad request"
// @Router			/seller/products/{SellerID} [patch]
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

// @Summary		Filter Products
// @Description	Filter products based on category, brand, product name, and price range.
// @Tags			Search
// @Accept			json
// @Produce		json
// @Param			category	query		string				false	"Category filter"
// @Param			brand		query		string				false	"Brand filter"
// @Param			product		query		string				false	"Product name filter"
// @Param			minprice	query		int					false	"Minimum price filter"
// @Param			maxprice	query		int					false	"Maximum price filter"
// @Success		200			{object}	response.Response	"Products filtered successfully"
// @Failure		400			{object}	response.Response	"Bad request. Please provide valid filter criteria."
// @Router			/filter [get]
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
