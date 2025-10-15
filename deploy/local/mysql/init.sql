CREATE DATABASE IF NOT EXISTS order_ez CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE order_ez;

CREATE TABLE IF NOT EXISTS oe_user (
    user_id BIGINT PRIMARY KEY,
    username VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS oe_order (
    order_id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE,
    total_price INT NOT NULL,
    CONSTRAINT fk_oe_order_user FOREIGN KEY (user_id) REFERENCES oe_user(user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS oe_order_item (
    order_item_id BIGINT PRIMARY KEY,
    order_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    item_id BIGINT NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    item_price INT NOT NULL,
    item_count INT NOT NULL,
    CONSTRAINT fk_oe_order_item_order FOREIGN KEY (order_id) REFERENCES oe_order(order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE INDEX idx_oe_order_item_order_id ON oe_order_item(order_id);
