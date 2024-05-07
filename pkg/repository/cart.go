package repository

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"errors"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.ICartRepository {
	return &cartRepository{DB: db}
}

func (d *cartRepository) IsProductExistInCart(ProductID string, userID string) (int, error) {
	var ProductCount int

	query := "SELECT count(*) FROM carts WHERE product_id=? AND user_id=? AND status='active' "
	result := d.DB.Raw(query, ProductID, userID).Scan(&ProductCount)
	if result.Error != nil {
		return 0, errors.New("face some issue while finding Product is exist in cart")
	}
	return ProductCount, nil
}

func (d *cartRepository) InsertToCart(cart *requestmodel.Cart) (*requestmodel.Cart, error) {

	query := "INSERT INTO carts (user_id, product_id, quantity) VALUES (?, ?,  ?)   RETURNING *;"
	result := d.DB.Raw(query, cart.UserID, cart.ProductID, cart.Quantity).Scan(&cart)

	if result.Error != nil {
		return nil, errors.New("face some issue while Product insert to cart ")
	}
	if result.RowsAffected == 0 {

		return nil, resCustomError.ErrNoRowAffected
	}
	return cart, nil
}

func (d *cartRepository) GetProductPrice(productID string) (uint, error) {
	var price uint
	query := "SELECT sale_price FROM products WHERE id= ? AND status = 'active'"
	result := d.DB.Raw(query, productID).Scan(&price)
	if result.Error != nil {
		return 0, errors.New("face some issue while get user profile ")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return price, nil
}

func (d *cartRepository) DeleteProductFromCart(ProductID string, userID string) error {

	query := "UPDATE carts SET status='delete' WHERE product_id = ? AND user_id= ? AND status= 'active'"
	result := d.DB.Exec(query, ProductID, userID)
	if result.Error != nil {
		return errors.New("face some issue while delete Product in cart")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}

func (d *cartRepository) GetSingleProduct(ProductID string, userID string) (*requestmodel.Cart, error) {

	var singleProduct *requestmodel.Cart
	query := "SELECT * FROM carts WHERE user_id=? AND product_id=? AND status='active'"
	result := d.DB.Raw(query, userID, ProductID).Scan(&singleProduct)
	if result.Error != nil {
		return nil, errors.New("face some issue while fetching product in cart")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return singleProduct, nil
}

func (d *cartRepository) UpdateQuantity(cart *requestmodel.Cart) (*requestmodel.Cart, error) {

	var singleProduct *requestmodel.Cart
	query := "UPDATE carts SET quantity= ? WHERE user_id=? AND product_id= ? AND status='active' RETURNING* ;"
	result := d.DB.Raw(query, cart.Quantity, cart.UserID, cart.ProductID).Scan(&singleProduct)
	if result.Error != nil {
		return nil, errors.New("face some issue while quantity updating cart")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return singleProduct, nil
}

func (d *cartRepository) GetCart(userID string) (*[]responsemodel.CartProduct, error) {

	var cartView *[]responsemodel.CartProduct
	query := "SELECT * FROM carts INNER JOIN products ON id=product_id LEFT JOIN category_offers ON category_offers.seller_id=products.seller_id AND category_offers.category_id=products.category_id AND category_offers.status='active' AND category_offers.end_date>now() WHERE carts.user_id=? AND carts.status='active'"
	result := d.DB.Raw(query, userID).Scan(&cartView)
	if result.Error != nil {
		return nil, errors.New("face some issue while  get cart")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user have no cart")
	}
	return cartView, nil
}

func (d *cartRepository) GetNetAmoutOfCart(userID string, productID string) (uint, error) {

	var NetCart uint
	query := "SELECT sale_price FROM carts INNER JOIN products ON id=product_id WHERE carts.user_id=? AND id= ? AND carts.status='active'"
	result := d.DB.Raw(query, userID, productID).Scan(&NetCart)
	if result.Error != nil {
		return 0, errors.New("user have no Cart")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("user have no cart")
	}
	return NetCart, nil
}

func (d *cartRepository) GetCartCriteria(userID string) (uint, error) {

	var count uint
	query := "SELECT SUM(quantity) FROM carts WHERE user_id=? AND status='active'"
	result := d.DB.Raw(query, userID)
	result.Row().Scan(&count)
	if result.Error != nil {
		return 0, errors.New("face some issue while  get cart")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return count, nil
}
