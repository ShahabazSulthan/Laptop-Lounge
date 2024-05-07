package domain

type Products struct {
	ID                 uint `gorm:"primary key"`
	ModelName          string
	Description        string
	BrandID            uint     `gorm:"not null"`
	Brand              Brand    `gorm:"foreignkey:BrandID;association_foreignkey:ID"`
	CategoryID         uint     `gorm:"not null"`
	Category           Category `gorm:"foreignkey:CategoryID;association_foreignkey:ID"`
	SellerID           string   `gorm:"not null"`
	Seller             Seller   `gorm:"foreignkey:SellerID;association_foreignkey:ID"`
	Mrp                int
	Discount           uint
	SalePrice          int
	Units              int64
	OperatingSystem    string
	ProcessorType      string
	ScreenSizeInInches float64
	GraphicsCard       string
	StorageCapacityGB  int
	BatteryCapacity    int
	ImageURL           string
	Status             status `gorm:"default:active"`
}

// type Rating struct {
// 	ID        uint `gorm:"primary_key"`
// 	ProductID uint
// 	Rating    float64
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
