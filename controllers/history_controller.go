package controllers

import (
	"encoding/json"
	"net/http"

	"gudang/services"

	"gorm.io/gorm"
)

type HistoryController struct {
	DB *gorm.DB
}

func NewHistoryController(db *gorm.DB) *HistoryController {
	return &HistoryController{DB: db}
}

func (c *HistoryController) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	histories, err := services.GetAllHistories(c.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{
		StatusCode: http.StatusOK,
		Message:    "Success fetching histories",
		Data:       histories,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
