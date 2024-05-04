package models

import (
	"github.com/google/uuid"
)

// Traveler est la structure spécifiant les données de la Traveler utilisé par le front web de l'application
type Traveler struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"Id"`
	First_name   string    `gorm:"type:varchar(255);not null" json:"first_name"`
	Last_name    string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Phone_number string    `gorm:"type:varchar(255);not null" json:"phone_number"`
	User_id      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
}

// TableName Traveler Spécifie à gorm le nom de la base de donnée
func (Traveler) TableName() string {
	return "traveler"
}
