package usecase

import (
	"Laptop_Lounge/pkg/config"
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"errors"
	"fmt"
)

type productUseCase struct {
	repo interfaces.IProductRepository
	s3   config.S3Bucket
}

func NewProductUseCase(repository interfaces.IProductRepository, s3aws *config.S3Bucket) interfaceUseCase.IProductUseCase {
	return &productUseCase{repo: repository, s3: *s3aws}
}

func (d *productUseCase) AddProduct(product *requestmodel.ProductReq) (*responsemodel.ProductRes, error) {

	// Calculate discounted price
	discountedPrice := helper.FindDiscount(float64(product.Mrp), float64(product.Discount))
	product.SalePrice = discountedPrice

	// Create product in repository
	products, err := d.repo.CreateProduct(product)
	if err != nil {
		return nil, fmt.Errorf("error creating product: %v", err)
	}

	return products, nil
}

func (r *productUseCase) BlockProduct(sellerID string, productID string) error {
	err := r.repo.BlockSingleProductBySeller(sellerID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (r *productUseCase) UnblockProduct(sellerID string, productID string) error {
	err := r.repo.UNBlockSingleProductBySeller(sellerID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (r *productUseCase) DeleteProduct(sellerID string, productID string) error {
	err := r.repo.DeleteProductBySeller(sellerID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (r *productUseCase) GetAllProducts() (*[]responsemodel.ProductShowcase, error) {

	Products, err := r.repo.GetAllProduct()
	if err != nil {
		return nil, err
	}

	for i, product := range *Products {
		if product.CategoryDiscount != 0 {
			(*Products)[i].NetDiscount = product.Discount + product.CategoryDiscount
			(*Products)[i].PriceAfterApplyCategoryDiscount = helper.FindDiscount(float64(product.Mrp), float64((*Products)[i].NetDiscount))
		}
	}
	return Products, nil
}

func (r *productUseCase) GetAProduct(productID string) (*responsemodel.ProductRes, error) {
	product, err := r.repo.GetAProducts(productID)
	if err != nil {
		return nil, fmt.Errorf("error getting product: %w", err) // Wrap the error for more context
	}

	if product.CategoryDiscount != 0 {
		product.NetDiscount = product.CategoryDiscount + product.Discount
		product.FinalPrice = helper.FindDiscount(float64(product.Mrp), float64(product.NetDiscount))
	}
	return product, nil
}

func (r *productUseCase) GetAProductHightoLow() (*[]responsemodel.ProductShowcase, error) {
	Products, err := r.repo.GetAProductHightoLow()
	if err != nil {
		return nil, err
	}

	for i, product := range *Products {
		if product.CategoryDiscount != 0 {
			(*Products)[i].NetDiscount = product.Discount + product.CategoryDiscount
			(*Products)[i].PriceAfterApplyCategoryDiscount = helper.FindDiscount(float64(product.Mrp), float64((*Products)[i].NetDiscount))
		}
	}
	return Products, nil
}

func (r *productUseCase) GetAProductLowtoHigh() (*[]responsemodel.ProductShowcase, error) {
	Products, err := r.repo.GetAProductLowtoHigh()
	if err != nil {
		return nil, err
	}

	for i, product := range *Products {
		if product.CategoryDiscount != 0 {
			(*Products)[i].NetDiscount = product.Discount + product.CategoryDiscount
			(*Products)[i].PriceAfterApplyCategoryDiscount = helper.FindDiscount(float64(product.Mrp), float64((*Products)[i].NetDiscount))
		}
	}
	return Products, nil
}

func (r *productUseCase) GetAProductAtoZ() (*[]responsemodel.ProductShowcase, error) {
	Products, err := r.repo.GetAProductAtoZ()
	if err != nil {
		return nil, err
	}

	for i, product := range *Products {
		if product.CategoryDiscount != 0 {
			(*Products)[i].NetDiscount = product.Discount + product.CategoryDiscount
			(*Products)[i].PriceAfterApplyCategoryDiscount = helper.FindDiscount(float64(product.Mrp), float64((*Products)[i].NetDiscount))
		}
	}
	return Products, nil
}

func (r *productUseCase) GetAProductZtoA() (*[]responsemodel.ProductShowcase, error) {
	Products, err := r.repo.GetAProductZtoA()
	if err != nil {
		return nil, err
	}

	for i, product := range *Products {
		if product.CategoryDiscount != 0 {
			(*Products)[i].NetDiscount = product.Discount + product.CategoryDiscount
			(*Products)[i].PriceAfterApplyCategoryDiscount = helper.FindDiscount(float64(product.Mrp), float64((*Products)[i].NetDiscount))
		}
	}
	return Products, nil
}

func (r *productUseCase) GetSellerProducts(page string, limit string, sellerID string) (*[]responsemodel.ProductShowcase, error) {

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	Products, err := r.repo.GetSellerProduct(offSet, limits, sellerID)
	if err != nil {
		return nil, err
	}

	for i, product := range *Products {
		if product.CategoryDiscount != 0 {
			(*Products)[i].NetDiscount = product.Discount + product.CategoryDiscount
			(*Products)[i].PriceAfterApplyCategoryDiscount = helper.FindDiscount(float64(product.Mrp), float64((*Products)[i].NetDiscount))
		}
	}
	return Products, nil
}

func (r *productUseCase) EditProduct(EditProduct *requestmodel.EditProduct) (*responsemodel.ProductRes, error) {

	Product, err := r.repo.GetAProducts(EditProduct.ID)
	if err != nil {
		return nil, err
	}
	if Product.SellerID != EditProduct.SellerID {
		return nil, resCustomError.ErrNoRowAffected
	}

	// fill data if it's empty
	if EditProduct.Units == 0 {
		EditProduct.Units = Product.Units
	}

	if EditProduct.Discount == 0 {
		EditProduct.Discount = Product.Discount
	}

	if EditProduct.Mrp == 0 {
		EditProduct.Mrp = Product.Mrp
	}

	if EditProduct.Saleprice == 0 {
		EditProduct.Saleprice = Product.SalePrice
	}

	// modify data to reach my criteria
	if EditProduct.Mrp != 0 {
		EditProduct.Saleprice = helper.FindDiscount(float64(EditProduct.Mrp), float64(Product.Discount))
	}

	if EditProduct.Discount != 0 {
		if EditProduct.Discount > 99 {
			return nil, errors.New("discount must be 1 to 99")
		}
		EditProduct.Saleprice = helper.FindDiscount(float64(Product.Mrp), float64(EditProduct.Discount))
	}

	if EditProduct.Mrp != 0 && EditProduct.Discount != 0 {
		if EditProduct.Discount > 99 {
			return nil, errors.New("discount must be 0 to 99")
		}
		EditProduct.Saleprice = helper.FindDiscount(float64(EditProduct.Mrp), float64(EditProduct.Discount))
	}

	updatedData, err := r.repo.UpdateProduct(EditProduct)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}

func (r *productUseCase) GetProductFilter(criterion *requestmodel.FilterCriterion) (*[]responsemodel.FilterProduct, error) {
	filteredProduct, err := r.repo.GetProductFilter(criterion)
	if err != nil {
		return nil, err
	}
	for i, product := range *filteredProduct {
		if product.CategoryDiscount != 0 {
			(*filteredProduct)[i].NetDiscount = product.Discount + product.CategoryDiscount
			(*filteredProduct)[i].PriceAfterApplyCategoryDiscount = helper.FindDiscount(float64(product.Mrp), float64((*filteredProduct)[i].NetDiscount))
		}
	}
	return filteredProduct, nil
}
