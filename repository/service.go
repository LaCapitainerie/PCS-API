package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

func ServiceCreateOrUpdateNewService(service models.Service) models.Service {
	utils.DB.Save(&service)
	return service
}
