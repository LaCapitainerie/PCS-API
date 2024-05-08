package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// CreateLessor reçoit en argument un lessor
// Crée un "lessor" dans la table et renvoie lessor mis à jour
func CreateLessor(lessor models.Lessor) (models.Lessor, error) {
	err := utils.DB.Create(&lessor)
	return lessor, err.Error
}
