package service

import (
	"PCS-API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllProperty_image Récupère la liste de tous les Property_image
// @Summary Property_image
// @Schemes
// @Description Récupère tous les Property_image
// @Tags Property_image
// @Produce json
// @Success 200 {array} models.Property_image
// @Router /api/Property_image [get]
func GetAllProperty_image(c *gin.Context) {
	Property_images := repository.GetAllProperty_image()
	c.JSON(http.StatusOK, gin.H{"Property_image": Property_images})
}
