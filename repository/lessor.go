package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

// CreateLessor reçoit en argument un lessor
// Crée un "lessor" dans la table et renvoie lessor mis à jour
func CreateLessor(lessor models.Lessor) (models.Lessor, error) {
	err := utils.DB.Create(&lessor)
	return lessor, err.Error
}

func IsALessor(id uuid.UUID) bool {
	var count int64
	utils.DB.Model(&models.Lessor{}).Where("user_id = ?", id).Count(&count)
	return count > 0
}

func GetLessorIdByUserId(id uuid.UUID) uuid.UUID {
	var lessor models.Lessor
	utils.DB.Where("user_id + ?", id).Find(&lessor)
	return lessor.ID
}
