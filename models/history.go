package models

import (
	"gudang/models/dto"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	ProductName       string
	Product           Product `gorm:"foreignKey:ProductName"`
	WarehouseLocation string
	Warehouse         Warehouse `gorm:"foreignKey:WarehouseLocation"`
	Event             int
}

func (h *History) MapToDTO() dto.HistoryDTO {
	return dto.HistoryDTO{
		ProductName:       h.ProductName,
		WarehouseLocation: h.WarehouseLocation,
		Event:             h.Event,
	}
}
