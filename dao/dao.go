package dao

import (
	"forever/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitPgsql() {
	dsn := os.Getenv("DSN") // 打包时这里不能读取环境变量，直接硬编码
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:    true,            // 取消事务
		DefaultTransactionTimeout: 5 * time.Second, // 事务默认超时时间
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL database: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	err = db.AutoMigrate(&model.Data{})
	if err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}

	DB = db
}
