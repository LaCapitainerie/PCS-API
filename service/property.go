package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"github.com/google/uuid"
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

func PostAProperty(c *gin.Context) {
	var property models.Property
	var err error
	if err = c.BindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idBrut, exist := c.Get("idUser")
	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}

	idUser, _ := uuid.Parse(idBrut.(string))

	if !repository.IsALessor(idUser) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "14"})
		return
	}

	if len(property.Name) > 1 {

	}
}
