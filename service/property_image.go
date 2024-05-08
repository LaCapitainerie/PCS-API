package service

import (
	"PCS-API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllPropertyImage Récupère la liste de tous les PropertyImage
// @Summary PropertyImage
// @Schemes
// @Description Récupère tous les PropertyImage
// @Tags PropertyImage
// @Produce json
// @Success 200 {array} models.PropertyImage
// @Router /api/Property_image [get]
func GetAllPropertyImage(c *gin.Context) {
	Property_images := repository.GetAllPropertyImage()
	c.JSON(http.StatusOK, gin.H{"Property_image": Property_images})
}
