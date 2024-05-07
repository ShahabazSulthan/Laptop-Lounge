package interfaces

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IPaymentRepository interface {
	CreateOrUpdateWallet(string, uint) (uint, error)
	OnlinePayment(string, string) (*responsemodel.OnlinePayment, error)
	GetFinalPriceByorderID(string) (uint, error)
	UpdateOnlinePaymentSucess(string) (*[]responsemodel.OrderDetails, error)
	GetWallet(string) (*responsemodel.UserWallet, error)
	UpdateWalletReduceBalance(string, uint) error
	GetWalletbalance(userID string) (*uint, error)
	WalletTransaction(requestmodel.WalletTransaction) error
	GetWalletTransaction(string) (*[]responsemodel.WalletTransaction, error)
}
