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

type OrderHandler struct {
	useCase interfaceUseCase.IOrderUseCase
}

func NewOrderHandler(orderUseCase interfaceUseCase.IOrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: orderUseCase}
}

//---------------------------------Create a NewOrder-----------------------------------//

// NewOrder creates a new order
// @Summary Create a new order
// @Description Create a new order with the input payload
// @Tags orders
// @Accept  json
// @Produce  json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param order body requestmodel.Order true "Order data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders [post]
func (u *OrderHandler) NewOrder(c *gin.Context) {

	var order *requestmodel.Order

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		fmt.Println("aaa", finalReslt)
		return
	}

	fmt.Println("rrr", order)
	if err := c.BindJSON(&order); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		fmt.Println("bbb", finalReslt)
		return
	}

	data, err := helper.Validation(*order)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		fmt.Println("ccc", finalReslt)
		return
	}

	order.UserID = userID

	orderDetais, err := u.useCase.NewOrder(order)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		fmt.Println("ddd", finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
		fmt.Println("eee", finalReslt)
	}
}

//---------------------------------Get All Orders-----------------------------------//

// ShowAbstractOrders retrieves all orders
// @Summary Get all orders
// @Description Get a list of all orders for the logged in user
// @Tags orders
// @Accept  json
// @Produce  json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders [get]
func (u *OrderHandler) ShowAbstractOrders(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.OrderShowcase(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------------Get Single Order-----------------------------------//

// SingleOrderDetails retrieves a single order's details
// @Summary Get order details
// @Description Get the details of a single order by its ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param orderItemID path string true "Order Item ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/{orderItemID} [get]
func (u *OrderHandler) SingleOrderDetails(c *gin.Context) {

	orderID, _ := c.Params.Get("orderItemID")

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	orderDetais, err := u.useCase.SingleOrder(orderID, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------------Cancel User Order-----------------------------------//

// CancelUserOrder cancels a user's order
// @Summary Cancel order
// @Description Cancel an order by its ID for the logged-in user
// @Tags orders
// @Accept  json
// @Produce  json
// @Security BearerTokenAuth
// @Security RefreshToken
// @Param orderItemID path string true "Order Item ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/{orderItemID} [delete]
func (u *OrderHandler) CancelUserOrder(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	orderID := c.Param("orderItemID")

	orderDetais, err := u.useCase.CancelUserOrder(orderID, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------------Return User Order-----------------------------------//

// ReturnUserOrder returns a user's order
// @Summary Return order
// @Description Return an order by its ID for the logged-in user
// @Tags orders
// @Accept  json
// @Produce  json
// @Security BearerTokenAuth
// @Security RefreshToken
// @Param orderItemID path string true "Order Item ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/return/{orderItemID} [post]
func (u *OrderHandler) ReturnUserOrder(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	orderID := c.Param("orderItemID")

	orderDetais, err := u.useCase.ReturnUserOrder(orderID, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------Seller Control Orders------------------------------------\\

//---------------------------------Get All Seller Order-----------------------------------//

// GetSellerOrders gets all orders for a seller
// @Summary Get all seller orders
// @Description Get all orders for the specified seller ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param SellerID path string true "Seller ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/seller/{SellerID} [get]
func (u *OrderHandler) GetSellerOrders(c *gin.Context) {

	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	remainingQuery := " IN ('processing','delivered','cancel')"
	orderDetais, err := u.useCase.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// GetSellerOrdersProcessing gets all processing orders for a seller
// @Summary Get seller processing orders
// @Description Get all processing orders for the specified seller ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param SellerID path string true "Seller ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/seller/processing/{SellerID} [get]
func (u *OrderHandler) GetSellerOrdersProcessing(c *gin.Context) {
	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	remainingQuery := " IN ('processing')"
	orderDetais, err := u.useCase.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------------Get Seller Order Delivered-----------------------------------//

// GetSellerOrdersDelivered gets all delivered orders for a seller
// @Summary Get seller delivered orders
// @Description Get all delivered orders for the specified seller ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param SellerID path string true "Seller ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/seller/delivered/{SellerID} [get]
func (u *OrderHandler) GetSellerOrdersDeliverd(c *gin.Context) {
	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	remainingQuery := " IN ('delivered')"
	orderDetais, err := u.useCase.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------------Get Seller Order Cancelled-----------------------------------//

// GetSellerOrdersCancelled gets all cancelled orders for a seller
// @Summary Get seller cancelled orders
// @Description Get all cancelled orders for the specified seller ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param SellerID path string true "Seller ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/seller/cancelled/{SellerID} [get]
func (u *OrderHandler) GetSellerOrdersCancelled(c *gin.Context) {
	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	remainingQuery := " IN ('cancel')"
	orderDetais, err := u.useCase.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------------Seller Conform Delivered-----------------------------------//

// ConfirmDelivered confirms an order as delivered
// @Summary Confirm order delivered
// @Description Confirm an order as delivered by its ID for the specified seller ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param SellerID path string true "Seller ID"
// @Param orderItemID path string true "Order Item ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/seller/{SellerID}/confirm/{orderItemID} [post]
func (u *OrderHandler) ConfirmDeliverd(c *gin.Context) {
	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	orderID := c.Param("orderItemID")
	orderDetais, err := u.useCase.ConfirmDeliverd(sellerID, orderID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------------Seller Cancel Order-----------------------------------//

// CancelOrder cancels an order by the seller
// @Summary Cancel order
// @Description Cancel an order by its ID for the specified seller ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param SellerID path string true "Seller ID"
// @Param orderID path string true "Order ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/seller/{SellerID}/cancel/{orderID} [delete]
func (u *OrderHandler) CancelOrder(c *gin.Context) {

	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	orderID := c.Param("orderID")

	orderDetais, err := u.useCase.CancelOrder(orderID, sellerID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------Sales Report YEAR-MONTH-DAY------------------------------------\\

// SalesReportCustomDays generates a sales report for the past custom number of days for a seller
// @Summary Generate sales report for custom days
// @Description Generate a sales report for the specified seller ID for the past number of days
// @Tags sales
// @Accept  json
// @Produce  json
// @Param SellerID path string true "Seller ID"
// @Param days path int true "Number of days"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /sales/seller/{SellerID}/report/days/{days} [get]
func (u *OrderHandler) SalesReport(c *gin.Context) {

	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	year := c.Query("year")
	month := c.Query("month")
	day := c.Query("day")

	report, err := u.useCase.GetSalesReport(sellerID, year, month, day)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------------Sales Report in Days-----------------------------------//

// SalesReportCustomDays generates a custom sales report for a seller based on the number of days
// @Summary Generate custom sales report
// @Description Generate a custom sales report for the specified seller ID for the given number of days
// @Tags sales
// @Accept  json
// @Produce  json
// @Param SellerID path string true "Seller ID"
// @Param days path int true "Number of days"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /sales/seller/{SellerID}/report/days/{days} [get]
func (u *OrderHandler) SalesReportCustomDays(c *gin.Context) {
	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	day := c.Param("days")
	ans, _ := strconv.Atoi(day)

	report, err := u.useCase.GetSalesReportByDays(sellerID, ans)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//---------------------------------Get Sales Report in Excel-----------------------------------//

// SalesReportXLSX generates a sales report in XLSX format
// @Summary Generate sales report in XLSX
// @Description Generate a sales report in XLSX format for the specified seller ID
// @Tags sales
// @Accept  json
// @Produce  json
// @Param SellerID path string true "Seller ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /sales/seller/{SellerID}/report/xlsx [get]
func (u *OrderHandler) SalesReportXlSX(c *gin.Context) {
	sellerID := c.Param("SellerID")
	if sellerID == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.GenerateXlOfSalesReport(sellerID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "sales report create succesfully", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------Invoice------------------------------------\\

// GetInvoice generates an invoice for an order
// @Summary Generate invoice
// @Description Generate an invoice for the specified order item ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Security BearerTokenAuth
// @Security RefreshToken
// @Param orderItemID path string true "Order Item ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/invoice/{orderItemID} [get]
func (u *OrderHandler) GetInvoice(c *gin.Context) {

	orderItemID := c.Param("orderItemID")
	pdfLink, err := u.useCase.OrderInvoiceCreation(orderItemID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "invoice successfully created", pdfLink, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
