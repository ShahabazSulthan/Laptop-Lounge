package requestmodel

import "mime/multipart"

type ProductReq struct {
	ModelName         string                `form:"modelName" validate:"required,min=3,max=100"`
	Description       string                `form:"description" validate:"required,min=5"`
	BrandID           uint                  `form:"brandID" validate:"required,number"`
	CategoryID        uint                  `form:"categoryID" validate:"required,number"`
	SellerID          uint                  `form:"sellerID" validate:"number"`
	Mrp               uint                  `form:"mrp" validate:"required,min=0,number"`
	Discount          uint                  `form:"discount" validate:"required,min=0,max=99,number"`
	SalePrice         uint                  `form:"salePrice"`
	Units             uint64                `form:"units" validate:"required,min=0,number"`
	OperatingSystem   string                `form:"operatingSystem" validate:"required"`
	ProcessorType     string                `form:"processorType" validate:"required"`
	ScreenSize        float64               `form:"screenSize" validate:"required,min=10"`
	GraphicsCard      string                `form:"graphicsCard" validate:"required"`
	StorageCapacityGB uint                  `form:"storageCapacityGB" validate:"required,min=128"`
	BatteryCapacity   uint                  `form:"batteryCapacity" validate:"required,min=3000"`
	Image             *multipart.FileHeader `form:"image"`
	ImageURL          string
	Status            string
}

type EditProduct struct {
	ID        string `json:"id" validate:"required"`
	Mrp       uint   `json:"mrp" validate:"required,min=0"`
	Discount  uint   `form:"discount" validate:"required,min=0,max=99,number"`
	Saleprice uint   `form:"saleprice" swaggerignore:"true"`
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
