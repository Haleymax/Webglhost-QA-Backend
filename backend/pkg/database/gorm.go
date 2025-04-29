package database

import (
	"fmt"
	"github.com/Webglhost-QA-Backend/backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func InitDB(cfg *config.MySQLConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.USER,
		cfg.PASSWORD,
		cfg.HOST,
		cfg.PORT,
		cfg.DBNAME)

	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("connect to database failed: %w", err)
	}

	log.Println("Successfully connected to database")
	DB = db
	return DB, nil
}

func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
