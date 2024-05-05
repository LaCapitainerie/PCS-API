package service

import (
	"PCS-API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllAdmin Récupère la liste de tous les Admin
// @Summary Admin
// @Schemes
// @Description Récupère tous les Admin
// @Tags administration
// @Produce json
// @Success 200 {array} models.Admin
// @Router /api/Admin [get]
func GetAllAdmin(c *gin.Context) {
	Admins := repository.GetAllAdmin()
	c.JSON(http.StatusOK, gin.H{"Admin": Admins})
}
