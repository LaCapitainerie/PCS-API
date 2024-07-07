package repository

import (
	"PCS-API/models"
	"PCS-API/utils"

	"github.com/google/uuid"
)

// CreateLessor reçoit en argument un lessor
// Crée un "lessor" dans la table et renvoie lessor mis à jour
func CreateLessor(lessor models.Lessor) (models.Lessor, error) {
	err := utils.DB.Save(&lessor)
	return lessor, err.Error
}

func IsALessor(id uuid.UUID) bool {
	var count int64
	utils.DB.Model(&models.Lessor{}).Where("user_id = ?", id).Count(&count)
	return count > 0
}

func GetLessorIdByUserId(id uuid.UUID) uuid.UUID {
	var lessor models.Lessor
	utils.DB.Where("user_id = ?", id).Find(&lessor)
	return lessor.ID
}

func LessorGetByUserId(id uuid.UUID) models.Lessor {
	var lessor models.Lessor
	utils.DB.Where("user_id = ?", id).Find(&lessor)
	return lessor
}

func GetUserByLessorId(id uuid.UUID) string {
	var user string
	utils.DB.Model(&models.Lessor{}).Select("user_id").Where("id = ?", id).Scan(&user)
	return user
}

func LessorGetById(id uuid.UUID) models.Lessor {
	var lessor models.Lessor
	utils.DB.First(&lessor, id)
	return lessor
}

func lessorDeleteByIdUser(id uuid.UUID) {
	utils.DB.Where("user_id = ?", id).Delete(&models.Lessor{})
}
