// Package repository
// package spécifiant les fonctions utilisé pour les requêtes avec la base de donnée
package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// CreateUser reçoit en argument un user
// Crée un "users" dans la table et renvoie l'user mis à jour
func CreateUser(users models.Users) (models.Users, error) {
	err := utils.DB.Create(&users)
	return users, err.Error
}

func VerifyUserEmailAndPassword(mail string, phoneNumber string) bool {
	var count int64
	utils.DB.Model(&models.Users{}).Where("mail = ? OR phone_number ?", mail, phoneNumber).Count(&count)
	return count > 0
}
