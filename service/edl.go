package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEdl(c *gin.Context) {
	edl := repository.GetEdl(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"edl": edl})
}

func PostEdl(c *gin.Context) {
	var edl models.RemarksDTO
	if err := c.ShouldBindJSON(&edl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	edl, err := repository.CreateEdl(edl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"edl": edl})
}
