// Package controller
// package contenant le code qui réceptionne toutes les requêtes, les envoie au "middleware" et aux fonctions services associées
package controller

import (
	"PCS-API/service"
	"github.com/gin-gonic/gin"
)

// Users # Controller
//
// Users réceptionne toutes les requêtes ayant pour endpoint '/users'
// Il les envoie aux fonctions services liés
func Users(api *gin.RouterGroup) {
	api.POST("/users", service.CreateUser)
}