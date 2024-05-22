package requestmodel

import responsemodel "Laptop_Lounge/pkg/models/responseModel"

type Order struct {
	ID             uint                        `json:"-"`
	UserID         string                      `json:"userid" validate:"required,numeric"`
	AddressID      string                      `json:"addressid" validate:"required,numeric"`
	Payment        string                      `json:"paymentMethod" validate:"required,alpha,uppercase"`
	Coupon         string                      `json:"couponid"`
	OrderIDRazopay string                      `json:"-"`
	FinalPrice     uint                        `json:"-"`
	CouponDiscount uint                        `json:"-"`
	OrderStatus    string                      `json:"-"`
	PaymentStatus  string                      `json:"-"`
	Cart           []responsemodel.CartProduct `json:"-"`
}

type OnlinePaymentVerification struct {
	PaymentID string `json:"payment_id" validate:"required"`
	OrderID   string `json:"order_id" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}
