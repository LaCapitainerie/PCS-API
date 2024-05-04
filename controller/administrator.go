package controller

import (
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// Admin # Controller
//
// Admin réceptionne toutes les requêtes ayant pour endpoint '/Admin'
// Il les envoie aux fonctions services liés
func Admin(api *gin.RouterGroup) {
	api.GET("/admin", service.GetAllAdmin)
}
