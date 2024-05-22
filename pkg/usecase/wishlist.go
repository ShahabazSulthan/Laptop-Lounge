package usecase

import (
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"errors"
)

type wishlistUseCase struct {
	repo interfaces.IwishlistRepo
}

func NewWishlisttUseCase(repository interfaces.IwishlistRepo) interfaceUseCase.IwishlistRepo {
	return &wishlistUseCase{repo: repository}
}

func (w *wishlistUseCase) AddProductToWishlist(ProductID string, userID string) error {
	err := w.repo.CheckExitWishList(ProductID, userID)
	if err != nil {
		return err
	}
	err = w.repo.AddProductToWishlist(ProductID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (w *wishlistUseCase) ViewUserWishlist(userID string) (*[]responsemodel.ProductShowcase, error) {
	WishedProducts, err := w.repo.GetWishlistsProducts(userID)
	if err != nil {
		return nil, err
	}
	return WishedProducts, nil
}

func (w *wishlistUseCase) RemoveProductFromWishlist(productID string, userID string) error {
	err := w.repo.CheckExitWishList(productID, userID)
	if err == nil {
		return errors.New(`no product found in wishlist with this id`)
	}
	err = w.repo.RemoveProductFromWishliist(productID, userID)
	if err != nil {
		return err
	}
	return nil
}
