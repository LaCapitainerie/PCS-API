package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ReviewGetAll(c *gin.Context) {
	var review models.Review
	var err error
	if err = c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reviews := repository.ReviewGetAll(review.IdTarget)
	c.JSON(http.StatusOK, gin.H{"review": reviews})
}

func ReviewPost(c *gin.Context) {
	var review models.Review
	var err error
	if err = c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if review.Mark < 0 || review.Mark > 5 {
		c.JSON(http.StatusOK, gin.H{"error": "Note invalide"})
		return
	}

	IdUserStr, _ := c.Get("idUser")
	idUser, _ := uuid.Parse(IdUserStr.(string))

	review.IdUser = idUser

	review = repository.ReviewSave(review)
	c.JSON(http.StatusOK, gin.H{"review": review})
}

func ReviewDelete(c *gin.Context) {
	IdUserStr, _ := c.Get("idUser")
	idUser, _ := uuid.Parse(IdUserStr.(string))
	var review models.Review
	var err error
	if err = c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review.IdUser = idUser
	repository.ReviewDelete(review)
	c.JSON(http.StatusOK, gin.H{})
}
