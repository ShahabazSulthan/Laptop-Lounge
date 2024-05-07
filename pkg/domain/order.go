package domain

import "time"

type Order struct {
	ID        uint `gorm:"primary key"`
	UserID    uint
	User      Users `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	AddressID uint
	Location  Address `gorm:"foreignkey:AddressID;association_foreignkey:ID"`
	PaymentMethod  string
	OrderIDRazopay string
	CouponCode     string
	Coupon         Coupons `gorm:"foreignkey:CouponCode;association_foreignkey:ID"`
}

type OrderProducts struct {
	ItemID        uint `gorm:"primarykey"`
	OrderID       uint
	Orderid       Order `gorm:"foreignkey:OrderID;association_foreignkey:ID"`
	ProductID     uint
	Product       Products `gorm:"foreignkey:ProductID;association_foreignkey:ID"`
	SellerID      uint     `gorm:"not null"`
	Seller        Seller   `gorm:"forgienKey:SellerID;association_foreignkey:ID"`
	Quantity      uint
	Price         uint
	Discount      uint `gorm:"default:0"`
	PayableAmount uint
	OrderDate     time.Time
	EndDate       time.Time
	PaymentStatus status
	OrderStatus   string
}
