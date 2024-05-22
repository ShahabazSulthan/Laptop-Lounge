package interfaces

import responsemodel "Laptop_Lounge/pkg/models/responseModel"

type IwishlistRepo interface {
	CheckExitWishList(string, string) error
	AddProductToWishlist(string, string) error
	GetWishlistsProducts(string) (*[]responsemodel.ProductShowcase, error)
	RemoveProductFromWishliist(string, string) error
}
