package database

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultMySQLDSN = "root:1234@tcp(192.168.233.136:3306)/order_ez?charset=utf8mb4&parseTime=True&loc=Local"

func mysqlDSN() string {
	if dsn, ok := os.LookupEnv("MYSQL_DSN"); ok && dsn != "" {
		return dsn
	}
	return defaultMySQLDSN
}

func InitMySQL() (*gorm.DB, *sql.DB, error) {
	db, err := gorm.Open(mysql.Open(mysqlDSN()), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to access underlying sql.DB: %w", err)
	}
	return db, sqlDB, nil
}

func CloseMySQL(sqlDB *sql.DB) error {
	if sqlDB == nil {
		return nil
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close MySQL connection: %w", err)
	}
	return nil
}
