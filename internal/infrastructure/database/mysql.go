package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMySQL() (*gorm.DB, *sql.DB, error) {
	dsn := "root:123@tcp(192.168.32.137:3306)/order_ez?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	return db, sqlDB, err

	// 自动迁移模型
	//db.AutoMigrate(&model.Order{})
	// 其他模型迁移...
}

func CloseMySQL(sqlDB *sql.DB) {
	err := sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
