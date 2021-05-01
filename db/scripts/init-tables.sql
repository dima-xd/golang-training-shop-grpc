-- customers table
CREATE TABLE customers
(
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING(30) NOT NULL,
    surname CHARACTER VARYING(30) NOT NULL,
    contact CHARACTER VARYING(30) NOT NULL
);

-- offices table
CREATE TABLE offices
(
    id SERIAL PRIMARY KEY,
    city CHARACTER VARYING(100) NOT NULL,
    phone CHARACTER VARYING(15) NOT NULL,
    address CHARACTER VARYING(50) NOT NULL
);

-- product_categories table
CREATE TABLE product_categories
(
    id SERIAL PRIMARY KEY,
    category CHARACTER VARYING(100) NOT NULL
);

-- employees table
CREATE TABLE employees
(
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING(30) NOT NULL,
    surname CHARACTER VARYING(30) NOT NULL,
    office_id INTEGER REFERENCES offices (id) NOT NULL
);

-- products table
CREATE TABLE products
(
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING(100) NOT NULL,
    product_category_id INTEGER REFERENCES product_categories (id) NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price MONEY NOT NULL
);

-- ratings table
CREATE TABLE ratings
(
    id SERIAL PRIMARY KEY,
    rating REAL NOT NULL,
    product_id INTEGER REFERENCES products (id) NOT NULL
);

-- orders table
CREATE TABLE orders
(
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers (id) NOT NULL,
    product_id INTEGER REFERENCES products (id) NOT NULL,
    employee_id INTEGER REFERENCES employees (id) NOT NULL,
    date DATE NOT NULL,
    price MONEY NOT NULL,
    delivered_status BOOLEAN NOT NULL
);