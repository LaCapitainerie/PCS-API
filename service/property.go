package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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
	var err error
	var propertyDTO models.PropertyDTO
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
	property.Lat, property.Lon, err = utils.LocateWithAddress(
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

	// image

	var images []models.PropertyImage
	for _, value := range propertyDTO.Images {
		var image models.PropertyImage
		image.ID = uuid.New()
		image.Path = value
		image.PropertyId = property.ID
		image = repository.PropertyImageCreate(image)
		images = append(images, image)
	}

	// DTO Création - Rendue1
	propertyDTO = createPropertyDTOwithProperty(property, []models.PropertyImage{}, idUser)
	c.JSON(http.StatusOK, gin.H{"property": propertyDTO})
}

func createPropertyDTOwithProperty(property models.Property, images []models.PropertyImage, idUser uuid.UUID) models.PropertyDTO {
	imagesPath := make([]string, len(images))
	for i, v := range images {
		imagesPath[i] = v.Path
	}

	return models.PropertyDTO{
		ID:                      property.ID,
		Type:                    property.Type,
		Price:                   property.Price,
		Surface:                 property.Surface,
		Room:                    property.Room,
		Bathroom:                property.Bathroom,
		Garage:                  property.Garage,
		Description:             property.Description,
		Address:                 property.Address,
		City:                    property.City,
		ZipCode:                 property.ZipCode,
		Lon:                     property.Lon,
		Lat:                     property.Lat,
		Images:                  imagesPath,
		Country:                 property.Country,
		AdministratorValidation: property.AdministratorValidation,
		UserId:                  idUser,
	}
}

func PropertyDeleteById(c *gin.Context) {
	IDUSER, exist := c.Get("idUser")
	idUser, _ := uuid.Parse(IDUSER.(string))
	idProperty, _ := uuid.Parse(c.Param("id"))
	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}

	lessor := repository.LessorGetByUserId(idUser)
	supp := repository.PropertyDeleteWithIdUserAndPropertyId(idProperty, lessor.ID)
	if supp != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": supp.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

/*
// Fichiers

	files := c.Request.MultipartForm.File["files"]
	if len(files) == 0 || len(files) > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "16"})
		return
	}

	propertyDTO.Images = nil
	var propertyImages []models.PropertyImage
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

		var propertyImage models.PropertyImage
		propertyImage.ID = uuid.New()
		propertyImage.PropertyId = property.ID
		propertyImage.Path = "public/images/" + utils.GenerateUniqueFileName(fileHeader.Filename)
		repository.PropertyImageCreate(propertyImage)
		propertyImages = append(propertyImages, propertyImage)

		if err = c.SaveUploadedFile(fileHeader, propertyImage.Path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
*/
