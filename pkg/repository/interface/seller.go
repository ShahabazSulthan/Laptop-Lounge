package interfaces

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type ISellerRepo interface {
	IsSellerExist(string) (int, error)
	CreateSeller(*requestmodel.SellerSignup) error
	GetHashPassAndStatus(string) (string, string, string, error)
	GetPasswordByMail(string) string
	AllSellers(int, int) (*[]responsemodel.SellerDetails, error)
	SellerCount(chan int)
	BlockSeller(string) error
	UnblockSeller(string) error
	GetPendingSellers(int, int) (*[]responsemodel.SellerDetails, error)
	GetSingleSeller(string) (*responsemodel.SellerDetails, error)

	GetSellerProfile(string) (*responsemodel.SellerProfile, error)
	UpdateSellerProfile(*requestmodel.SellerEditProfile) (*responsemodel.SellerProfile, error)
	UpdateSellerCredit(string, uint) error
	GetSellerCredit(string) (uint, error)

	GetDashBordOrderCount(string, string) (uint, error)
	GetDashBordOrderSum(string, string) (uint, error)
	// GetSellerCredit(string) (uint, error)
	GetLowStokesProduct(string) ([]uint, error)

}
