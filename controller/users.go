// Package controller
// package contenant le code qui réceptionne toutes les requêtes, les envoie au "middleware" et aux fonctions services associées
package controller

import (
	"github.com/gin-gonic/gin"
)

// Users réceptionne toutes les requêtes ayant pour endpoint '/users'
// Il les envoie aux fonctions services liées
func Users(api *gin.RouterGroup) {

}
