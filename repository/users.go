package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

func CreateUser(user models.Users) models.Users {
	if err := utils.DB.Create(&user).Error; err != nil {
		panic("Impossible d'insérer l'utilisateur")
	}
	return user
}
