package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// createLessor crée un nouveau bailleur
// la fonction ne peut être appelé hors du package
func createLessor(c *gin.Context, userDTO models.UsersDTO) {
	user := convertUserDTOtoUser(userDTO)
	lessor := createLessorWithUserDTO(userDTO)
	var err error

	if len(lessor.FirstName) < 1 &&
		len(lessor.LastName) < 1 &&
		(len(lessor.PhoneNumber) < 6 && len(lessor.PhoneNumber) > 15) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "1"})
		return
	}

	user, err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lessor, err = repository.CreateLessor(lessor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDTO = createUserDTOwithUserAndLessor(user, lessor)
	userDTO.Password = ""
	c.JSON(http.StatusOK, gin.H{"users": userDTO})
}

// createLessorWithUserDTO Crée un bailleur à partir d'un UserDTO
func createLessorWithUserDTO(dto models.UsersDTO) models.Lessor {
	return models.Lessor{
		ID:          uuid.New(),
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		PhoneNumber: dto.PhoneNumber,
		UserId:      dto.ID,
	}
}
