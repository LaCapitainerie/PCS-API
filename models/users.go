package models

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Mail               string    `json:"mail"`
	Password           string    `json:"password"`
	RegisterDate       time.Time `gorm:"type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP" json:"register_date"`
	LastConnectionDate time.Time `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"last_connection_date"`
}

func (Users) TableName() string {
	return "users"
}
