package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func serviceConvertToServiceDTO(service models.Service, userId uuid.UUID) models.ServiceDTO {
	return models.ServiceDTO{
		Service: service,
		UserId:  userId,
	}
}

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
	service.ProviderId = provider.ID
	service, _ = repository.ServiceCreateNewService(service)
	serviceDTO := serviceConvertToServiceDTO(service, idUser)
	c.JSON(http.StatusOK, gin.H{"service": serviceDTO})
}

func ServiceUpdate(c *gin.Context) {
	idService, _ := uuid.Parse(c.Param("id"))
	service, err := repository.ServiceGetWithServiceId(idService)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var serviceTransfert models.Service
	if err = c.BindJSON(&serviceTransfert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if serviceTransfert.Price > 1 &&
		(serviceTransfert.TargetCustomer != models.LessorType && serviceTransfert.TargetCustomer != models.TravelerType) &&
		serviceTransfert.RangeAction < 0 &&
		serviceTransfert.Description != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "19"})
		return
	}
	serviceTransfert.Lat, serviceTransfert.Lon, err = utils.LocateWithAddress(service.Address, service.City, service.ZipCode, service.Country)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	serviceTransfert.ID = service.ID
	serviceTransfert.ProviderId = service.ProviderId
	serviceTransfert = repository.ServiceUpdate(serviceTransfert)
	ServiceDTO := serviceConvertToServiceDTO(serviceTransfert,
		repository.ProviderGetUserIdWithProviderId(serviceTransfert.ProviderId))
	c.JSON(http.StatusOK, gin.H{"service": ServiceDTO})
}

func ServiceGetAll(c *gin.Context) {
	services, err := repository.ServiceGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"service": services})
}
