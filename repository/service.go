package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

func ServiceCreateNewService(service models.Service) (models.Service, error) {
	err := utils.DB.Create(&service).Error
	return service, err
}

func ServiceUpdate(service models.Service) models.Service {
	utils.DB.Save(&service)
	return service
}

func ServiceGetWithServiceId(id uuid.UUID) (models.Service, error) {
	var service models.Service
	utils.DB.Where("id = ?", id).First(&service)
	err := utils.DB.First(&service, id).Error
	return service, err
}
