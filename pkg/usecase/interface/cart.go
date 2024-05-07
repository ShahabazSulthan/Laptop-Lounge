package interfaceUseCase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type ICartUseCase interface {
	CreateCart(*requestmodel.Cart) (*requestmodel.Cart, error)
	DeleteProductFromCart(string, string) error
	QuantityIncriment(string, string) (*requestmodel.Cart, error)
	QuantityDecrease(string, string) (*requestmodel.Cart, error)
	ShowCart(string) (*responsemodel.UserCart, error)
}