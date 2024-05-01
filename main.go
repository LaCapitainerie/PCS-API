// Initialise l'ensemble de l'API
package main

import (
	"PCS-API/controller"
	"PCS-API/utils"
	"github.com/gin-gonic/gin"
)

// Initialise l'API
func main() {
	utils.LoadConfig()
	router := gin.Default()

	api := router.Group("/api")

	controller.Users(api)
	controller.Sidebar(api)

	err := router.Run(":" + utils.PortApp)
	if err != nil {
		return
	}
}
