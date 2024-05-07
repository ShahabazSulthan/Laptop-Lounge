package interfaces

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IOrderRepository interface {
	CreateOrder(*requestmodel.Order) (*responsemodel.Order, error)
	GetOrderShowcase(string) (*[]responsemodel.OrderShowcase, error)
	GetSingleOrder(string, string) (*responsemodel.SingleOrder, error)
	GetProductUnits(string) (*uint, error)
	UpdateProductUnits(string, uint) error
	GetOrderPrice(string) (uint, error)
	UpdateUserOrderCancel(string, string) (*responsemodel.OrderDetails, error)
	GetPaymentType(string) (string, error)
	UpdateDeliveryTimeByUser(string, string) error
	GetOrderExistOfUser(string, string) error
	GetAddressExist(string, string) error
	AddProdutToOrderProductTable(*requestmodel.Order, *responsemodel.Order) (*responsemodel.Order, error)
	UpdateUserOrderReturn(string, string) (*responsemodel.OrderDetails, error)
	GetOrderFullDetails(string) (*responsemodel.Invoice, error)
	GetAddressForInvoice(string) (*requestmodel.Address, error)
	GetAInventoryForInvoice(id string) (*responsemodel.ProductRes, error)
	GetOrderXlSalesReport(string) (*[]responsemodel.XlSalesReport, error)

	GetSellerOrders(string, string) (*[]responsemodel.OrderDetails, error)
	UpdateOrderDelivered(string, string) (*responsemodel.OrderDetails, error)
	UpdateDeliveryTime(string, string) error
	UpdateOrderCancel(string, string) (*responsemodel.OrderDetails, error)
	UpdateOrderPaymetSuccess(string, string) error
	GetOrderExistOfSeller(string, string) error
	 CheckCouponAppliedOrNot(string, string) uint

	GetSalesReport(string, string, string, string) (*responsemodel.SalesReport, error)
	GetSalesReportByDays(string, string) (*responsemodel.SalesReport, error)

	 GetCategoryOffers(string) uint
}
