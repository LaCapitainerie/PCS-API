package service

import (
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
