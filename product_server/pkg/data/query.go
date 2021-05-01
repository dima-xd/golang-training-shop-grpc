package data

const (
	selectMainOrderInfo = `orders.id, customers.name AS customer_name, customers.surname, 
		products.name AS product_name, customers.contact, date, price`
	joinCustomers            = `JOIN customers ON customers.id = customer_id`
	joinProducts             = `JOIN products ON products.id = product_id`
	selectFromProducts       = `SELECT * FROM "products"`
	selectFromProductsWithID = `SELECT * FROM "products" WHERE "id" = $1`
	insertProduct            = `INSERT INTO "products" ("name","product_category_id","quantity","unit_price","id") 
		VALUES ($1,$2,$3,$4,$5) RETURNING "id"`
)
