package repository

import (
	"errors"

	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	interfaces "Laptop_Lounge/pkg/repository/interface"

	"gorm.io/gorm"
)

type WishlistRepo struct {
	DB *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) interfaces.IwishlistRepo {
	return &WishlistRepo{DB: db}
}

func (w *WishlistRepo) CheckExitWishList(ProductID, UserID string) error {
	var count int64
	query := "SELECT COUNT(*) FROM wishlists WHERE product_id = ? AND user_id = ?"
	if err := w.DB.Raw(query, ProductID, UserID).Scan(&count).Error; err != nil {
		return errors.New("encountered an issue while checking wishlist")
	}

	if count > 0 {
		return errors.New("product is already in the wishlist")
	}
	return nil
}

func (w *WishlistRepo) AddProductToWishlist(ProductID, UserID string) error {
	query := "INSERT INTO wishlists (product_id, user_id) VALUES (?, ?)"
	if err := w.DB.Exec(query, ProductID, UserID).Error; err != nil {
		return errors.New("encountered an issue while inserting into wishlist")
	}
	return nil
}

func (w *WishlistRepo) GetWishlistsProducts(userID string) (*[]responsemodel.ProductShowcase, error) {
	var products []responsemodel.ProductShowcase
	query := `
		SELECT products.id, products.model_name, products.mrp, products.discount, products.sale_price, products.seller_id, products.units, products.image_url
		FROM products 
		INNER JOIN wishlists ON products.id = wishlists.product_id
		WHERE wishlists.user_id = ?
	`
	result := w.DB.Raw(query, userID).Scan(&products)
	if result.Error != nil {
		return nil, errors.New("encountered an issue while fetching products from wishlist")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &products, nil
}

func (w *WishlistRepo) RemoveProductFromWishliist(ProductID, UserID string) error {
	query := "DELETE FROM wishlists WHERE product_id = ? AND user_id = ?"
	result := w.DB.Exec(query, ProductID, UserID)
	if result.Error != nil {
		return errors.New("encountered an issue while deleting from wishlist")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}
