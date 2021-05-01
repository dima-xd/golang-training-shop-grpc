package data

import (
	"fmt"

	"gorm.io/gorm"
)

// Product - table products in struct form
type Product struct {
	ID                int64  `gorm:"primaryKey"`          // primary key
	Name              string `gorm:"name"`                // product's name
	ProductCategoryID int64  `gorm:"product_category_id"` // foreign key of the product_categories table
	Quantity          int64  `gorm:"quantity"`            // number of product in stock
	UnitPrice         string `gorm:"money"`               // price of product by unit
}

// ProductData - has connection to db
type ProductData struct {
	db *gorm.DB
}

// NewProductData - creates copy of ProductData to control operations with db
func NewProductData(db *gorm.DB) *ProductData {
	return &ProductData{db: db}
}

// ReadAll - gets array of products
func (p ProductData) ReadAll() ([]Product, error) {
	var products []Product
	result := p.db.Table("products").Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read products from database, error: %w", result.Error)
	}
	return products, nil
}

// Read - gets product by the id
func (p ProductData) Read(id int64) (Product, error) {
	var product Product
	result := p.db.Where("id", id).Find(&product)
	if result.Error != nil {
		return Product{}, fmt.Errorf("can't read product by the id from database, error: %w", result.Error)
	}
	return product, nil
}

// Create - adds product to the product table
func (p ProductData) Create(product Product) (int64, error) {
	result := p.db.Create(&product)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create product from database, error: %w", result.Error)
	}
	return product.ID, nil
}

// Update - update unit price in the product table by the id
func (p ProductData) Update(id int64, unitPrice string) error {
	result := p.db.Table("products").Where("id", id).Update("unit_price", unitPrice)
	if result.Error != nil {
		return fmt.Errorf("can't update product's unit price from database, error: %w", result.Error)
	}
	return nil
}

// Delete - deletes a row by the id
func (p ProductData) Delete(id int64) error {
	result := p.db.Delete(&Product{}, id)
	if result.Error != nil {
		return fmt.Errorf("can't dalete product from database, error: %w", result.Error)
	}
	return nil
}
