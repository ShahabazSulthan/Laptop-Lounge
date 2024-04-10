package interfaceUseCase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IAdminUseCAse interface {
	AdminLogin(*requestmodel.AdminLoginData) (*responsemodel.AdminLoginres, error)
	GetAllSellersDetailAdminDashboard() (*responsemodel.AdminDashBoard, error)
}
