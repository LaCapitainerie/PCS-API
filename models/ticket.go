package models

import (
	"github.com/google/uuid"
	"time"
)

// Ticket est la structure spécifiant les données des tickets, un type de chat
type Ticket struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Type        string    `gorm:"type:varchar(64);notnull" json:"type"`
	State       time.Time `gorm:"type:varchar(16);notnull" json:"state"`
	Description string    `gorm:"type:text;notnull" json:"description"`
	ChatId      uuid.UUID `gorm:"type:uuid" json:"chatId"`
}

// TableName Ticket Spécifie à gorm le nom de la table
func (Ticket) TableName() string {
	return "ticket"
}
