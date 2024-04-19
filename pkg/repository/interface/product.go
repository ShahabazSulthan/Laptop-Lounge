package interfaces

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IProductRepository interface {
	CreateProduct(*requestmodel.ProductReq) (*responsemodel.ProductRes, error)
	BlockSingleProductBySeller(string, string) error
	UNBlockSingleProductBySeller(string, string) error
	DeleteProductBySeller(string, string) error
	GetAllProduct(int, int) (*[]responsemodel.ProductShowcase, error)
	GetAProducts(string) (*responsemodel.ProductRes, error)
	GetSellerProduct(int, int, string) (*[]responsemodel.ProductShowcase, error)
	UpdateProduct(*requestmodel.EditProduct) (*responsemodel.ProductRes, error)

	GetProductFilter(*requestmodel.FilterCriterion) (*[]responsemodel.FilterProduct, error)
}
