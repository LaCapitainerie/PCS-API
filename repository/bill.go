package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

func BillCreate(bill models.Bill) (models.Bill, error) {
	err := utils.DB.Create(&bill).Error
	return bill, err
}

func BillGetByReservationId(id uuid.UUID) (models.Bill, error) {

}
