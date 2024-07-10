package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
	"time"
)

func SubscribeTypeGetByTypeAndAnnuel(typeSubscribe string, annuel bool) models.Subscribe {
	var result models.Subscribe
	utils.DB.Where("type = ? AND annuel = ?", typeSubscribe, annuel).First(&result)
	return result
}

func SubscribeGetByTravelerId(travelerId uuid.UUID) models.SubscribeTraveler {
	var subscribe models.SubscribeTraveler
	utils.DB.Where("traveler_id = ?", travelerId).Order("end_date DESC").First(&subscribe)
	return subscribe
}

func SubscribeTypeById(subscribeId uuid.UUID) models.Subscribe {
	var subscribe models.Subscribe
	utils.DB.Where("id = ?", subscribeId).First(&subscribe)
	return subscribe
}

func SubscribeCreateNewTraveler(subscribe models.SubscribeTraveler) models.SubscribeTraveler {
	utils.DB.Create(&subscribe)
	return subscribe
}

func SubscribeDeleteDateNow(travelerId uuid.UUID) {
	utils.DB.Where("traveler_id = ? AND ? BETWEEN begin_date AND end_date", travelerId, time.Now()).Delete(models.SubscribeTraveler{})
}

// SubscribeGetAll Récupère dans la bdd tous les types d'abonnement
func SubscribeGetAll() []models.Subscribe {
	var subscribe []models.Subscribe
	utils.DB.Find(&subscribe)
	return subscribe
}
