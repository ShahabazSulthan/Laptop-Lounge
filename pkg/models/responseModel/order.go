package responsemodel

import "time"

type OrderShowcase struct {
	Model_name    string `json:"model_name" validate:"required,min=3,max=100"`
	SingleOrderID string `json:"singleorderid" gorm:"column:item_id"`
	ID            string `gorm:"column:order_id" json:"orderID" validate:"required,number"`
	UserID        string `gorm:"column:user_id" json:"userid"`
	SellerID      string `json:"seller_id" gorm:"column:seller_id"`
	ProductID     string `gorm:"column:product_id" json:"productid"`
	Price         uint   `json:"total-amout"`
	Saleprice     uint   `json:"userPayableAmount" gorm:"column:payable_amount" `
	OrderStatus   string `json:"orderstatus,omitempty"`
	PaymentStatus string `json:"paymentStatus,omitempty"`
	Quantity      uint   `json:"quantity"`
	ImageURL      string `json:"imageURL"`
}

type OrderDetails struct {
	ID            string `json:"orderid" gorm:"column:order_id"`
	ItemID        string `json:"itemID" gorm:"column:item_id"`
	UserID        string `json:"user_id" gorm:"column:user_id"`
	Address       string `json:"address_id" gorm:"column:address_id"`
	Payment       string `json:"payment_method" gorm:"column:payment_method"`
	SellerID      string `json:"seller_id" gorm:"column:seller_id"`
	ProductID     string `json:"productid" gorm:"column:product_id"`
	Quantity      uint   `json:"quantity"`
	Price         uint   `json:"sale_price"`
	Saleprice     uint   `json:"userPayableAmount" gorm:"column:payable_amount"`
	OrderStatus   string `json:"orderstatus,omitempty" gorm:"column:order_status"`
	PaymentStatus string `json:"paymentStatus,omitempty" gorm:"column:payment_status"`
	WalletBalance uint   `json:"walletBelance,omitempty" gorm:"-"`
}

type OrderProducts struct {
	ItemID           string    `json:"itemID"`
	OrderID          string    `json:"parentOrderID"`
	ProductID        string    `json:"proSductID"`
	SellerID         string    `json:"sellerID"`
	CategoryID       string    `json:"categoryID"`
	Discount         uint      `json:"discount,omitempty"`
	Price            uint      `json:"price"`
	Quantity         uint      `json:"quantity"`
	CategoryDiscount uint      `json:"categoryDiscount,omitempty"`
	FinalPrice       uint      `json:"payableAmount,omitempty"`
	PayableAmount    uint      `json:"PayableAmount"`
	OrderDate        time.Time `json:"orderDate"`
	DeliveryDate     string    `json:"delivaryDate,omitempty"`
	OrderStatus      string    `json:"OrderStatus,omitempty"`
	PaymentStatus    string    `json:"paymentStatus,omitempty"`
}

type Order struct {
	ID             string `json:"orderID"`
	UserID         string `gorm:"column:user_id" json:"userid"`
	Address        string `gorm:"column:address_id" json:"address_id"`
	Payment        string `gorm:"column:payment_method" json:"payment"`
	TotalPrice     uint   `json:"payable_amount"`
	Wallet         uint   `json:"walletBalance,omitempty"`
	OrderIDRazopay string `json:"razopayOrderID,omitempty"`
	Coupon         string `json:"coupon,omitempty"`
	Orders         []OrderProducts
}

type SingleOrder struct {
	Model_name    string `json:"model_name" validate:"required,min=3,max=100"`
	SingleOrderID string `json:"singleorderid" gorm:"column:item_id"`
	ID            string `gorm:"column:order_id" json:"orderID" validate:"required,number"`
	SingleUnit    uint   `json:"PriceOfAUnit" gorm:"column:saleprice"`
	Price         uint   `json:"totalAmout" `
	Quantity      uint   `json:"quantity"`
	OrderStatus   string `json:"orderStatus"`
	Coupon        string `json:"coupon,omitempty"`
	OrderDate     string `json:"orderdate"`
	EndDate       string `json:"delivaryDate,omitempty"`
	ImageURL      string `json:"imageURL"`
	FirstName     string `json:"firstName" validate:"required"`
	LastName      string `json:"lastName,omitempty"`
	Street        string `json:"street" validate:"required,alpha"`
	City          string `json:"city" validate:"required,alpha"`
	State         string `json:"state" validate:"required,alpha"`
	Pincode       string `json:"pincode" validate:"min=6"`
	LandMark      string `json:"landmark" validate:"required"`
	PhoneNumber   string `json:"phoneNumber" validate:"required,len=10,number"`
}

type SalesReport struct {
	Orders   uint `json:"total -orders"`
	Quantity uint `json:"total-unit-saled"`
	Price    uint `json:"total-revenue"`
}

type DashBord struct {
	SellerID           string `json:"sellerID"`
	TotalOrders        uint   `json:"totalOrders"`
	DeliveredOrders    uint   `json:"deliveredOrders"`
	OngoingOrders      uint   `json:"OngoingOrders"`
	CancelledOrders    uint   `json:"cancelledOrders"`
	TotalRevenue       uint   `json:"totalRevenue"`
	TotalSelledProduct uint   `json:"totalSelledProduct"`
	AdminCredit        uint   `json:"adminCredit"`
	LowStockProductID  []uint `json:"LowStockProductID"`
}

type OnlinePayment struct {
	OrderID     string `json:"orderID" gorm:"column:order_id_razopay" `
	User        string `gorm:"column:first_name" json:"user"`
	FinalPrice  uint   `json:"finalPrice"`
	PhoneNumber uint   `json:"phoneNumber" `
}

type Invoice struct {
	AddressID     string
	UserID        string
	PaymentMethod string
	ProductID     string
	SellerID      string
	Quantity      uint
	PayableAmount uint
	PaymentStatus string
	OrderDate     time.Time
}

type XlSalesReport struct {
	ItemID        string
	ProductID     string
	Model_name    string
	Quantity      uint
	PayableAmount uint
	OrderDate     time.Time
	EndDate       time.Time
}
