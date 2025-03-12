package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user" // 表名是 "user" 而不是默认的 "users"
}

//-- 创建数据库（如果不存在）
//CREATE DATABASE IF NOT EXISTS order_ez;
//
//-- 使用数据库
//USE order_ez;
//
//-- 创建 user 表
//CREATE TABLE user (
//id INT AUTO_INCREMENT PRIMARY KEY,
//username VARCHAR(255) NOT NULL UNIQUE,
//password VARCHAR(255) NOT NULL,
//created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
//);
