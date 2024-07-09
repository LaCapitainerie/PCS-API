package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"time"

	"github.com/google/uuid"
)

func ReservationGetAll(id uuid.UUID) ([]models.Reservation, error) {
	var reservations []models.Reservation
	var properties []models.Property

	lessor_id := GetLessorIdByUserId(id)

	// Get all properties for the lessor
	err := utils.DB.Where("lessor_id = ?", lessor_id).Find(&properties).Error
	if err != nil {
		return reservations, err
	}

	// Get all reservations for each property
	for _, property := range properties {
		propertyReservations, err := ReservationGetAllByIdProperty(property.ID)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, propertyReservations...)
	}

	return reservations, nil
}

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

func ReservationGetAllByIdProperty(id uuid.UUID) ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := utils.DB.Where("property_id = ? AND annulation = ?", id, false).Order("begin_date DESC").Find(&reservations).Error
	return reservations, err
}

func ReservationSetAnnulation(id uuid.UUID) error {
	return utils.DB.Model(&models.Reservation{}).Where("id = ?", id).Update("Annulation", true).Error
}

func ReservationSetReport(id uuid.UUID, begin time.Time, end time.Time) error {
	return utils.DB.Model(&models.Reservation{}).Where("id = ?", id).Updates(map[string]interface{}{
		"begin_date": begin,
		"end_date":   end,
	}).Error

}
