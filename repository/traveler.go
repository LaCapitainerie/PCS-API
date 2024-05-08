package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// GetAllTraveler
// Renvoie la liste de tous les "Traveler"
func GetAllTraveler() []models.Traveler {
	var Travelers []models.Traveler
	if err := utils.DB.Find(&Travelers); err.Error != nil {
		panic("Unable to get Travelers " + err.Error.Error())
	}
	return Travelers
}

// CreateTraveler reçoit en argument un traveler
// Crée un "traveler" dans la table et renvoie le voyageur mis à jour
func CreateTraveler(traveler models.Traveler) (models.Traveler, error) {
	err := utils.DB.Create(&traveler)
	return traveler, err.Error
}
