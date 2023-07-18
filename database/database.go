package database

import (
	models "gudang/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var con *gorm.DB

func connect() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetDBConnection() *gorm.DB {
	if con == nil {
		con = connect()
	}
	return con
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Warehouse{},
		&models.History{},
		&models.Monthly{})
}
