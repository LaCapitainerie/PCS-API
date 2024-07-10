package controller

import (
	"PCS-API/middleware"
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// edl # Controller
//
// edl réceptionne toutes les requêtes ayant pour endpoint '/edl'
// Il les envoie aux fonctions services liés
func Edl(api *gin.RouterGroup) {
	edl := api.Group("/edl")
	edl.Use(middleware.AuthMiddleware())
	{
		edl.GET("/:id", service.GetEdl)
		edl.POST("/:id", service.PostEdl)
	}
}
