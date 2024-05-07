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

type CouponHandler struct {
	useCase interfaceUseCase.ICouponUseCase
}

func NewCouponHandler(useCase interfaceUseCase.ICouponUseCase) *CouponHandler {
	return &CouponHandler{useCase: useCase}
}

func (u *CouponHandler) CreateCoupon(c *gin.Context) {
	var newCoupon requestmodel.Coupon

	if err := c.BindJSON(&newCoupon); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(newCoupon)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	coupon, err := u.useCase.CreateCoupon(&newCoupon)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", coupon, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *CouponHandler) GetCoupons(c *gin.Context) {
	coupon, err := u.useCase.GetCoupons()
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", coupon, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *CouponHandler) UnblockCoupon(c *gin.Context) {
	couponID := c.Param("couponID")
	coupon, err := u.useCase.UpdateCouponStatus(couponID, "active")
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", coupon, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *CouponHandler) BlockCoupon(c *gin.Context) {
	couponID := c.Param("couponID")
	coupon, err := u.useCase.UpdateCouponStatus(couponID, "block")
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", coupon, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}