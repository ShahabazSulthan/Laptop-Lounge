package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	useCase interfaceUseCase.IPaymentUseCase
}

func NewPaymentHandler(useCase interfaceUseCase.IPaymentUseCase) *PaymentHandler {
	return &PaymentHandler{useCase: useCase}
}

func (u *PaymentHandler) OnlinePayment(c *gin.Context) {
	userID := c.Query("userID")
	orderID := c.Query("orderID")
	fmt.Println("**", userID, orderID)
	orderDetails, err := u.useCase.OnlinePayment(userID, orderID)
	fmt.Println("--------",orderDetails)
	if err != nil {
		c.HTML(http.StatusBadRequest, "razopay.html", gin.H{"badRequest": "Refine your request"})
	} else {
		c.HTML(http.StatusOK, "razopay.html", orderDetails)
	}
}

func (u *PaymentHandler) VerifyOnlinePayment(c *gin.Context) {
	var onlinePaymentDetails requestmodel.OnlinePaymentVerification

	if err := c.BindJSON(&onlinePaymentDetails); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(onlinePaymentDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	order, err := u.useCase.OnlinePaymentVerification(&onlinePaymentDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", order, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *PaymentHandler) ViewWallet(c *gin.Context) {
	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userWallet, err := u.useCase.GetUserWallet(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", userWallet, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *PaymentHandler) GetWalletTransaction(c *gin.Context) {
	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	walletTransactions, err := u.useCase.GetWalletTransaction(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", walletTransactions, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
