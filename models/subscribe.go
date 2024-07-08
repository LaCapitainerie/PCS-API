package models

import "github.com/google/uuid"

type Subscribe struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Type     string    `gorm:"type:varchar(64);notnull" json:"type"`
	Annuel   bool      `gorm:"type:boolean;notnull" json:"annuel"`
	Price    float64   `gorm:"type:numeric(10,2);notnull" json:"price"`
	IdStripe string    `gorm:"type:varchar(32);notnull" json:"idStripe"`
}

// TableName Subscribe Spécifie à gorm le nom de la table
func (Subscribe) TableName() string {
	return "subscribe"
}
