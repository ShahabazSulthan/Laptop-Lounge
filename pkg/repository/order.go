package repository

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.IOrderRepository {
	return &orderRepository{DB: db}
}

//-------------------------Create the Order-----------------------------------//

func (d *orderRepository) CreateOrder(order *requestmodel.Order) (*responsemodel.Order, error) {

	var orderSucess = &responsemodel.Order{}

	query := "INSERT INTO orders (user_id, address_id, payment_method, order_id_razopay) VALUES(?, ?, ?, ?) RETURNING*"
	result := d.DB.Raw(query, order.UserID, order.AddressID, order.Payment, order.OrderIDRazopay).Scan(&orderSucess)
	if result.Error != nil {
		return nil, errors.New("face some issue while creating order")
	}
	if result.RowsAffected == 0 {

		return nil, errors.New("creating order row is not affected (no data matched the specified criteria)")
	}
	return orderSucess, nil
}

//-------------------------Add Products Order_Product Table-----------------------------------//

func (d *orderRepository) AddProdutToOrderProductTable(order *requestmodel.Order, orderDetails *responsemodel.Order) (*responsemodel.Order, error) {
	var orderProduct responsemodel.OrderProducts
	today := time.Now().Format("2006-01-02 15:04:05")

	for _, data := range order.Cart {
		query := "INSERT INTO order_products (order_id, product_id, seller_id, quantity, order_date, order_status, payment_status, price, discount, payable_amount) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *"
		result := d.DB.Raw(query, orderDetails.ID, data.ProductID, data.SellerID, data.Quantity, today, order.OrderStatus, order.PaymentStatus, data.Price, data.Discount, data.FinalPrice).Scan(&orderProduct)
		if result.Error != nil {
			return nil, result.Error
		}
		orderDetails.Orders = append(orderDetails.Orders, orderProduct)
	}
	return orderDetails, nil
}

//-------------------------Check Address is Exist-----------------------------------//

func (d *orderRepository) GetAddressExist(userID, addressesID string) error {
	query := "SELECT * FROM addresses WHERE userid= ? AND id= ?"
	result := d.DB.Exec(query, userID, addressesID)
	if result.Error != nil {
		return errors.New("face some issue while chcking address is exist of user")
	}
	if result.RowsAffected == 0 {
		return errors.New("user does not have specified address (no data matched the specified criteria)")
	}
	return nil
}

//-------------------------Get the All Orders-----------------------------------//

func (d *orderRepository) GetOrderShowcase(userID string) (*[]responsemodel.OrderShowcase, error) {
	fmt.Println("***", userID)
	var OrderShowcase []responsemodel.OrderShowcase
	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id=order_products.order_id INNER JOIN products ON products.id=order_products.product_id WHERE orders.user_id=? ORDER BY order_products.item_id DESC"
	result := d.DB.Raw(query, userID).Scan(&OrderShowcase)
	if result.Error != nil {
		return nil, errors.New("face some issue while order showcase")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("show all orders row is not affected (no data matched the specified criteria)")
	}
	return &OrderShowcase, nil
}

//-------------------------Get single order by item_id-----------------------------------//

func (d *orderRepository) GetSingleOrder(orderID string, userID string) (*responsemodel.SingleOrder, error) {

	var OrderShowcase *responsemodel.SingleOrder
	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id=order_products.order_id INNER JOIN products ON products.id=order_products.product_id INNER JOIN addresses ON addresses.id= orders.address_id WHERE  order_products.item_id=? AND orders.user_id=?"
	result := d.DB.Raw(query, orderID, userID).Scan(&OrderShowcase)
	if result.Error != nil {
		return nil, errors.New("face some issue while get single order")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("get single order row is not affected (no data matched the specified criteria)")
	}
	return OrderShowcase, nil
}

//-------------------------Check the product Units-----------------------------------//

func (d *orderRepository) GetProductUnits(ProductID string) (*uint, error) {

	var units uint
	query := "SELECT units FROM products WHERE id=?"
	result := d.DB.Raw(query, ProductID).Scan(&units)
	if result.Error != nil {
		return nil, errors.New("face some issue while get product units")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("get product units order row is not affected (no data matched the specified criteria)")
	}
	return &units, nil
}

//-------------------------Create the Order-----------------------------------//

