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
