package controller

import (
	"PCS-API/middleware"
	"PCS-API/models"
	"PCS-API/service"
	"github.com/gin-gonic/gin"
)

func Reservation(api *gin.RouterGroup) {
	reservationGroup := api.Group("/reservation/property")
	reservationGroup.Use(middleware.AuthMiddleware())
	reservationGroup.Use(middleware.BlockTypeMiddleware(models.TravelerType))
	{
		reservationGroup.POST("", service.ReservationPropertyCreate)
	}
}