func (d *orderRepository) GetPaymentType(orderItemID string) (string, error) {

	var paymentType string
	fmt.Println("@@", orderItemID)
	query := "SELECT payment_method FROM orders INNER JOIN order_products ON orders.id=order_products.order_id WHERE order_products.item_id=?"
	result := d.DB.Raw(query, orderItemID).Scan(&paymentType)
	if result.Error != nil {
		return "", errors.New("face some issue while get payment type of order")
	}
	if result.RowsAffected == 0 {
		return "", errors.New("get payment type row is not affected (no data matched the specified criteria)")
	}
	return paymentType, nil
}

//-------------------------Create the Order-----------------------------------//

func (d *orderRepository) UpdateProductUnits(ProductID string, units uint) error {

	query := "UPDATE products SET units= ? WHERE id =?"
	result := d.DB.Exec(query, units, ProductID)
	if result.Error != nil {
		return errors.New("face some issue while updating Products unit")
	}
	if result.RowsAffected == 0 {
		return errors.New("updating Products units row is not affected (no data matched the specified criteria)")
	}
	return nil
}

//-------------------------Get Order Price-----------------------------------//

func (d *orderRepository) GetOrderPrice(orderID string) (uint, error) {
	fmt.Println("&&77", orderID)
	var price uint
	query := "SELECT price FROM orders WHERE id =?"
	result := d.DB.Raw(query, orderID).Scan(&price)
	if result.Error != nil {
		return 0, errors.New("face some issue while get credit from seller table")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("get order price row is not affected (no data matched the specified criteria)")
	}
	return price, nil
}

//-------------------------Update User Order Cancel -----------------------------------//

func (d *orderRepository) UpdateUserOrderCancel(orderItemID string, userID string) (*responsemodel.OrderDetails, error) {
	var cancelOrder responsemodel.OrderDetails
	today := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE order_products SET order_status= 'cancelled', payment_status= 'refunded', end_date=? FROM orders WHERE orders.id=order_products.order_id AND item_id=? AND user_id= ? AND order_status='processing' RETURNING*"
	result := d.DB.Raw(query, today, orderItemID, userID).Scan(&cancelOrder)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is canceling")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("update user order cancel row is not affected (no data matched the specified criteria)")
	}
	return &cancelOrder, nil
}

//-------------------------Update User Order Return -----------------------------------//

func (d *orderRepository) UpdateUserOrderReturn(orderItemID string, userID string) (*responsemodel.OrderDetails, error) {

	var returnOrder responsemodel.OrderDetails
	today := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE order_products SET order_status= 'return', payment_status= 'refunded', end_date=? FROM orders WHERE orders.id=order_products.order_id AND item_id=? AND user_id= ? AND order_status='delivered' RETURNING*"
	result := d.DB.Raw(query, today, orderItemID, userID).Scan(&returnOrder)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is return")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no deliverd order exist for the given order item id of the user (no data matched the specified criteria)")
	}
	return &returnOrder, nil
}

//-------------------------Update Delivery Time-----------------------------------//

func (d *orderRepository) UpdateDeliveryTimeByUser(userID string, orderItemID string) error {

	delivaryTime := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE orders SET delivery_date= ? WHERE user_id= ? AND id = ?"
	result := d.DB.Exec(query, delivaryTime, userID, orderItemID)
	if result.Error != nil {
		return errors.New("face some issue while updating delivary time")
	}
	if result.RowsAffected == 0 {
		return errors.New("update delivery time row is not affected (no data matched the specified criteria)")
	}
	return nil
}

//-------------------------Check Order Exists of User-----------------------------------//

func (d *orderRepository) GetOrderExistOfUser(orderItemID, userID string) error {

	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id=order_products.order_id WHERE item_id= $1 AND user_id= $2"

	result := d.DB.Exec(query, orderItemID, userID)
	if result.Error != nil {
		return errors.New("encountered an issue while checking if the order exists")
	}
	if result.RowsAffected == 0 {
		return errors.New("no orders were found matching the specified criteria (no data matched the specified criteria)")
	}
	if result.RowsAffected != 0 {
		return nil
	}
	return nil
}

// ------------------------------------------Seller Control Orders------------------------------------\\

//-------------------------Get All Orders-----------------------------------//

func (d *orderRepository) GetSellerOrders(sellerID string, remainingQuery string) (*[]responsemodel.OrderDetails, error) {

	var orderList *[]responsemodel.OrderDetails
	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id= order_products.order_id WHERE seller_id=? AND order_status" + remainingQuery
	result := d.DB.Raw(query, sellerID).Scan(&orderList)
	if result.Error != nil {
		return nil, errors.New("face some issue while get user orders")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("get all seller order row is not affected (no data matched the specified criteria)")
	}
	return orderList, nil
}

