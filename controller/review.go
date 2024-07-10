package controller

import (
	"PCS-API/middleware"
	"PCS-API/service"
	"github.com/gin-gonic/gin"
)

func Review(api *gin.RouterGroup) {
	api.GET("/review", service.ReviewGetAll)
	reviewGroup := api.Group("/review")
	reviewGroup.Use(middleware.AuthMiddleware())
	{
		reviewGroup.POST("", service.ReviewPost)
		reviewGroup.DELETE("", service.ReviewDelete)
	}
}
