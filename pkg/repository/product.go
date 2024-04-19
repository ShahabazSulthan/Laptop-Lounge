package repository

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.IProductRepository {
	return &productRepository{DB: db}
}

func (d *productRepository) CreateProduct(ProductReq *requestmodel.ProductReq) (*responsemodel.ProductRes, error) {
	var insertedData responsemodel.ProductRes

	query := `
	INSERT INTO products (
		model_name, 
		description, 
		brand_id, 
		category_id, 
		mrp, 
		sale_price, 
		units, 
		operating_system, 
		processor_type, 
		screen_size_in_inches, 
		graphics_card,
		storage_capacity_gb, 
		battery_capacity , 
		seller_id, 
		image_url, 
		discount,
		status
	) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?) 
	RETURNING *;
`

	// Log the generated SQL query for debugging
	fmt.Println("Generated SQL Query:", query)

	result := d.DB.Raw(query,
		ProductReq.ModelName, ProductReq.Description, ProductReq.BrandID, ProductReq.CategoryID,
		ProductReq.Mrp, ProductReq.SalePrice, ProductReq.Units,
		ProductReq.OperatingSystem, ProductReq.ProcessorType, ProductReq.ScreenSize,
		ProductReq.GraphicsCard, ProductReq.StorageCapacityGB, ProductReq.BatteryCapacity,
		ProductReq.SellerID, ProductReq.ImageURL, ProductReq.Discount, ProductReq.Status,
	).Scan(&insertedData)

	if result.Error != nil {
		return nil, errors.New("product is not added into database")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("product is not added in database, faced some error")
	}
	return &insertedData, nil

}

func (d *productRepository) BlockSingleProductBySeller(SellerId string, productID string) error {
	query := "UPDATE products SET status='block' WHERE id= $1"
	err := d.DB.Exec(query, productID).Error
	if err != nil {
		return errors.New("can't change the status of product")
	}
	return nil
}

func (d *productRepository) UNBlockSingleProductBySeller(SellerId string, productID string) error {
	query := "UPDATE products SET status='active' WHERE id= $1"
	err := d.DB.Exec(query, productID).Error
	if err != nil {
		return errors.New("can't change the status of product in inverntories")
	}
	return nil
}

func (d *productRepository) DeleteProductBySeller(SellerID string, productID string) error {
	query := "UPDATE products SET status='delete' WHERE id= $1"
	result := d.DB.Exec(query, productID)
	if result.Error != nil {
		return errors.New("can't change the status of product in inverntories")
	}
	if result.RowsAffected == 0 {
		return errors.New("no inventory exist in table for deletion")
	}
	return nil
}

func (d *productRepository) GetAllProduct(offset int, limit int) (*[]responsemodel.ProductShowcase, error) {
	var Product []responsemodel.ProductShowcase

	query := "SELECT * FROM category_offers RIGHT JOIN products ON category_offers.category_id= products.category_id AND products.seller_id=category_offers.seller_id AND category_offers.status='active' AND category_offers.end_date>=now() WHERE products.status='Active' ORDER BY products.id OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offset, limit).Scan(&Product).Error
	if err != nil {
		return nil, errors.New("can't get products data from db")
	}
	return &Product, nil
}

func (d *productRepository) GetAProducts(id string) (*responsemodel.ProductRes, error) {
	var Products responsemodel.ProductRes

	query := "SELECT * FROM category_offers RIGHT JOIN products ON category_offers.category_id= products.category_id AND products.seller_id=category_offers.seller_id AND category_offers.status='active' AND category_offers.end_date>=now() WHERE products.id=? "
	result := d.DB.Raw(query, id).Scan(&Products)
	if result.Error != nil {
		return nil, errors.New("can't get products data from db or products is not active state")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &Products, nil
}


func (d *productRepository) GetSellerProduct(offSet int, limit int, sellerID string) (*[]responsemodel.ProductShowcase, error) {
	var inventory []responsemodel.ProductShowcase

	query := "SELECT * FROM category_offers RIGHT JOIN products ON category_offers.category_id= products.category_id AND products.seller_id=category_offers.seller_id AND category_offers.status='active' AND category_offers.end_date>=now() WHERE products.seller_id= ? ORDER BY products.id OFFSET ? LIMIT ?"

	err := d.DB.Raw(query, sellerID, offSet, limit).Scan(&inventory).Error
	if err != nil {
		return nil, errors.New("can't get inventory data from db")
	}

	return &inventory, nil
}

func (d *productRepository) UpdateProduct(product *requestmodel.EditProduct) (*responsemodel.ProductRes, error) {
	var updatedData responsemodel.ProductRes

	query := "UPDATE products SET mrp=?, discount= ?, saleprice= ?, units= ? WHERE id=? RETURNING*;"
	result := d.DB.Raw(query, product.Mrp, product.Discount, product.Saleprice, product.Units, product.ID).Scan(&updatedData)
	if result.Error != nil {
		return nil, errors.New("inventory is not updated into database")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("inventory is not updated in database , face some error")
	}
	return &updatedData, nil
}

func (d *productRepository) GetProductFilter(criterion *requestmodel.FilterCriterion) (*[]responsemodel.FilterProduct, error) {
	fmt.Println("##", criterion.MinPrice)
	var sortedProduct []responsemodel.FilterProduct

	query := `SELECT products.id AS productID, * FROM products
	 INNER JOIN categories ON categories.id= products.category_id 
	 INNER JOIN brands ON brands.id= products.brand_id 
	 LEFT JOIN category_offers ON category_offers.category_id=products.category_id 
	 AND products.seller_id=category_offers.seller_id AND category_offers.status='active' 
	 AND category_offers.end_date>=now()  WHERE categories.name ILIKE '%' || $1 || '%' 
	 AND brands.name ILIKE '%' || $2 || '%' AND products.model_name ILIKE '%' || $3 || '%' 
	 AND ($4 = 0 OR $4 < products.sale_price AND ($5 = 0 OR $5 >= products.sale_price))`
	 
	result := d.DB.Raw(query, criterion.Category, criterion.Brand, criterion.Product, criterion.MinPrice, criterion.MaxPrice).Scan(&sortedProduct)
	if result.Error != nil {
		return nil, errors.New("face some issue while filter product")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	fmt.Println("**", sortedProduct)
	return &sortedProduct, nil
}
