package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ServiceCreateNewService(c *gin.Context) {
	var service models.Service
	var err error
	if err = c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if service.Price > 1 &&
		(service.TargetCustomer != models.LessorType && service.TargetCustomer != models.TravelerType) &&
		service.RangeAction < 0 &&
		service.Description != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "19"})
		return
	}
	service.Lat, service.Lon, err = utils.LocateWithAddress(service.Address, service.City, service.ZipCode, service.Country)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idBrut, exist := c.Get("idUser")
	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	idUser, err := uuid.Parse(idBrut.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "18"})
		return
	}

	//TODO: Penser à la sécurité (imaginons que le provider n'existe plus ?
	provider := repository.ProviderGetByUserId(idUser)
	service.ID = uuid.New()
	service.UserId = idUser
	service.ProviderId = provider.ID
	service = repository.ServiceCreateOrUpdateNewService(service)
	c.JSON(http.StatusOK, gin.H{"service": service})
}
