package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// CreateProvider reçoit en argument un provider
// Crée un "provider" dans la table et renvoie le prestataire mis à jour
func CreateProvider(traveler models.Provider) (models.Provider, error) {
	err := utils.DB.Create(&traveler)
	return traveler, err.Error
}
