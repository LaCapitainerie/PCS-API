package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// GetAllAdmin
// Renvoie la liste de tous les "Admin"
func GetAllAdmin() []models.Admin {
	var Admins []models.Admin
	if err := utils.DB.Find(&Admins); err.Error != nil {
		panic("Unable to get Admins " + err.Error.Error())
	}
	return Admins
}
