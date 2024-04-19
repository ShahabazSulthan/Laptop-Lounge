package interfaces

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type ICategoryRepository interface {
	InsertCategory(*requestmodel.Category) error
	GetAllCategory(int, int) (*[]responsemodel.CategoryDetails, error)
	EditCategoryName(*requestmodel.CategoryDetails) (*responsemodel.CategoryDetails, error)
	DeleteCategory(string) error

	InsertBrand(*requestmodel.Brand) error
	GetAllBrand(int, int) (*[]responsemodel.BrandRes, error)
	EditBrandName(*requestmodel.BrandDetails) error
	DeleteBrand(string) error

	InsertCategoryOffer(*requestmodel.CategoryOffer) (*responsemodel.CategoryOffer, error)
	ChekSellerHaveCategoryOffer(string, string) (*uint, error)
	ChangeStatus(string, string) (*responsemodel.CategoryOffer, error)
	GetAllCategoryOffers(string) (*[]responsemodel.CategoryOffer, error)
	UpdateCategoryOffer(*requestmodel.EditCategoryOffer) (*responsemodel.CategoryOffer, error)
}
