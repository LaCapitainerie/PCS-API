package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

func ReviewSave(review models.Review) models.Review {
	var existReview models.Review
	if err := utils.DB.Where("id_target = ? AND id_user = ?", review.IdTarget, review.IdUser).First(&existReview).Error; err != nil {
		review.ID = uuid.New()
		utils.DB.Create(&review)
	} else {
		utils.DB.Model(&models.Review{}).Where("id = ?", existReview.ID).Updates(models.Review{Comment: review.Comment, Mark: review.Mark})
		review.ID = existReview.ID
	}
	return review
}

func ReviewDelete(review models.Review) {
	utils.DB.Where("id_target = ? AND id_user = ?", review.IdTarget, review.IdUser).Delete(&models.Review{})
}

func ReviewGetAll(id uuid.UUID) []models.Review {
	var reviews []models.Review
	utils.DB.Where("id_target = ?", id.String()).Find(&reviews)
	return reviews
}
