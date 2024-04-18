package controller

import (
	"PCS-API/service"
	"github.com/gin-gonic/gin"
)

func User(api *gin.RouterGroup) {
	api.POST("/users", service.CreateUser)
}
