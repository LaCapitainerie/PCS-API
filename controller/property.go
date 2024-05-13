package controller

import (
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// Property # Controller
//
// Property réceptionne toutes les requêtes ayant pour endpoint '/property'
// Il les envoie aux fonctions services liés
func Property(api *gin.RouterGroup) {
	api.GET("/property", service.GetAllProperty)
	api.POST("/property", service.PostAProperty)
}
