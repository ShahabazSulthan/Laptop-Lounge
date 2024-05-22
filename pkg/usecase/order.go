package usecase

import (
	"Laptop_Lounge/pkg/config"
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	interfaces "Laptop_Lounge/pkg/repository/interface"
	"Laptop_Lounge/pkg/service"
	interfaceUseCase "Laptop_Lounge/pkg/usecase/interface"
	"Laptop_Lounge/pkg/utils/helper"
	"errors"
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jung-kurt/gofpdf"
)

type orderUseCase struct {
	repo             interfaces.IOrderRepository
	cartrepo         interfaces.ICartRepository
	sellerRepository interfaces.ISellerRepo
	paymentRepo      interfaces.IPaymentRepository
	couponrepo       interfaces.ICouponRepository
	razopay          *config.Razopay
}

func NewOrderUseCase(repository interfaces.IOrderRepository, cartrepository interfaces.ICartRepository, sellerRepository interfaces.ISellerRepo, paymentRepository interfaces.IPaymentRepository, coupon interfaces.ICouponRepository, razopay *config.Razopay) interfaceUseCase.IOrderUseCase {
	return &orderUseCase{repo: repository, cartrepo: cartrepository, sellerRepository: sellerRepository, paymentRepo: paymentRepository, couponrepo: coupon, razopay: razopay}
}

//---------------------------------Create a new order-----------------------------------//

func (r *orderUseCase) NewOrder(order *requestmodel.Order) (resp *responsemodel.Order, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("error creating order: %v", err)
		}
	}()

	var couponData *responsemodel.Coupon
	// check payment method
	switch order.Payment {
	case "COD":
		order.OrderStatus = "processing"
		order.PaymentStatus = "pending"
	case "ONLINE":
		order.OrderStatus = "pending"
		order.PaymentStatus = "pending"
	case "WALLET":
		order.OrderStatus = "processing"
		order.PaymentStatus = "success"
	}
	// check Address
	if err = r.repo.GetAddressExist(order.UserID, order.AddressID); err != nil {
		return nil, err
	}
	// check usercart
	userCart, err := r.cartrepo.GetCart(order.UserID)
	if err != nil {
		return nil, err
	}
	order.Cart = *userCart

	for _, data := range order.Cart {
		unit, err := r.repo.GetProductUnits(data.ProductID)
		if err != nil {
			return nil, err
		}
		if *unit < data.Quantity {
			return nil, fmt.Errorf("insufficient stock for product ID %s (Available: %d, Requested: %d)", data.ProductID, *unit, data.Quantity)
		}
		newUnit := *unit - data.Quantity
		if err := r.repo.UpdateProductUnits(data.ProductID, newUnit); err != nil {
			return nil, err
		}
	}

	for _, product := range order.Cart {
		ProductPrice, err := r.cartrepo.GetProductPrice(product.ProductID)
		if err != nil {
			return nil, err
		}
		order.FinalPrice += ProductPrice
	}
	// check coupon
	if order.Coupon != "" {
		couponData, err = r.couponrepo.CheckCouponExpired(order.Coupon)
		if err != nil {
			return nil, err
		}
		if order.FinalPrice < couponData.MinimumRequired || order.FinalPrice >= couponData.MaximumAllowed {
			return nil, fmt.Errorf("order price does not meet coupon requirements (Total Price: %d, Coupon: %s, Maximum Allowed: %d)", order.FinalPrice, order.Coupon, couponData.MaximumAllowed)
		}
		if couponData.EndDate.Before(time.Now()) {
			return nil, errors.New("coupon has expired")
		}
		if exist := r.repo.CheckCouponAppliedOrNot(order.UserID, order.Coupon); exist > 0 {
			return nil, fmt.Errorf("coupon %s already applied %d times", order.Coupon, exist)
		}
		order.CouponDiscount = couponData.Discount
	}

	order.FinalPrice = 0
	for i, product := range order.Cart {
		order.Cart[i].Price = helper.FindDiscount(float64(product.Price), float64(product.CategoryDiscount+product.Discount)) * product.Quantity
		order.Cart[i].Discount = product.Discount + product.CategoryDiscount
		order.Cart[i].FinalPrice = helper.FindDiscount(float64(product.Price), float64(product.Discount+product.CategoryDiscount+order.CouponDiscount)) * product.Quantity
		order.FinalPrice += order.Cart[i].FinalPrice
	}

	switch order.Payment {
	case "ONLINE":
		orderID, err := service.Razopay(order.FinalPrice, r.razopay.RazopayKey, r.razopay.RazopaySecret)
		if err != nil {
			return nil, err
		}
		order.OrderIDRazopay = orderID
	case "WALLET":
		userWallet, err := r.paymentRepo.GetWallet(order.UserID)
		if err != nil {
			return nil, err
		}
		if userWallet.Balance < order.FinalPrice {
			return nil, fmt.Errorf("insufficient balance in the wallet (Available: %d, Required: %d)", userWallet.Balance, order.FinalPrice)
		}
		if err := r.paymentRepo.UpdateWalletReduceBalance(order.UserID, order.FinalPrice); err != nil {
			return nil, err
		}
		var walletTransactions requestmodel.WalletTransaction
		walletTransactions.UserID = order.UserID
		walletTransactions.Debit = order.FinalPrice
		walletTransactions.TotalAmount = userWallet.Balance - order.FinalPrice
		if err := r.paymentRepo.WalletTransaction(walletTransactions); err != nil {
			return nil, err
		}
	}

	orderResponse, err := r.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	OrderSuccessDetails, err := r.repo.AddProdutToOrderProductTable(order, orderResponse)
	if err != nil {
		return nil, err
	}

	for _, data := range order.Cart {
		if err := r.cartrepo.DeleteProductFromCart(data.ProductID, order.UserID); err != nil {
			return nil, err
		}
	}

	orderResponse.TotalPrice = order.FinalPrice
	return OrderSuccessDetails, nil
}

