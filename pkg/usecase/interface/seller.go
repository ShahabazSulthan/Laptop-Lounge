package interfaceUseCase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type ISellerUseCase interface {
	SellerSignup(*requestmodel.SellerSignup) (*responsemodel.SellerSignupRes, error)
	SellerLogin(*requestmodel.SellerLogin) (*responsemodel.SellerLoginRes, error)
	GetAllSellers(string, string) (*[]responsemodel.SellerDetails, *int, error)
}
