package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// GetAllPropertyImage
// Renvoie la liste de tous les "Property_image"
func GetAllPropertyImage() []models.PropertyImage {
	var PropertyImages []models.PropertyImage
	if err := utils.DB.Find(&PropertyImages); err.Error != nil {
		panic("Unable to get Property_images " + err.Error.Error())
	}
	return PropertyImages
}

func PropertyImageCreate(image models.PropertyImage) models.PropertyImage {
	utils.DB.Create(&image)
	return image
}
