// Package service spécifie le code "métier" de l'API
package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreateUser Crée un utilisateur, cette fonction ne peut être appelé par un contrôleur, elle est forcément
// appelé par une autre fonction service (CreateLessor,
func createUser(c *gin.Context) {
	var user models.Users
	var err error
	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uuid.New()
	user, err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"users": user})
}
