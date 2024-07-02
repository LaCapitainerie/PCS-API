package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

func SubscribeGetByTypeAndAnnuel(typeSubscribe string, annuel bool) models.Subscribe {
	var result models.Subscribe
	utils.DB.Where("type = ? AND annuel = ?", typeSubscribe, annuel).First(&result)
	return result
}
