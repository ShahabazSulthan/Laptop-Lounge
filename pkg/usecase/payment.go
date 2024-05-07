package usecase

import (
	"Laptop_Lounge/pkg/config"
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"Laptop_Lounge/pkg/service"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"errors"
	"fmt"
)

type paymentUseCase struct {
	repo    interfaces.IPaymentRepository
	razopay *config.Razopay
}

func NewPaymentUseCase(repository interfaces.IPaymentRepository, razopay *config.Razopay) interfaceUseCase.IPaymentUseCase {
	return &paymentUseCase{repo: repository, razopay: razopay}
}

func (r *paymentUseCase) OnlinePayment(userID, orderID string) (*responsemodel.OnlinePayment, error) {
	fmt.Println("---",r)
	paymentDetails, err := r.repo.OnlinePayment(userID, orderID)
	if err != nil {
		return nil, err
	}
	fmt.Println("&&", paymentDetails)

	paymentDetails.FinalPrice, err = r.repo.GetFinalPriceByorderID(orderID)
	if err != nil {
		return nil, err
	}
	return paymentDetails, nil
}

func (r *paymentUseCase) OnlinePaymentVerification(details *requestmodel.OnlinePaymentVerification) (*[]responsemodel.OrderDetails, error) {
	result := service.VerifyPayment(details.OrderID, details.PaymentID, details.Signature, r.razopay.RazopaySecret)
	if !result {
		return nil, errors.New("payment is unsuccessful")
	}

	orders, err := r.repo.UpdateOnlinePaymentSucess(details.OrderID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *paymentUseCase) GetUserWallet(userID string) (*responsemodel.UserWallet, error) {
	userWallet, err := r.repo.GetWallet(userID)
	if err != nil {
		return nil, err
	}
	return userWallet, err
}

func (r *paymentUseCase) GetWalletTransaction(userID string) (*[]responsemodel.WalletTransaction, error) {
	transaction, err := r.repo.GetWalletTransaction(userID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}