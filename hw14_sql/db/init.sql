CREATE USER application WITH PASSWORD 'application123' SUPERUSER;
CREATE DATABASE store;
GRANT ALL PRIVILEGES ON DATABASE store TO application;
/*to new db*/
\c store;
CREATE SCHEMA AUTHORIZATION application;
CREATE TABLE IF NOT exists store.application.Users (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) unique NOT null,
   email VARCHAR(255),
   password VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS store.application.Orders (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES application.Users(id),
  order_date timestamp NOT null,
  total_amount NUMERIC(10, 2) DEFAULT 0
);
CREATE TABLE IF NOT EXISTS store.application.Products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) unique NOT null,
  price NUMERIC(10, 2) NOT null
);
CREATE TABLE IF NOT exists store.application.OrderProducts (
  order_id INT REFERENCES application.Orders(id),
  product_id INT REFERENCES application.Products(id)
);
/*create indexes*/
CREATE INDEX if not EXISTS orders_user_id_idx ON application.orders (user_id);
CREATE INDEX if not EXISTS orderproducts_order_id_idx ON application.orderproducts (order_id);