//---------------------------------Get All Products-----------------------------------//

func (r *orderUseCase) OrderShowcase(userID string) (*[]responsemodel.OrderShowcase, error) {
	abstractOrder, err := r.repo.GetOrderShowcase(userID)
	if err != nil {
		return nil, err
	}
	return abstractOrder, nil
}

//---------------------------------Get Single order using order_id-----------------------------------//

func (r *orderUseCase) SingleOrder(orderID string, userID string) (*responsemodel.SingleOrder, error) {
	singleOrder, err := r.repo.GetSingleOrder(orderID, userID)
	if err != nil {
		return nil, err
	}
	return singleOrder, nil
}

//---------------------------------Cancel User Order-----------------------------------//

func (r *orderUseCase) CancelUserOrder(orderItemID string, userID string) (*responsemodel.OrderDetails, error) {

	err := r.repo.GetOrderExistOfUser(orderItemID, userID)
	if err != nil {
		return nil, err
	}

	orderDetails, err := r.repo.UpdateUserOrderCancel(orderItemID, userID)
	if err != nil {
		return nil, err
	}

	paymentType, err := r.repo.GetPaymentType(orderItemID)
	if err != nil {
		return nil, err
	}

	if paymentType == "ONLINE" || paymentType == "WALLET" {

		orderDetails.WalletBalance, err = r.paymentRepo.CreateOrUpdateWallet(userID, orderDetails.Saleprice)
		if err != nil {
			return nil, err
		}

		currentBalance, err := r.paymentRepo.GetWalletbalance(userID)
		if err != nil {
			return nil, err
		}

		var walletTransactions requestmodel.WalletTransaction

		walletTransactions.UserID = userID
		walletTransactions.Credit = orderDetails.Saleprice
		walletTransactions.TotalAmount = *currentBalance
		fmt.Println("****", walletTransactions, *currentBalance, orderDetails.Saleprice)
		err = r.paymentRepo.WalletTransaction(walletTransactions)
		if err != nil {
			return nil, err
		}
	}

	units, err := r.repo.GetProductUnits(orderDetails.ProductID)
	if err != nil {
		return nil, err
	}

	updatedUnit := *units + orderDetails.Quantity

	err = r.repo.UpdateProductUnits(orderDetails.ProductID, updatedUnit)
	if err != nil {
		return nil, err
	}
	return orderDetails, nil
}

//---------------------------------Return User Order-----------------------------------//

func (r *orderUseCase) ReturnUserOrder(orderItemID, userID string) (*responsemodel.OrderDetails, error) {

	orderDetails, err := r.repo.UpdateUserOrderReturn(orderItemID, userID)
	if err != nil {
		return nil, err
	}

	orderDetails.WalletBalance, err = r.paymentRepo.CreateOrUpdateWallet(userID, orderDetails.Saleprice)
	if err != nil {
		return nil, err
	}

	currentBalance, err := r.paymentRepo.GetWalletbalance(userID)
	if err != nil {
		return nil, err
	}

	var walletTransactions requestmodel.WalletTransaction

	walletTransactions.UserID = userID
	walletTransactions.Credit = orderDetails.Saleprice
	walletTransactions.TotalAmount = *currentBalance

	err = r.paymentRepo.WalletTransaction(walletTransactions)
	if err != nil {
		return nil, err
	}

	units, err := r.repo.GetProductUnits(orderDetails.ProductID)
	if err != nil {
		return nil, err
	}

	updatedUnit := *units + orderDetails.Quantity

	err = r.repo.UpdateProductUnits(orderDetails.ProductID, updatedUnit)
	if err != nil {
		return nil, err
	}

	sellerCredit, err := r.sellerRepository.GetSellerCredit(orderDetails.SellerID)
	if err != nil {
		return nil, err
	}

	err = r.sellerRepository.UpdateSellerCredit(orderDetails.SellerID, sellerCredit-orderDetails.Price)
	if err != nil {
		return nil, err
	}

	return orderDetails, nil
}

