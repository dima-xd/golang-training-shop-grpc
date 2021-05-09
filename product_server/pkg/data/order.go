package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Order - table orders in struct form
type Order struct {
	ID              int       `gorm:"primaryKey"`       // primary key
	CustomerID      int       `gorm:"customer_id"`      // foreign key of the customers table
	ProductID       int       `gorm:"product_id"`       // foreign key of the products table
	EmployeeID      int       `gorm:"employee_id"`      // foreign key of the employees table
	Date            time.Time `gorm:"date"`             // order's date
	Price           string    `gorm:"price"`            // order's price
	DeliveredStatus bool      `gorm:"delivered_status"` // is delivered order
}

// ResultOrder - represents main information of orders table
type ResultOrder struct {
	ID           int       // primary key
	CustomerName string    `gorm:"column:customer_name"` // customer's name
	Surname      string    // customer's surname
	Contact      string    // customer's contact
	ProductName  string    `gorm:"column:product_name"` // product's name
	Date         time.Time // order's date
	Price        string    // order's price
}

// OrderData - has connection to db
type OrderData struct {
	db *gorm.DB
}

// NewOrderData - creates copy of OrderData to control operations with db
func NewOrderData(db *gorm.DB) *OrderData {
	return &OrderData{db: db}
}

// ReadAll - gets array of ResultOrder
func (o OrderData) ReadAll() ([]ResultOrder, error) {
	var results []ResultOrder
	result := o.db.Table("orders").
		Select(selectMainOrderInfo).
		Joins(joinCustomers).
		Joins(joinProducts).
		Scan(&results)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read orders from database, error: %w", result.Error)
	}
	return results, nil
}
