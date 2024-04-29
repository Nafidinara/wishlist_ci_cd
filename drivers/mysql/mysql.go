package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-wishlist-api/drivers/mysql/user"
	"go-wishlist-api/drivers/mysql/wishlist"
)

type Config struct {
	DBName     string
	DBUsername string
	DBPass     string
	DBHost     string
	DBPort     string
}

func ConnectDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	MigrationUser(db)
	return db
}

func MigrationUser(db *gorm.DB) {
	db.AutoMigrate(&wishlist.Wishlist{}, &user.User{})
}
