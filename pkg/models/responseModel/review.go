package responsemodel

type ReviewResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}
