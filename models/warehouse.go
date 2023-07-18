package models

type Warehouse struct {
	Location string `json:"location" gorm:"primaryKey"`
}
