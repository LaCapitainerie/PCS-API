package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// GetAllProperty
// Renvoie la liste de tous les "Property"
func GetAllProperty() []models.Property {
	var Propertys []models.Property
	if err := utils.DB.Find(&Propertys); err.Error != nil {
		panic("Unable to get Propertys " + err.Error.Error())
	}
	return Propertys
}
