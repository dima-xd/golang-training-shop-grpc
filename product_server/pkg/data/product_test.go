package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	ReadAll() ([]Product, error)
}

var testProduct = Product{
	ID:                1,
	Name:              "test",
	ProductCategoryID: 3,
	Quantity:          4,
	UnitPrice:         "5",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

func NewGormDB(db *sql.DB) *gorm.DB {
	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       db,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gdb
}

func TestProductData_ReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gdb := NewGormDB(db)
	productData := NewProductData(gdb)
	rows := sqlmock.NewRows([]string{"id", "name", "product_category_id", "quantity", "unit_price"}).
		AddRow(testProduct.ID, testProduct.Name, testProduct.ProductCategoryID, testProduct.Quantity, testProduct.UnitPrice)
	mock.ExpectQuery(selectFromProducts).WillReturnRows(rows)
	products, err := productData.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(products)
	assert.Equal(products[0], testProduct)
	assert.Len(products, 1)
}

func TestProductData_Read(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gdb := NewGormDB(db)
	productData := NewProductData(gdb)
	rows := sqlmock.NewRows([]string{"id", "name", "product_category_id", "quantity", "unit_price"}).
		AddRow(testProduct.ID, testProduct.Name, testProduct.ProductCategoryID, testProduct.Quantity, testProduct.UnitPrice)
	mock.ExpectQuery(selectFromProductsWithID).WithArgs(1).WillReturnRows(rows)
	product, err := productData.Read(1)
	assert.NoError(err)
	assert.Equal(product, testProduct)
}

func TestProductData_Add(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gdb := NewGormDB(db)
	productData := NewProductData(gdb)
	mock.ExpectBegin()
	mock.ExpectQuery(insertProduct).
		WithArgs(testProduct.Name, testProduct.ProductCategoryID, testProduct.Quantity, testProduct.UnitPrice, testProduct.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(fmt.Sprint(1)))
	mock.ExpectCommit()
	id, err := productData.Create(testProduct)
	assert.NoError(err)
	assert.Equal(id, 1)
}

func TestProductData_ReadAll_Error(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gdb := NewGormDB(db)
	productData := NewProductData(gdb)
	mock.ExpectQuery(selectFromProducts).WillReturnError(errors.New("something went wrong..."))
	products, err := productData.ReadAll()
	assert.Error(err)
	assert.Empty(products)
}

func TestProductData_Read_Error(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gdb := NewGormDB(db)
	productData := NewProductData(gdb)
	mock.ExpectQuery(selectFromProductsWithID).WithArgs(1).WillReturnError(errors.New("something went wrong..."))
	product, err := productData.Read(1)
	assert.Error(err)
	assert.NotEqual(product, testProduct)
}

func TestProductData_Add_Error(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gdb := NewGormDB(db)
	productData := NewProductData(gdb)
	mock.ExpectBegin()
	mock.ExpectQuery(insertProduct).
		WithArgs(testProduct.Name, testProduct.ProductCategoryID, testProduct.Quantity, testProduct.UnitPrice, testProduct.ID).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	id, err := productData.Create(testProduct)
	assert.Error(err)
	assert.Equal(id, -1)
}
