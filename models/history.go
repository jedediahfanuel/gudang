package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	ProductName       string
	Product           Product `gorm:"foreignKey:ProductName"`
	WarehouseLocation string
	Warehouse         Warehouse `gorm:"foreignKey:WarehouseLocation"`
	Event             int
}
