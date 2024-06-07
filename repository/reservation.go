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

func ReservationValidation(id uuid.UUID) (models.Reservation, error) {
	var reservation models.Reservation
	result := utils.DB.Model(&models.Reservation{}).Where("id = ?", id).First(&reservation)

	if result.Error != nil && reservation.ID == uuid.Nil {
		return reservation, result.Error
	}
	err := utils.DB.Model(&models.Reservation{}).Where("id = ?", id).Update("Annulation", false).Error
	if err != nil {
		return reservation, err
	}
	return reservation, nil
}

func ReservationGetById(id uuid.UUID) (models.Reservation, error) {
	var reservation models.Reservation
	err := utils.DB.First(&reservation, id).Error
	return reservation, err
}
