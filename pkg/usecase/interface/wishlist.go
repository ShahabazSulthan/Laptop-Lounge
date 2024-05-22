package interfaceUseCase

import responsemodel "Laptop_Lounge/pkg/models/responseModel"

type IwishlistRepo interface {
	AddProductToWishlist(string, string) error
	ViewUserWishlist(string) (*[]responsemodel.ProductShowcase, error)
	RemoveProductFromWishlist(string, string) error
}
