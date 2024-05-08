// Package service spécifie le code "métier" de l'API
package service

import (
	"PCS-API/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"regexp"
	"unicode"
)

// CreateUser Crée un utilisateur
// @Summary User
// @Schemes
// @Description Crée un nouvel utilisateur
// @Tags Création
// @Produce json
// @Success 200 {object} model.UsersDTO
// @Router /api/user [post]
func CreateUser(c *gin.Context) {
	var user models.UsersDTO
	var err error
	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if validityPassword(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "1"})
		return
	}

	if validityEmail(user.Mail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "2"})
		return
	}

	user.ID = uuid.New()

	if user.TypeUser == models.TravelerType {
		createTraveler(c, user)
	} else if user.TypeUser == models.ProviderType {
		createProvider(c, user)
	} else if user.TypeUser == models.LessorType {
		createLessor(c, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "3"})
	}
}

// convertUserDTOtoUser Crée un utilisateur à partir d'un UserDTO
func convertUserDTOtoUser(userDTO models.UsersDTO) models.Users {
	return models.Users{
		ID:                 userDTO.ID,
		Mail:               userDTO.Mail,
		Password:           userDTO.Password,
		RegisterDate:       userDTO.RegisterDate,
		LastConnectionDate: userDTO.LastConnectionDate,
	}
}

// createUserDTOwithUserAndLessor Crée un userDTO à partir d'un utilisateurr et d'un bailleur
func createUserDTOwithUserAndLessor(users models.Users, lessor models.Lessor) models.UsersDTO {
	return models.UsersDTO{
		ID:                 lessor.ID,
		TypeUser:           models.LessorType,
		Mail:               users.Mail,
		Password:           users.Password,
		RegisterDate:       users.RegisterDate,
		LastConnectionDate: users.LastConnectionDate,
		FirstName:          lessor.FirstName,
		LastName:           lessor.LastName,
		PhoneNumber:        lessor.PhoneNumber,
	}
}

// validityPassword Vérifie la validité d'un mot de passe
func validityPassword(password string) bool {
	var check [4]bool
	if len(password) < 8 || len(password) > 128 {
		return false
	}
	for _, char := range password {
		if unicode.IsUpper(char) {
			check[0] = true
		} else if unicode.IsLower(char) {
			check[1] = true
		} else if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			check[2] = true
		} else if unicode.IsDigit(char) {
			check[3] = true
		}

		if check[0] == true &&
			check[1] == true &&
			check[2] == true &&
			check[3] == true {
			return true
		}
	}
	return false
}

func validityEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}
