package services

import (
	"gudang/models"
	"gudang/models/dto"
	"gudang/repositories"

	"gorm.io/gorm"
)

func GetAllHistories(db *gorm.DB) ([]dto.HistoryDTO, error) {
	histories, err := repositories.GetAllHistories(db)
	if err != nil {
		return nil, err
	}

	return MapToHistoryDTO(histories), nil
}

func MapToHistoryDTO(h []models.History) []dto.HistoryDTO {
	var dto []dto.HistoryDTO
	for _, v := range h {
		dto = append(dto, v.MapToDTO())
	}
	return dto
}
