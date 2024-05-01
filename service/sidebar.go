package service

import (
	"PCS-API/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllSidebar Gère dans le servicve l'interaction avec le repository pour récupérer tous les sidebar
func GetAllSidebar(c *gin.Context) {
	sidebars := repository.GetAllSidebar()
	c.JSON(http.StatusOK, gin.H{"Sidebar": sidebars})
}
