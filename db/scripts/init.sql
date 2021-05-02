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

-- insert customers
INSERT INTO customers(name, surname, contact) VALUES
('Usman', 'Small', 'aardo@sbcglobal.net'),
('Brendon', 'Hines', 'jandrese@outlook.com'),
('Jill', 'Lang', 'simone@live.com'),
('Philippa', 'Merritt', 'tedrlord@yahoo.com'),
('Nevaeh', 'Drummond', 'cantu@gmail.com');

-- insert offices
INSERT INTO offices(city, phone, address) VALUES
('Ely', '07766492336', '59 Botley Road'),
('Carlisle', '07701826255', '110 Trehafod Road'),
('Chichester', '07073916002', '57 Fox Lane'),
('Chester', '07050012552', '125 Boat Lane'),
('Sheffield', '07754990182', '61 Balsham Road');

-- insert product_categories
INSERT INTO product_categories(category) VALUES
('electronics'),
('home'),
('computers'),
('furniture'),
('software');

-- insert employees
INSERT INTO employees(name, surname, office_id) VALUES
('Emelia', 'Hussain', 1),
('Firat', 'Welch', 4),
('Amal', 'Mcnally', 4),
('Zeynep', 'Kline', 2),
('Terri', 'Fields', 5);

-- insert products
INSERT INTO products(name, product_category_id, quantity, unit_price) VALUES
('Kingston DDR4-1600 8 Gb', 3, 15, 67),
('Samsung Galaxy Buds', 1, 8, 120),
('Samsung Electronics EVO Select 256GB MicroSDXC', 1, 26, 30),
('Windows 10 Pro Upgrade', 5, 40, 100),
('Kraus KAG-2MB Dishwasher Air Gap', 2, 31, 25);

-- insert ratings
INSERT INTO ratings(rating, product_id) VALUES
(8.3, 1),
(5.6, 3),
(8.1, 2),
(7.1, 4),
(6.2, 5);

-- insert orders
INSERT INTO orders(customer_id, product_id, employee_id, date, price, delivered_status) VALUES
(3, 1, 2, '2021-04-02', 75.71, true),
(5, 4, 1, '2021-04-03', 113, true),
(3, 4, 4, '2021-04-03', 113, false),
(1, 2, 3, '2021-04-04', 135.6, false),
(2, 5, 1, '2021-04-06', 28.25, false);