package main

import (
	"PCS-API/controller"
	"PCS-API/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig()
	router := gin.Default()

	api := router.Group("/api")

	controller.User(api)

	err := router.Run(":" + utils.PortApp)
	if err != nil {
		return
	}
}