// ------------------------------------------Seller Control Orders------------------------------------\\

//---------------------------------Get All seller Orders-----------------------------------//

func (r *orderUseCase) GetSellerOrders(sellerID string, remainingQuery string) (*[]responsemodel.OrderDetails, error) {
	userOrders, err := r.repo.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		return nil, err
	}
	return userOrders, nil
}

//---------------------------------Seller Conform Delivered-----------------------------------//

func (r *orderUseCase) ConfirmDeliverd(sellerID string, orderItemID string) (*responsemodel.OrderDetails, error) {

	fmt.Println("ooooiiiooo")
	err := r.repo.UpdateDeliveryTime(sellerID, orderItemID)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("hiii")

	orderDetails, err := r.repo.UpdateOrderDelivered(sellerID, orderItemID)
	if err != nil {
		return nil, err
	}

	err = r.repo.UpdateOrderPaymetSuccess(sellerID, orderItemID)
	if err != nil {
		return nil, err
	}

	sellerCredit, err := r.sellerRepository.GetSellerCredit(sellerID)
	if err != nil {
		return nil, err
	}

	err = r.sellerRepository.UpdateSellerCredit(sellerID, sellerCredit+orderDetails.Price)
	if err != nil {
		return nil, err
	}

	return orderDetails, nil
}

//---------------------------------Seller Cancel Order-----------------------------------//

func (r *orderUseCase) CancelOrder(orderID string, sellerID string) (*responsemodel.OrderDetails, error) {
	err := r.repo.GetOrderExistOfSeller(orderID, sellerID)
	if err != nil {
		return nil, err
	}
	orderDetails, err := r.repo.UpdateOrderCancel(orderID, sellerID)
	if err != nil {
		return nil, err
	}

	units, err := r.repo.GetProductUnits(orderDetails.ProductID)
	if err != nil {
		return nil, err
	}

	updatedUnit := *units + orderDetails.Quantity

	err = r.repo.UpdateProductUnits(orderDetails.ProductID, updatedUnit)
	if err != nil {
		return nil, err
	}

	return orderDetails, nil
}

// ------------------------------------------Seller Sales Report------------------------------------\\

//---------------------------------Get Sales Report Year-Month-Day-----------------------------------//

func (r *orderUseCase) GetSalesReport(sellerID, year, month, days string) (*responsemodel.SalesReport, error) {
	report, err := r.repo.GetSalesReport(sellerID, year, month, days)
	if err != nil {
		return nil, err
	}
	return report, nil
}

//---------------------------------Get Sales Reports by days-----------------------------------//

func (r *orderUseCase) GetSalesReportByDays(sellerID string, days string) (*responsemodel.SalesReport, error) {
	report, err := r.repo.GetSalesReportByDays(sellerID, days)
	if err != nil {
		return nil, err
	}
	return report, nil
}

// ------------------------------------------Order Invoice Pdf------------------------------------\\

