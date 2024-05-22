package requestmodel

type Cart struct {
	UserID    string `json:"cartid_userid"`
	ProductID string `json:"productid" validate:"required,number"`
	Quantity  uint   `json:"quantity" validate:"required,number,min=1,max=5"`
	Price     uint   `json:"price,omitempty"`
}
