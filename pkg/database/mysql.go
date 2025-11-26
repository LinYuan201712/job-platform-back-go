package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"job-platform-go/internal/config"
	"log"
	"time"
)

var DB *gorm.DB

func InitDB() {
	dsnConfig := config.GlobalConfig.Datasource
	// 构建 DSN (Data Source Name)
	// 格式: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dsnConfig.Username,
		dsnConfig.Password,
		dsnConfig.Host,
		dsnConfig.Port,
		dsnConfig.Database,
		dsnConfig.Charset,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//GORM 日志，方便调试 SQL
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	//连接配置池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大复用时间
	log.Println("Database connected successfully")
}
