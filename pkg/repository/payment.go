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

type paymentRepo struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) interfaces.IPaymentRepository {
	return &paymentRepo{DB: db}
}

func (d *paymentRepo) CreateOrUpdateWallet(userID string, creditAmount uint) (uint, error) {

	var balance uint
	query := "INSERT INTO wallets (user_id, balance) VALUES ($1, $2) ON CONFLICT(user_id) DO UPDATE SET balance=wallets.balance + $2 RETURNING balance"
	result := d.DB.Raw(query, userID, creditAmount).Scan(&balance)
	if result.Error != nil {
		return 0, errors.New("face some issue while intract with wallet for made change in balance")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return balance, nil
}

func (d *paymentRepo) GetWalletbalance(userID string) (*uint, error) {

	var currentBalance uint
	query := "SELECT balance FROM wallets WHERE user_id= ?"
	result := d.DB.Raw(query, userID).Scan(&currentBalance)
	if result.Error != nil {
		return nil, errors.New("face some issue while fetch user wallet balance")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &currentBalance, nil
}

func (d *paymentRepo) WalletTransaction(transaction requestmodel.WalletTransaction) error {
	fmt.Println("**", transaction)
	query := "INSERT INTO wallet_transactions (user_id, credit, debit, event_date, total_amount) VALUES (?,?,?,now(),?)"
	result := d.DB.Exec(query, transaction.UserID, transaction.Credit, transaction.Debit, transaction.TotalAmount)
	if result.Error != nil {
		return errors.New("face some issue while add transaction on wallet")
	}
	return nil
}

func (d *paymentRepo) GetWalletTransaction(userID string) (*[]responsemodel.WalletTransaction, error) {
	var transaction []responsemodel.WalletTransaction
	query := "SELECT * FROM wallet_transactions WHERE user_id= ? ORDER BY transaction_id DESC"
	result := d.DB.Raw(query, userID).Scan(&transaction)
	if result.Error != nil {
		return nil, errors.New("face some issue while fetch user wallet transaction")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &transaction, nil
}

// ------------------------------------------Online Payment------------------------------------\\

func (d *paymentRepo) OnlinePayment(userID, orderID string) (*responsemodel.OnlinePayment, error) {
   
    var orderDetails responsemodel.OnlinePayment
	fmt.Println(orderDetails,"///////")
	query := "SELECT * FROM users INNER JOIN orders ON orders.user_id = users.id INNER JOIN addresses ON addresses.id = address_id INNER JOIN order_products ON order_products.order_id=orders.id WHERE orders.user_id=? AND payment_status = 'pending' AND payment_method= 'ONLINE' AND orders.id=?"
	result := d.DB.Raw(query, userID, orderID).Scan(&orderDetails)	
    if result.Error != nil {
        return nil, fmt.Errorf("error executing database query: %v", result.Error)
    }
    if result.RowsAffected == 0 {
        fmt.Println("no order matched")
        return nil, resCustomError.ErrNoRowAffected
    }
    return &orderDetails, nil
}

func (d *paymentRepo) GetFinalPriceByorderID(orderID string) (uint, error) {
	var finalPrice uint
	query := "SELECT sum(payable_amount) FROM orders INNER JOIN order_products ON order_products.order_id= orders.id WHERE orders.id= ?"
	result := d.DB.Raw(query, orderID).Scan(&finalPrice)
	if result.Error != nil {
		return 0, errors.New("face some issue while getting tatal amount of order by using order id")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return finalPrice, nil
}

func (d *paymentRepo) UpdateOnlinePaymentSucess(orderID string) (*[]responsemodel.OrderDetails, error) {

	var orders []responsemodel.OrderDetails
	query := "UPDATE order_products SET payment_status='success', order_status='processing' FROM orders WHERE order_products.order_id=orders.id AND orders.order_id_razopay = ? RETURNING*"
	result := d.DB.Raw(query, orderID).Scan(&orders)
	if result.Error != nil {
		return nil, errors.New("face some issue while update order status and payment status on verify online payment success")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &orders, nil
}

func (d *paymentRepo) GetWallet(userID string) (*responsemodel.UserWallet, error) {

	var userWallet responsemodel.UserWallet
	query := "SELECT COALESCE(balance,0),* FROM wallets WHERE user_id= ?"
	result := d.DB.Raw(query, userID).Scan(&userWallet)
	if result.Error != nil {
		return nil, errors.New("face some issue while get user wallet")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &userWallet, nil
}

func (d *paymentRepo) UpdateWalletReduceBalance(userID string, amount uint) error {
	fmt.Println("##", userID, amount)
	query := "UPDATE wallets SET balance=balance-$1 WHERE user_id = $2"
	result := d.DB.Exec(query, amount, userID)
	if result.Error != nil {
		return errors.New("face some issue while update wallet balance")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}
