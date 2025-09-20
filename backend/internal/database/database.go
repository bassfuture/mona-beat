package database

import (
	"fmt"
	"time"

	"nft-capture-game/internal/config"
	"nft-capture-game/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 获取底层的sql.DB对象来配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func Migrate(db *gorm.DB) error {
	// 按依赖顺序迁移表
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	
	err = db.AutoMigrate(&models.Capture{})
	if err != nil {
		return err
	}
	
	err = db.AutoMigrate(&models.NFT{})
	if err != nil {
		return err
	}
	
	return nil
}