package controller

import (
	"PCS-API/middleware"
	"PCS-API/models"
	"PCS-API/service"
	"github.com/gin-gonic/gin"
)

func Subscribe(api *gin.RouterGroup) {
	subscribeGroup := api.Group("/subscribe")
	subscribeGroup.Use(middleware.AuthMiddleware())
	subscribeGroup.Use(middleware.BlockTypeMiddleware(models.TravelerType))
	{
		subscribeGroup.POST("", service.SubscribeCreateSession)
		subscribeGroup.GET("/check", service.SubscribeSessionCheck)
	}
}
