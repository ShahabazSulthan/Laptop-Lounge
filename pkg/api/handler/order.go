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

type OrderHandler struct {
	useCase interfaceUseCase.IOrderUseCase
}

func NewOrderHandler(orderUseCase interfaceUseCase.IOrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: orderUseCase}
}

func (u *OrderHandler) NewOrder(c *gin.Context) {

	var order *requestmodel.Order

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	if err := c.BindJSON(&order); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(*order)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	order.UserID = userID

	orderDetais, err := u.useCase.NewOrder(order)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}


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


func (u *OrderHandler) SingleOrderDetails(c *gin.Context) {

	orderID, _ := c.Params.Get("orderID")

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


func (u *OrderHandler) CancelUserOrder(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	orderID := c.Param("orderID")

	orderDetais, err := u.useCase.CancelUserOrder(orderID, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *OrderHandler) ReturnUserOrder(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	orderID := c.Param("orderID")

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

func (u *OrderHandler) GetSellerOrders(c *gin.Context) {
	
	sellerID := c.Param("SellerID")
    if sellerID == "" {
    finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
    c.JSON(http.StatusBadRequest, finalReslt)
    return
}


	remainingQuery := " IN ('processing','delivered','cancelled')"
	orderDetais, err := u.useCase.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}


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

func (u *OrderHandler) GetSellerOrdersCancelled(c *gin.Context) {
	sellerID := c.Param("SellerID")
    if sellerID == "" {
    finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
    c.JSON(http.StatusBadRequest, finalReslt)
    return
}

	remainingQuery := " IN ('cancelled')"
	orderDetais, err := u.useCase.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}


func (u *OrderHandler) ConfirmDeliverd(c *gin.Context) {
	sellerID := c.Param("SellerID")
    if sellerID == "" {
    finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
    c.JSON(http.StatusBadRequest, finalReslt)
    return
}

	orderID := c.Param("orderID")
	orderDetais, err := u.useCase.ConfirmDeliverd(sellerID, orderID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

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

// ------------------------------------------Sales Report------------------------------------\\


func (u *OrderHandler) SalesReport(c *gin.Context) {

	sellerID := c.Param("SellerID")
    if sellerID == "" {
    finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
    c.JSON(http.StatusBadRequest, finalReslt)
    return
}
	year := c.Param("year")
	month := c.Param("month")
	day := c.Param("day")

	report, err := u.useCase.GetSalesReport(sellerID, year, month, day)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *OrderHandler) SalesReportCustomDays(c *gin.Context) {
	sellerID := c.Param("SellerID")
    if sellerID == "" {
    finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetSellerIDinContexr)
    c.JSON(http.StatusBadRequest, finalReslt)
    return
}
	day := c.Param("days")

	report, err := u.useCase.GetSalesReportByDays(sellerID, day)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}


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

func (u *OrderHandler) GetInvoice(c *gin.Context) {

	orderItemID := c.Param("OrderID")
	pdfLink, err := u.useCase.OrderInvoiceCreation(orderItemID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "invoice successfully created", pdfLink, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}