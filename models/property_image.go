package models

import (
	"github.com/google/uuid"
)

// Property_image est la structure spécifiant les données de la Property_image utilisé par le front web de l'application
type Property_image struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"ID"`
	Path       string    `gorm:"type:varchar(255);notnull" json:"Path"`
	PropertyId uuid.UUID `gorm:"type:uuid;notnull" json:"property_id"`
}

// TableName Property_image Spécifie à gorm le nom de la table
func (Property_image) TableName() string {
	return "property_image"
}
