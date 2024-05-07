package requestmodel

type Cart struct {
	UserID    string `json:"cartid userid"`
	ProductID string `json:"productid" validate:"required,number"`
	Quantity  uint   `json:"quantity"`
	Price     uint   `json:"price,omitempty"`
}
