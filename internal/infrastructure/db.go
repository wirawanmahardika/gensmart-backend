package infrastructure

import (
	"fmt"
	"gensmart/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() *gorm.DB {
	dialect := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true",
		config.AppConfig.DBUser,
		config.AppConfig.DBPass,
		config.AppConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dialect), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// db.AutoMigrate(&adminDomain.Entity{}, &schoolDomain.Entity{})

	if err != nil {
		panic(err)
	}
	return db
}
