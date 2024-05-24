package handler

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	"Laptop_Lounge/pkg/models/responseModel/response"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUseCase interfaceUseCase.IAdminUseCAse
}

func NewAdminHandler(useCase interfaceUseCase.IAdminUseCAse) *AdminHandler {
	return &AdminHandler{AdminUseCase: useCase}
}

// @Summary Admin Login
// @Description Using this handler, admins can log in and receive an authentication token.
// @Tags Admins
// @Accept json
// @Produce json
// @Param admin body requestmodel.AdminLoginData true "Admin login details"
// @Success 200 {object} response.Response{data=string} "Successfully logged in. Token returned."
// @Failure 400 {object} response.Response "Bad Request. Invalid input."
// @Failure 401 {object} response.Response "Unauthorized. Authentication failed."
// @Router /admin/login [post]
func (u *AdminHandler) AdminLogin(c *gin.Context) {
	var loginCredential requestmodel.AdminLoginData

	err := c.BindJSON(&loginCredential)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "json is wrong can't bind", nil, err.Error())
		c.JSON(http.StatusUnauthorized, finalReslt)
		return
	}

	data, err := helper.Validation(loginCredential)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.AdminUseCase.AdminLogin(&loginCredential)
	if err != nil {
		finalReslt := response.Responses(http.StatusUnauthorized, "", result, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully login", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Get Admin Dashboard Details
// @Description Retrieve details for the admin. Requires a valid Bearer token.
// @Tags Admins
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Admin details retrieved successfully"
// @Failure 401 {object} response.Response "Unauthorized. Authentication required."
// @Failure 500 {object} response.Response "Internal Server Error."
// @Router /admin/ [get]
func (u *AdminHandler) AdminDashBord(c *gin.Context) {
	result, err := u.AdminUseCase.GetAllSellersDetailAdminDashboard()
	fmt.Println("Errr", result)
	if err != nil {
		finalReslt := response.Responses(http.StatusUnauthorized, "", result, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully login", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
