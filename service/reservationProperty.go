package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func reservationDTOCreate(reservation models.Reservation, bill models.Bill, service []models.ServiceDTO) models.ReservationDTO {
	return models.ReservationDTO{
		Reservation: reservation,
		Bill:        bill,
		Service:     service,
	}
}

func reservationDateIntersect(entry models.Reservation, allEntry []models.Reservation) bool {
	for _, value := range allEntry {
		if !value.Annulation {
			if (entry.BeginDate.Before(value.EndDate) || entry.BeginDate.Equal(value.EndDate)) &&
				(entry.EndDate.After(value.BeginDate) || entry.EndDate.Equal(value.BeginDate)) {
				return true
			}
		}
	}
	return false
}

func ReservationPropertyCreate(c *gin.Context) {
	str, exist := c.Get("idUser")
	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}

	var reservationDTO models.ReservationDTO
	var err error
	if err = c.BindJSON(&reservationDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUser, _ := uuid.Parse(str.(string))
	if idUser != reservationDTO.TravelerId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "18"})
		return
	}

	property, err := repository.PropertyGetById(reservationDTO.PropertyId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "21"})
		return
	}

	services, err := reservationGetAllService(&reservationDTO)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "25"})
		return
	}

	//TODO: Vérifie que la propriété est dans le rayon d'actiond de tous les services

	var reservation models.Reservation
	reservation.BeginDate = reservationDTO.BeginDate
	reservation.EndDate = reservationDTO.EndDate
	reservation.TravelerId = repository.TravelerGetIdByUserId(idUser)
	if reservation.TravelerId == uuid.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "24"})
		return
	}

	if reservation.BeginDate.After(reservation.EndDate) {
		tmp := reservation.EndDate
		reservation.EndDate = reservation.BeginDate
		reservation.BeginDate = tmp
	}

	timeNow := time.Now()
	date := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())

	if !reservation.BeginDate.After(date) {
		c.JSON(http.StatusConflict, gin.H{"error": "22"})
		return
	}

	if reservationDateIntersect(reservation, repository.ReservationGetAllByIdPropertyWithEndDateAfterADate(property.ID, date)) {
		c.JSON(http.StatusConflict, gin.H{"error": "22"})
		return
	}

	reservation.ID = uuid.New()
	reservation.PropertyId = property.ID

	reservation, err = repository.ReservationCreate(reservation)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	serviceDTO, err := reservationServiceListCreate(&reservationDTO, services, &reservation.ID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "25"})
		return
	}

	bill, err := billCreate(property, reservation)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "23"})
		return
	}

	reservationDTO = reservationDTOCreate(reservation, bill, serviceDTO)
	c.JSON(http.StatusOK, gin.H{"reservation": reservationDTO})
}

func ReservationGetAllOfAProperty(c *gin.Context) {

}
