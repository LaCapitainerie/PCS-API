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

// UsersVerifyEmail reçoit en argument un string
// Vérifie dans la base de donnée si le mail dans user existe déjà
func UsersVerifyEmail(mail string) bool {
	var count int64
	utils.DB.Model(&models.Users{}).Where("mail = ?", mail).Count(&count)
	return count > 0
}

// UsersVerifyPhone reçoit en argument un string
// Vérifie dans la base de donnée si le numéro de téléphone dans user existe déjà
func UsersVerifyPhone(phoneNumber string) bool {
	var count int64
	utils.DB.Model(&models.Users{}).Where("phone_number = ?", phoneNumber).Count(&count)
	return count > 0
}

// UsersLoginVerify reçoit en argument deux string
// Vérifie les informations de connexion, renvoie l'utilisateur en question
func UsersLoginVerify(mail string, password string) models.Users {
	var user models.Users
	utils.DB.Where("mail = ? AND password = ?", mail, password).First(&user)
	return user
}
