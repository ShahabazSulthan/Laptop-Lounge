package responsemodel

type ProductRes struct {
	ID                uint    `json:"id"`
	ModelName         string  `json:"modelName" validate:"required,min=3,max=100"`
	Description       string  `json:"description" validate:"required,min=5"`
	BrandID           uint    `json:"brandID" validate:"required"`
	CategoryID        uint    `json:"categoryID" validate:"required"`
	SellerID          string  `json:"sellerID" validate:"required"`
	Mrp               uint    `json:"mrp" validate:"required,min=0"`
	Discount          uint    `json:"discount" validate:"required,min=0,max=99"`
	SalePrice         uint    `json:"salePrice" validate:"required,min=0"`
	CategoryDiscount  uint    `json:"categoryDiscount,omitempty"`
	NetDiscount       uint    `json:"netDiscount,omitempty"`
	FinalPrice        uint    `json:"finalPrice,omitempty"`
	Units             uint64  `json:"units" validate:"required,min=0"`
	OperatingSystem   string  `json:"operatingSystem"`
	ProcessorType     string  `json:"processorType" validate:"required"`
	ScreenSize        float64 `json:"screenSize" validate:"required,min=0"`
	GraphicsCard      string  `json:"graphicsCard"`
	StorageCapacityGB uint    `json:"storageCapacityGB" validate:"required,min=0"`
	BatteryCapacity   uint    `json:"batteryCapacity" validate:"required,min=0"`
	ImageURL          string  `json:"imageURL" validate:"required"`
}

type Error struct {
	Err string
}

type ProductShowcase struct {
	ID                              uint   `json:"id"`
	ModelName                       string `json:"model_name"`
	Mrp                             int    `json:"mrp" `
	Discount                        uint   `form:"discount" `
	Saleprice                       int    `json:"sale_price" `
	CategoryDiscount                uint   `json:"categoryDiscount,omitempty"`
	NetDiscount                     uint   `json:"netDiscount,omitempty"`
	PriceAfterApplyCategoryDiscount uint   `json:"priceApplyCategoryDiscount,omitempty"`
	SellerID                        string `json:"sellerID" `
	Units                           uint   `json:"units"`
}

type FilterProduct struct {
	ID                              uint   `json:"id" gorm:"column:productid"`
	ModelName                       string `json:"model_name"`
	Mrp                             int    `json:"mrp" `
	Discount                        uint   `form:"discount" validate:"required,min=0,max=99,number"`
	Saleprice                       int    `json:"sale_price"`
	CategoryDiscount                uint   `json:"categoryDiscount,omitempty"`
	NetDiscount                     uint   `json:"netDiscount,omitempty"`
	PriceAfterApplyCategoryDiscount uint   `json:"priceApplyCategoryDiscount,omitempty"`
	Title                           string `json:"categoryOfferTitle,omitempty"`
	SellerID                        string `json:"sellerID" `
	Category                        string `json:"categoty" gorm:"column:name"`
	Brand                           string `json:"brand" gorm:"column:name"`
	Units                           uint   `json:"units"`
}
