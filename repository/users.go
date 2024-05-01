// Package repository
// package spécifiant les fonctions utilisé pour les requêtes avec la base de donnée
package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// CreateUser reçoit en argument un user
// Crée un "users" dans la table et renvoie l'user mis à jour
func CreateUser(user models.Users) (models.Users, error) {
	err := utils.DB.Create(&user)
	return user, err.Error
}
