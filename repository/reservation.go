package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
	"time"
)

func ReservationGetAllByIdPropertyWithEndDateAfterADate(idProperty uuid.UUID, date time.Time) []models.Reservation {
	var reservations []models.Reservation
	utils.DB.Where("property_id = ? AND end_date > ?", idProperty, date).Find(&reservations)
	return reservations
}

func ReservationCreate(reservation models.Reservation) (models.Reservation, error) {
	err := utils.DB.Create(&reservation).Error
	return reservation, err
}