//-------------------------Update Delivery Time By Seller-----------------------------------//

func (d *orderRepository) UpdateDeliveryTime(sellerID string, orderItemID string) error {
	deliveryTime := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE order_products SET end_date= ? FROM orders WHERE orders.id= order_products.order_id AND seller_id= ? AND order_products.item_id= ? AND order_status='processing'"
	result := d.DB.Exec(query, deliveryTime, sellerID, orderItemID)
	if result.Error != nil {
		return errors.New("encountered an issue while updating delivery time")
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated (no data matched the specified criteria)")
	}
	return nil
}

//-------------------------Update Order Delivered-----------------------------------//

func (d *orderRepository) UpdateOrderDelivered(sellerID string, orderItemID string) (*responsemodel.OrderDetails, error) {
	var deliveryDetails responsemodel.OrderDetails
	query := "UPDATE order_products SET order_status='delivered',payment_status = 'success' FROM orders WHERE orders.id= order_products.order_id AND seller_id= ? AND order_products.item_id= ? RETURNING*"
	result := d.DB.Raw(query, sellerID, orderItemID).Scan(&deliveryDetails)
	if result.Error != nil {
		return nil, errors.New("face some issue while update order delivered")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows were update order delivered (no data matched the specified criteria)")
	}
	return &deliveryDetails, nil
}

//-------------------------Update Order Payment Success-----------------------------------//

func (d *orderRepository) UpdateOrderPaymetSuccess(sellerID string, orderItemID string) error {

	query := "UPDATE order_products SET payment_status = 'success' FROM orders WHERE orders.id= order_products.order_id AND seller_id= ? AND order_products.item_id= ?"

	result := d.DB.Exec(query, sellerID, orderItemID)
	if result.Error != nil {
		return errors.New("face some issue while update payment status success")
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows were update order payment success (no data matched the specified criteria)")
	}
	return nil
}

//-------------------------Update Order Cancel By Seller-----------------------------------//

func (d *orderRepository) UpdateOrderCancel(orderID string, sellerID string) (*responsemodel.OrderDetails, error) {

	var cancelOrder responsemodel.OrderDetails
	query := "UPDATE order_products SET order_status= 'cancel', payment_status='cancel' WHERE order_id=? AND seller_id= ? AND order_status='processing' RETURNING*"
	result := d.DB.Raw(query, orderID, sellerID).Scan(&cancelOrder)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is cancel")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows were update order cancel (no data matched the specified criteria)")
	}
	return &cancelOrder, nil
}

//-------------------------Check Order Exist By Using Seller_id-----------------------------------//

func (d *orderRepository) GetOrderExistOfSeller(orderID, sellerID string) error {

	query := "SELECT * FROM order_products WHERE order_id= $1 AND seller_id=$2"

	result := d.DB.Exec(query, orderID, sellerID)
	if result.Error != nil {
		return errors.New("encountered an issue while checking if the order exists")
	}
	if result.RowsAffected == 0 {
		return errors.New("no orders were found matching the specified criteria")
	}
	if result.RowsAffected != 0 {
		return nil
	}
	return nil
}

// ------------------------------------------Sales Report------------------------------------\\

//-------------------------Get Sales Report by Year-Month-Day-----------------------------------//

func (d *orderRepository) GetSalesReport(sellerID, year, month, day string) (*responsemodel.SalesReport, error) {

	var remainingQuery string

	if year != "" {
		remainingQuery = " EXTRACT(YEAR FROM order_date)=" + year
	}
	if year != "" && month != "" {
		remainingQuery = " EXTRACT(YEAR FROM order_date)=" + year + " AND EXTRACT(Month FROM order_date)=" + month
	}
	if year != "" && month != "" && day != "" {
		remainingQuery = " EXTRACT(YEAR FROM order_date)=" + year + " AND EXTRACT(Month FROM order_date)=" + month + " AND EXTRACT(Day FROM order_date)=" + day
	}

	var report responsemodel.SalesReport
	query := "SELECT COUNT(*) AS Orders, SUM(quantity) AS Quantity, SUM(price) AS Price FROM order_products WHERE seller_id= ? AND order_status='delivered' AND" + remainingQuery
	result := d.DB.Raw(query, sellerID).Scan(&report)
	if result.Error != nil {
		return nil, errors.New("face some issue while get report (no data matched the specified criteria)")
	}
	return &report, nil
}

