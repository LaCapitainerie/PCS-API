package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

func ReservationServiceListCreate(service models.ReservationService) (models.ReservationService, error) {
	err := utils.DB.Create(&service).Error
	return service, err
}
