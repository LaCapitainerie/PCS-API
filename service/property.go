package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"github.com/google/uuid"
	"image"
	"net/http"
	"path/filepath"

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
	err := c.Request.ParseMultipartForm(15 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No json"})
		return
	}

	var propertyDTO models.PropertyDTO
	var err error
	if err = c.BindJSON(&propertyDTO); err != nil {
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

	if err = c.BindJSON(&propertyDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(propertyDTO.Name) < 1 &&
		len(propertyDTO.Type) < 1 &&
		propertyDTO.Price < 1 &&
		propertyDTO.Surface < 8 &&
		propertyDTO.Room < 1 &&
		len(propertyDTO.ZipCode) < 5 &&
		len(propertyDTO.Address) < 1 &&
		len(propertyDTO.City) < 1 &&
		len(propertyDTO.Country) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "15"})
		return
	}

	// create Property

	var property models.Property
	property.ID = uuid.New()
	property.Name = propertyDTO.Name
	property.Type = propertyDTO.Type
	property.Price = propertyDTO.Price
	property.Surface = propertyDTO.Surface
	property.Room = propertyDTO.Room
	property.Bathroom = propertyDTO.Bathroom
	property.Garage = propertyDTO.Garage
	property.Description = propertyDTO.Description
	property.Address = propertyDTO.Address
	property.City = propertyDTO.City
	property.ZipCode = propertyDTO.ZipCode
	property.Country = propertyDTO.Country
	property.LessorId = repository.GetLessorIdByUserId(idUser)
	property.Position, err = utils.LocateWithAddress(
		property.Address,
		property.City,
		property.ZipCode,
		property.Country)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	property, err = repository.PropertyCreate(property)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Property non créer"})
		return
	}

	// Fichiers

	files := c.Request.MultipartForm.File["files"]
	if len(files) == 0 || len(files) > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "16"})
		return
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "16"})
			return
		}
		defer file.Close()

		extension := filepath.Ext(fileHeader.Filename)
		if extension != ".png" &&
			extension != ".jpg" &&
			extension != ".jpeg" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "16"})
			return
		}

		_, _, err = image.Decode(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "16"})
			return
		}

	}

	// DTO Création - Rendue
}
