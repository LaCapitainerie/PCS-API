package models

import (
	"github.com/google/uuid"
)

// Admin est la structure spécifiant les données de la Admin utilisé par le front web de l'application
type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"ID"`
	Site     string    `gorm:"type:varchar(64)" json:"Site"`
	Nickname string    `gorm:"type:varchar(64);notnull" json:"nickname"`
	User_id  uuid.UUID `gorm:"type:uuid;notnull" json:"user_id"`
}

// TableName Admin Spécifie à gorm le nom de la table
func (Admin) TableName() string {
	return "administrator"
}
