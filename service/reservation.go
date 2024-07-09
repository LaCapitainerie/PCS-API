package service

import (
	"PCS-API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ReservationGetAll(c *gin.Context) {
	idUserStr, exist := c.Get("idUser")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
	}
	idUser, _ := uuid.Parse(idUserStr.(string))

	reservations, err := repository.ReservationGetAll(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
