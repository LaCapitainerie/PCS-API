package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"github.com/google/uuid"
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
	user := convertUserDTOtoUser(userDTO)
	traveler := createTravelerWithUserDTO(userDTO)
	var err error

	if len(traveler.FirstName) < 1 &&
		len(traveler.LastName) < 1 &&
		(len(traveler.PhoneNumber) < 6 && len(traveler.PhoneNumber) > 15) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "1"})
		return
	}

	user, err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	traveler, err = repository.CreateTraveler(traveler)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDTO = createUserDTOwithUserAndTraveler(user, traveler)
	userDTO.Password = ""
	c.JSON(http.StatusOK, gin.H{"users": userDTO})
}

// createTravelerWithUserDTO Crée un voyageur à partir d'un UserDTO
func createTravelerWithUserDTO(dto models.UsersDTO) models.Traveler {
	return models.Traveler{
		ID:          uuid.New(),
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		PhoneNumber: dto.PhoneNumber,
		UserId:      dto.ID,
	}
}
