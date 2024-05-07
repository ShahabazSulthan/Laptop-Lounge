package usecase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"errors"
	"fmt"
)

type cartUseCase struct {
	repo interfaces.ICartRepository
}

func NewCartUseCase(repository interfaces.ICartRepository) interfaceUseCase.ICartUseCase {
	return &cartUseCase{repo: repository}
}

func (r *cartUseCase) CreateCart(cart *requestmodel.Cart) (*requestmodel.Cart, error) {

	count, err := r.repo.IsProductExistInCart(cart.ProductID, cart.UserID)
	if err != nil {
		return nil, err
	}

	if count >= 1 {
		return nil, errors.New("product alrady exist in cart now you can purchase")
	}

	cart.Quantity = 1

	inserCart, err := r.repo.InsertToCart(cart)
	if err != nil {
		return nil, err
	}
	return inserCart, nil
}

func (r *cartUseCase) DeleteProductFromCart(ProductID string, userID string) error {
	err := r.repo.DeleteProductFromCart(ProductID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *cartUseCase) QuantityIncriment(ProductID string, userID string) (*requestmodel.Cart, error) {

	singleProduct, err := r.repo.GetSingleProduct(ProductID, userID)
	if err != nil {
		return nil, err
	}

	singleProduct.Quantity += 1

	singleProduct, err = r.repo.UpdateQuantity(singleProduct)
	if err != nil {
		return nil, err
	}
	return singleProduct, nil
}

func (r *cartUseCase) QuantityDecrease(ProductID string, userID string) (*requestmodel.Cart, error) {

	singleProduct, err := r.repo.GetSingleProduct(ProductID, userID)
	if err != nil {
		return nil, err
	}

	if singleProduct.Quantity == 1 {
		return singleProduct, errors.New("reach the maximum limit")
	}

	singleProduct.Quantity -= 1

	singleProduct, err = r.repo.UpdateQuantity(singleProduct)
	if err != nil {
		return nil, err
	}
	return singleProduct, nil
}

func (r *cartUseCase) ShowCart(userID string) (*responsemodel.UserCart, error) {

	cart := &responsemodel.UserCart{}

	quantity, err := r.repo.GetCartCriteria(userID)
	if err != nil {
		return nil, err
	}

	cart.ProductCount = quantity
	cart.UserID = userID

	CartProducts, err := r.repo.GetCart(userID)
	if err != nil {
		return nil, err
	}

	for i, product := range *CartProducts {

		// price, err := r.repo.GetNetAmoutOfCart(userID, inventory.InventoryID)
		// if err != nil {
		// 	return nil, err
		// }
		(*CartProducts)[i].FinalPrice = helper.FindDiscount(float64(product.Price), float64(product.Discount+product.CategoryDiscount)) * product.Quantity
		fmt.Println("**", (*CartProducts)[i].FinalPrice)
		cart.TotalPrice += (*CartProducts)[i].FinalPrice
	}

	cart.Cart = *CartProducts
	return cart, nil
}