//-------------------------Get Sales Report by Days-----------------------------------//

func (d *orderRepository) GetSalesReportByDays(sellerID string, days string) (*responsemodel.SalesReport, error) {
	var report responsemodel.SalesReport
	remainingQuery := "(now() - interval '" + days + " day')"
	query := "SELECT COUNT(*) AS Orders, SUM(quantity) AS Quantity, SUM(price) AS Price FROM order_products WHERE seller_id = ? AND order_status='delivered' AND order_date >= " + remainingQuery
	result := d.DB.Raw(query, sellerID).Scan(&report)

	if result.Error != nil {
		return nil, errors.New("face some issue while get report by days (no data matched the specified criteria)")
	}
	return &report, nil
}

//-------------------------Get Excel Sales Report-----------------------------------//

func (d *orderRepository) GetOrderXlSalesReport(sellerID string) (*[]responsemodel.XlSalesReport, error) {
	var order []responsemodel.XlSalesReport

	query := "SELECT * FROM orders INNER JOIN order_products ON order_products.order_id=orders.id INNER JOIN products ON products.id=order_products.product_id WHERE order_products.order_status='delivered' AND order_products.seller_id=? "
	result := d.DB.Raw(query, sellerID).Scan(&order)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is cancel")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("face some issue while get report by xl report (no data matched the specified criteria)")
	}
	return &order, nil
}

// ------------------------------------------category_offers------------------------------------\\

// ------------------------------------------Check Coupon Applied Or Not ------------------------------------\\

func (d *orderRepository) CheckCouponAppliedOrNot(userID, couponID string) uint {
	var exist uint
	query := "SELECT COUNT(*) FROM orders WHERE user_id=? AND coupon_code= ?"
	d.DB.Raw(query, userID, couponID).Scan(&exist)
	return exist
}

// ------------------------------------------Get Category_offers to make discount------------------------------------\\

func (d *orderRepository) GetCategoryOffers(productID string) uint {
	var categoryDiscount uint
	query := "SELECT category_discount FROM category_offers RIGHT JOIN products ON products.seller_id=category_offers.seller_id AND category_offers.category_id=products.category_id AND category_offers.status='active' AND category_offers.end_date>now() WHERE products.status='active'  AND products.id=?"
	d.DB.Raw(query, productID).Scan(&categoryDiscount)
	return categoryDiscount
}

// ------------------------------------------Get Full Order Details------------------------------------\\

func (d *orderRepository) GetOrderFullDetails(orderItemID string) (*responsemodel.Invoice, error) {
	var orderDetails responsemodel.Invoice
	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id=order_products.order_id WHERE  order_products.item_id= ?;	"
	result := d.DB.Raw(query, orderItemID).Scan(&orderDetails)
	if result.Error != nil {
		return nil, errors.New("face some issue while get order details")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("face some issue while get full order details (no data matched the specified criteria)")
	}
	return &orderDetails, nil
}

// ------------------------------------------Get Address For Invoice------------------------------------\\

func (d *orderRepository) GetAddressForInvoice(addressID string) (*requestmodel.Address, error) {

	var address *requestmodel.Address
	query := "SELECT * FROM addresses WHERE id= ?"
	result := d.DB.Raw(query, addressID).Scan(&address)
	if result.Error != nil {
		return nil, errors.New("face some issue while address fetch")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("face some issue while fetch address for invoice (no data matched the specified criteria)")
	}
	return address, nil
}

// ------------------------------------------Get Product Invoice------------------------------------\\

func (d *orderRepository) GetAInventoryForInvoice(id string) (*responsemodel.ProductRes, error) {
	var inventory responsemodel.ProductRes

	query := "SELECT * FROM category_offers RIGHT JOIN products ON category_offers.category_id= products.category_id AND products.seller_id=category_offers.seller_id AND category_offers.status='active' AND category_offers.end_date>=now() WHERE products.id=? AND products.status='active'"
	result := d.DB.Raw(query, id).Scan(&inventory)
	if result.Error != nil {
		return nil, errors.New("can't get product data from db or product is not active state")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("face some issue while get product invoice (no data matched the specified criteria)")
	}
	return &inventory, nil
}
