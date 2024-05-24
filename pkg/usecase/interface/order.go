package interfaceUseCase

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
)

type IOrderUseCase interface {
	NewOrder(*requestmodel.Order) (*responsemodel.Order, error)
	OrderShowcase(string) (*[]responsemodel.OrderShowcase, error)
	SingleOrder(string, string) (*responsemodel.SingleOrder, error)
	CancelUserOrder(string, string) (*responsemodel.OrderDetails, error)
	ReturnUserOrder(string, string) (*responsemodel.OrderDetails, error)

	GetSellerOrders(string, string) (*[]responsemodel.OrderDetails, error)
	ConfirmDeliverd(string, string) (*responsemodel.OrderDetails, error)
	CancelOrder(string, string) (*responsemodel.OrderDetails, error)

	GetSalesReport(string, string, string, string) (*responsemodel.SalesReport, error)
	GetSalesReportByDays(string, int) (*responsemodel.SalesReport, error)

	OrderInvoiceCreation(string) (*string, error)
	GenerateXlOfSalesReport(string) (string, error)
}
