package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllTraveler Récupère la liste de tous les Traveler
// @Summary Traveler
// @Schemes
// @Description Récupère tous les Traveler
// @Tags administration
// @Produce json
// @Success 200 {array} models.Traveler
// @Router /api/Traveler [get]
func GetAllTraveler(c *gin.Context) {
	Travelers := repository.GetAllTraveler()
	c.JSON(http.StatusOK, gin.H{"Traveler": Travelers})
}

// createTraveler crée un nouveau voyageur
// la fonction ne peut être appelé hors du package
func createTraveler(c *gin.Context, userDTO models.UsersDTO) {

}
