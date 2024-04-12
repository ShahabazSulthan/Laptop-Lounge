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
}
