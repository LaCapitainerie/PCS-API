package controller

import (
	"PCS-API/middleware"
	"PCS-API/models"
	"PCS-API/service"
	"github.com/gin-gonic/gin"
)

func reservationCheckout(reservation *gin.RouterGroup) {
	reservationCheckoutGroup := reservation.Group("/checkout")
	reservationCheckoutGroup.Use(middleware.BlockTypeMiddleware(models.TravelerType))
	{
		reservationCheckoutGroup.POST("/session/:id/:quantity", service.CheckoutCreateSession)
	}
}
