package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

func BillCreate(bill models.Bill) (models.Bill, error) {
	err := utils.DB.Create(&bill).Error
	return bill, err
}
