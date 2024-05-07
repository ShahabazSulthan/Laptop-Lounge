package interfaceUseCase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IPaymentUseCase interface {
	OnlinePayment(string, string) (*responsemodel.OnlinePayment, error)
	OnlinePaymentVerification(*requestmodel.OnlinePaymentVerification) (*[]responsemodel.OrderDetails, error)
	GetUserWallet(string) (*responsemodel.UserWallet, error)
	GetWalletTransaction(string) (*[]responsemodel.WalletTransaction, error)
}
