package repositories

import (
	"gudang/models"

	"gorm.io/gorm"
)

func GetAllHistories(db *gorm.DB) ([]models.History, error) {
	var histories []models.History
	res := db.Find(&histories)
	return histories, res.Error
}
