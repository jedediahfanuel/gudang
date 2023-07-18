package models

type Product struct {
	Name string `json:"name" gorm:"primaryKey"`
	Last int    `json:"last"`
}
