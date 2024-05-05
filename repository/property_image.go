package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// GetAllProperty_image
// Renvoie la liste de tous les "Property_image"
func GetAllProperty_image() []models.Property_image {
	var Property_images []models.Property_image
	if err := utils.DB.Find(&Property_images); err.Error != nil {
		panic("Unable to get Property_images " + err.Error.Error())
	}
	return Property_images
}
