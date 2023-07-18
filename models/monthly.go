package models

import "gorm.io/gorm"

type Monthly struct {
	gorm.Model
	ProductName       string
	Product           Product `gorm:"foreignKey:ProductName"`
	WarehouseLocation string
	Warehouse         Warehouse `gorm:"foreignKey:WarehouseLocation"`
	Stock             uint
}
