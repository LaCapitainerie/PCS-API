package service

import (
	"PCS-API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllProperty Récupère la liste de tous les Property
// @Summary Property
// @Schemes
// @Description Récupère tous les Property
// @Tags administration
// @Produce json
// @Success 200 {array} models.Property
// @Router /api/Property [get]
func GetAllProperty(c *gin.Context) {
	Propertys := repository.GetAllProperty()
	c.JSON(http.StatusOK, gin.H{"Property": Propertys})
}
