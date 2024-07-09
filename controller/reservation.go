package controller

import (
	"PCS-API/middleware"
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

func Reservation(api *gin.RouterGroup) {
	reservation := api.Group("/reservation")
	reservation.Use(middleware.AuthMiddleware())
	reservation.GET("/all", service.ReservationGetAll)
	reservationProperty(reservation)
	reservationCheckout(reservation)
}
