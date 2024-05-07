package interfaceUseCase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type ISellerUseCase interface {
	SellerSignup(*requestmodel.SellerSignup) (*responsemodel.SellerSignupRes, error)
	SellerLogin(*requestmodel.SellerLogin) (*responsemodel.SellerLoginRes, error)
	GetAllSellers(string, string) (*[]responsemodel.SellerDetails, *int, error)
	BlockSeller(string) error
	ActiveSeller(string) error
	GetAllPendingSellers(string, string) (*[]responsemodel.SellerDetails, error)
	FetchSingleSeller(string) (*responsemodel.SellerDetails, error)

	
	GetSellerProfile(string) (*responsemodel.SellerProfile, error)
	UpdateSellerProfile(*requestmodel.SellerEditProfile) (*responsemodel.SellerProfile, error)

	GetSellerDashbord(string) (*responsemodel.DashBord, error)

}
