package domain

type Cart struct {
	UserID    uint
	User      Users `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	ProductID uint
	Product   Products `gorm:"foreignkey:ProductID;association_foreignkey:ID"`
	Quantity  uint     `gorm:"default:1"`
	Status    status   `gorm:"default:active"`
}

