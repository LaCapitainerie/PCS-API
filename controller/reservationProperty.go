package controller

import (
	"PCS-API/middleware"
	"PCS-API/models"
	"PCS-API/service"
	"github.com/gin-gonic/gin"
)

func reservationProperty(reservation *gin.RouterGroup) {
	reservationPropertyGroup := reservation.Group("/property")
	reservationPropertyGroup.Use(middleware.BlockTypeMiddleware(models.TravelerType))
	{
		reservationPropertyGroup.POST("", service.ReservationPropertyCreate)
	}
}
