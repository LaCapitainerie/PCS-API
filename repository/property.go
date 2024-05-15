package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

// GetAllProperty
// Renvoie la liste de tous les "Property"
func GetAllProperty() []models.Property {
	var Propertys []models.Property
	if err := utils.DB.Find(&Propertys); err.Error != nil {
		panic("Unable to get Propertys " + err.Error.Error())
	}
	return Propertys
}

func PropertyCreate(property models.Property) (models.Property, error) {
	err := utils.DB.Create(&property)
	return property, err.Error
}

/*func PropertyDeleteWithIdUserAndPropertyId(propertyId uuid.UUID, userId uuid.UUID) bool {
	utils.DB.Where("property_id = ?", userId, propertyId)
	return false
}

func propertyVerifExistenceById(lessorId uuid.UUID) bool {

}*/
