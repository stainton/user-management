
-- 有48个产品，不确定数量的用户，我需要完成以下的事情
-- 1. 获取指定用户在某天购买了哪些产品以及购买各个产品的花费
-- 2. 获取某天所有用户购买了哪些产品以及购买各个产品的花费

USE testdb;

CREATE TABLE customers (
    user_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'unique ID of user',
    username VARCHAR(255) NOT NULL COMMENT 'username',
    telnum VARCHAR(255) COMMENT 'telphone number',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'user rigister time'
) COMMENT 'table of customers';

CREATE TABLE products (
    product_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'unique id of product',
    product_area VARCHAR(255) NOT NULL COMMENT 'product area'
) COMMENT 'table of products';

CREATE TABLE orders (
    order_id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'unique ID of order',
    user_id INT NOT NULL COMMENT 'user ID',
    order_date DATE NOT NULL COMMENT 'order date',
    total_amount DECIMAL(10, 2) NOT NULL COMMENT 'total cost',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'order created time',
    FOREIGN KEY (user_id) REFERENCES customers(user_id) ON DELETE CASCADE
) COMMENT 'table of orders';

CREATE TABLE order_details (
    order_detail_id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'unique ID of order detail',
    order_id BIGINT NOT NULL COMMENT 'orderID',
    product_id INT NOT NULL COMMENT 'product ID',
    total_price DECIMAL(10, 2) NOT NULL COMMENT 'cost of this order',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'order created time',
    FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
) COMMENT 'table of order details';

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'unique ID of user',
    username VARCHAR(255) NOT NULL COMMENT 'username',
    password VARCHAR(255) NOT NULL COMMENT 'encrypted password',
    email VARCHAR(255) NOT NULL COMMENT 'email',
    role ENUM('admin', 'viewer') NOT NULL DEFAULT 'viewer',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'user created time'
) COMMENT 'admin user';