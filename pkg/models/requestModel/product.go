package requestmodel

type ProductReq struct {
	ModelName         string  `json:"modelName" validate:"required,min=3,max=100"`
	Description       string  `json:"description" validate:"required,min=5"`
	BrandID           uint    `json:"brandID" validate:"required,number"`
	CategoryID        uint    `json:"categoryID" validate:"required,number"`
	SellerID          uint    `json:"sellerID" validate:"number"`
	Mrp               uint    `json:"mrp" validate:"required,min=0,number"`
	Discount          uint    `json:"discount" validate:"required,min=0,max=99,number"`
	SalePrice         uint    `json:"salePrice"`
	Units             uint64  `json:"units" validate:"required,min=0,number"`
	OperatingSystem   string  `json:"operatingSystem" validate:"required"`
	ProcessorType     string  `json:"processorType" validate:"required"`
	ScreenSize        float64 `json:"screenSize" validate:"required,min=10"`
	GraphicsCard      string  `json:"graphicsCard" validate:"required"`
	StorageCapacityGB uint    `json:"storageCapacityGB" validate:"required,min=128"`
	BatteryCapacity   uint    `json:"batteryCapacity" validate:"required,min=3000"`
	ImageURL          string  `json:"imageURL"`
}

type EditProduct struct {
	ModelName string `json:"modelName" validate:"required,min=3,max=100"`
	ID        string `json:"id" validate:"required"`
	Mrp       uint   `json:"mrp" validate:"required,min=0"`
	Discount  uint   `json:"discount" validate:"required,min=0,max=99,number"`
	Saleprice uint   `json:"saleprice"`
	Units     uint64 `json:"units" validate:"required,min=0"`
	SellerID  string `json:"-"`
}

type FilterCriterion struct {
	Category string `json:"category" validate:"alpha"`
	Brand    string `json:"brand" validate:"alpha"`
	Product  string `json:"product" validate:"alpha"`
	MinPrice uint   `json:"minprice" validate:"numeric"`
	MaxPrice uint   `json:"maxprice" validate:"numeric"`
}
