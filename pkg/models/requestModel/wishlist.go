package requestmodel

type WishlistRequest struct {
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
}
