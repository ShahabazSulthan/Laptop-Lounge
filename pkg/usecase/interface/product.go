package interfaceUseCase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IProductUseCase interface {
	AddProduct(*requestmodel.ProductReq) (*responsemodel.ProductRes, error)
	BlockProduct(string, string) error
	UnblockProduct(string, string) error
	DeleteProduct(string, string) error
	GetAllProducts() (*[]responsemodel.ProductShowcase, error)
	GetAProduct(string) (*responsemodel.ProductRes, error)
	GetAProductHightoLow() (*[]responsemodel.ProductShowcase, error)
	GetAProductLowtoHigh() (*[]responsemodel.ProductShowcase, error)
	GetAProductAtoZ() (*[]responsemodel.ProductShowcase, error)
	GetAProductZtoA() (*[]responsemodel.ProductShowcase, error)
	GetSellerProducts(string, string, string) (*[]responsemodel.ProductShowcase, error)
	EditProduct(*requestmodel.EditProduct) (*responsemodel.ProductRes, error)

	GetProductFilter(*requestmodel.FilterCriterion) (*[]responsemodel.FilterProduct, error)
}