func (r *orderUseCase) OrderInvoiceCreation(orderItemID string) (*string, error) {
	// Get order details
	orderDetails, err := r.repo.GetOrderFullDetails(orderItemID)
	if err != nil {
		return nil, err
	}

	// Get seller details
	sellerDetails, err := r.sellerRepository.GetSingleSeller(orderDetails.SellerID)
	if err != nil {
		return nil, err
	}

	// Get user address details
	userAddresses, err := r.repo.GetAddressForInvoice(orderDetails.AddressID)
	if err != nil {
		return nil, err
	}

	// Get product details
	product, err := r.repo.GetAInventoryForInvoice(orderDetails.ProductID)
	if err != nil {
		return nil, err
	}

	// Create PDF
	marginX := 10.0
	marginY := 20.0
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(marginX, marginY, marginX)
	pdf.AddPage()

	// Set font and color for header
	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(51, 102, 153) // Dark blue color
	pdf.Cell(0, 10, "Tax Invoice")
	pdf.Ln(10)

	// Set font for the company name
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(0, 10, "Laptop Lounge")
	pdf.Ln(10)

	// Order details
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(0, 0, 0) // Black color
	pdf.Cell(0, 10, "Order ID: "+orderItemID)
	pdf.Ln(5)
	pdf.Cell(0, 10, "Order Date: "+orderDetails.OrderDate.Format("2006-01-02 15:04:05"))
	pdf.Ln(5)
	pdf.Cell(0, 10, "Payment Status: "+orderDetails.PaymentStatus)
	pdf.Ln(10)

	pdf.Cell(0, 10, "Seller: "+sellerDetails.Name)
	pdf.Ln(10)

	// Address
	pdf.Cell(0, 10, "Address")
	pdf.Ln(5)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Name: %s %s", userAddresses.FirstName, userAddresses.LastName))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Address: %s, %s, %s - %s", userAddresses.Street, userAddresses.City, userAddresses.State, userAddresses.Pincode))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Landmark: %s", userAddresses.LandMark))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Phone Number: %s", userAddresses.PhoneNumber))
	pdf.Ln(15)

	// Table headers
	pdf.SetFont("Arial", "B", 12)
	pdf.SetFillColor(230, 230, 230) // Light gray background for headers
	header := [5]string{"No", "Product", "Quantity", "MRP", "Final Price"}
	colWidth := [5]float64{10.0, 75.0, 25.0, 40.0, 40.0}

	for colJ := 0; colJ < 5; colJ++ {
		pdf.CellFormat(colWidth[colJ], 7, header[colJ], "1", 0, "CM", true, 0, "")
	}
	pdf.Ln(7)

	// Table data
	pdf.SetFont("Arial", "B", 12)
	data := [5]string{"1", product.ModelName, fmt.Sprintf("%d", orderDetails.Quantity), fmt.Sprintf("%d", product.Mrp), fmt.Sprintf("%d", orderDetails.PayableAmount)}
	for colJ := 0; colJ < 5; colJ++ {
		pdf.CellFormat(colWidth[colJ], 7, data[colJ], "1", 0, "CM", false, 0, "")
	}
	pdf.Ln(7)

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(colWidth[3], 7, "Grand Total", "1", 0, "CM", true, 0, "")
	pdf.CellFormat(colWidth[4], 7, fmt.Sprintf("%d", orderDetails.PayableAmount), "1", 0, "CM", true, 0, "")
	pdf.Ln(15)

	// Footer
	pdf.SetFont("Arial", "I", 10)
	pdf.SetTextColor(100, 100, 100) // Gray color for footer
	pdf.Cell(0, 10, "Laptop_Lounge: Thanks for shopping!")

	// Save PDF to a local folder
	filePath := "C:\\Users\\shaha\\OneDrive\\Desktop\\GO-Workplace\\First Project\\Laptop_Lounge\\Report\\invoice.pdf"
	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &filePath, nil
}

// ------------------------------------------Sales Report in Excel------------------------------------\\

func (r *orderUseCase) GenerateXlOfSalesReport(sellerID string) (string, error) {
	orders, err := r.repo.GetOrderXlSalesReport(sellerID)
	if err != nil {
		return "", err
	}
	if orders == nil || len(*orders) == 0 {
		return "", errors.New("seller has no sales for creating a sales report")
	}

	f := excelize.NewFile()
	sheetName := "SalesReport"
	f.NewSheet(sheetName)

	// Set the company name heading
	f.MergeCell(sheetName, "A1", "G1")
	f.SetCellValue(sheetName, "A1", "Laptop Lounge Sales Report")
	style, _ := f.NewStyle(`{"font":{"bold":true,"size":16},"alignment":{"horizontal":"center"}}`)
	f.SetCellStyle(sheetName, "A1", "G1", style)

	// Set column headers with a color theme
	headers := []string{"ItemID", "ProductID", "Productname", "Quantity", "PayedAmount", "OrderDate", "EndDate"}
	headerStyle, _ := f.NewStyle(`{"font":{"bold":true,"color":"#FFFFFF"},"fill":{"type":"pattern","color":["#4F81BD"],"pattern":1}}`)
	for colIndex, header := range headers {
		cell := excelize.ToAlphaString(colIndex) + "2"
		f.SetCellValue(sheetName, cell, header)
		f.SetCellStyle(sheetName, cell, cell, headerStyle)
	}

	// Populate the sheet with data
	for rowIndex, record := range *orders {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex+3), record.ItemID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex+3), record.ProductID)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex+3), record.Model_name)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex+3), record.Quantity)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", rowIndex+3), record.PayableAmount)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", rowIndex+3), record.OrderDate.Format("2006-01-02 15:04:05"))
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", rowIndex+3), record.EndDate.Format("2006-01-02 15:04:05"))
	}

	// Auto-size columns for better readability
	for colIndex := 0; colIndex < len(headers); colIndex++ {
		col := excelize.ToAlphaString(colIndex)
		f.SetColWidth(sheetName, col, col, 15)
	}

	// Save the Excel file locally
	filePath := "C:\\Users\\shaha\\OneDrive\\Desktop\\GO-Workplace\\First Project\\Laptop_Lounge\\Report\\salesReport.xlsx"
	err = f.SaveAs(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return filePath, nil
}
