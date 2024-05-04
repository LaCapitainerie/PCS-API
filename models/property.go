package models

import (
	"github.com/google/uuid"
)

// Property est la structure spécifiant les données de la property utilisé par le front web de l'application
type Property struct {
	ID                       uuid.UUID `gorm:"type:uuid;primaryKey" json:"ID"`
	Name                     string    `json:"Name"`
	Type                     string    `json:"type"`
	Price                    float32   `json:"Price"`
	Surface                  int       `json:"Surface"`
	Room                     int       `json:"Room"`
	Bathroom                 int       `json:"Bathroom"`
	Garage                   int       `json:"Garage"`
	Description              string    `json:"Description"`
	Adress                   string    `json:"Adress"`
	City                     string    `json:"City"`
	Zip_Code                 string    `json:"Zip_Code"`
	Country                  string    `json:"Country"`
	Administrator_validation bool      `json:"AdministrationValidation"`
	Lessor_id                uuid.UUID `gorm:"type:uuid" json:"Lessor_id"`
}

// TableName Property Spécifie à gorm le nom de la table
func (Property) TableName() string {
	return "property"
}
