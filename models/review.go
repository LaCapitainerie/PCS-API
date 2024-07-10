package models

import "github.com/google/uuid"

type Review struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Mark     float64   `gorm:"type:numeric(10,2);notnull" json:"mark"`
	Comment  string    `gorm:"type:text" json:"comment"`
	IdTarget uuid.UUID `gorm:"type:uuid;notnull" json:"idTarget"`
	IdUser   uuid.UUID `gorm:"type:uuid;notnull" json:"idUser"`
}

// TableName Review Spécifie à gorm le nom de la base de donnée
func (Review) TableName() string {
	return "review"
}
