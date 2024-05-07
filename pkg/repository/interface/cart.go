package interfaces

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type ICartRepository interface {
	InsertToCart(*requestmodel.Cart) (*requestmodel.Cart, error)
	GetProductPrice(string) (uint, error)
	IsProductExistInCart(string, string) (int, error)
	DeleteProductFromCart(string, string) error
	GetSingleProduct(string, string) (*requestmodel.Cart, error)
	UpdateQuantity(*requestmodel.Cart) (*requestmodel.Cart, error)
	GetCart(string) (*[]responsemodel.CartProduct, error)
	GetNetAmoutOfCart(string, string) (uint, error)
	GetCartCriteria(string) (uint, error)
}
