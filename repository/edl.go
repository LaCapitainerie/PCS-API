package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

func GetEdl[edlT []models.Remarks](id string) edlT {
	var edl edlT
	utils.DB.Where("id_reservation = ?", id).Find(&edl)
	return edl
}

func CreateEdl(edl models.RemarksDTO) (models.RemarksDTO, error) {
	for _, e := range edl.Edl {
		if err := utils.DB.Create(&e).Error; err != nil {
			return edl, err
		}
	}
	return edl, nil
}
