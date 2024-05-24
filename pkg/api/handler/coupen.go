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

// CreateCoupon creates a new coupon by the admin.
// @Summary      Create Coupon (Admin)
// @Description  Create a new coupon by the admin.
// @Tags         Admin Coupons
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Param        coupon body requestmodel.Coupon true "Coupon details to be created"
// @Success      201 {object} response.Response "Coupon created successfully"
// @Failure      400 {object} response.Response "Bad request. Unable to create the coupon."
// @Router       /admin/coupon/ [post]
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

// GetCoupons retrieves a list of coupons for the admin.
// @Summary      Get Coupons (Admin)
// @Description  Retrieve a list of coupons for the admin.
// @Tags         Admin Coupons
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Success      200 {object} response.Response "Coupons retrieved successfully"
// @Failure      400 {object} response.Response "Bad request. Unable to retrieve coupons."
// @Router       /admin/coupon/ [get]
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

// / UnblockCoupon unblocks a coupon by the admin.
// @Summary      Unblock Coupon (Admin)
// @Description  Unblock a coupon by the admin.
// @Tags         Admin Coupons
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Param        couponID path int true "ID of the coupon to be unblocked"
// @Success      200 {object} response.Response "Coupon unblocked successfully"
// @Failure      400 {object} response.Response "Bad request. Please provide a valid coupon ID."
// @Router       /admin/coupon/unblock/{couponID} [patch]
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

// BlockCoupon blocks a coupon by the admin.
// @Summary      Block Coupon (Admin)
// @Description  Block a coupon by the admin.
// @Tags         Admin Coupons
// @Accept       json
// @Produce      json
// @Security     BearerTokenAuth
// @Param        couponID path int true "ID of the coupon to be blocked"
// @Success      200 {object} response.Response "Coupon blocked successfully"
// @Failure      400 {object} response.Response "Bad request. Please provide a valid coupon ID."
// @Router       /admin/coupon/block/{couponID} [patch]
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
