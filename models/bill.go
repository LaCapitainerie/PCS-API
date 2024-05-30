package models

import (
	"github.com/google/uuid"
	"time"
)

type Bill struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Price   float64   `gorm:"type:numeric(10,2);notnull" json:"price"`
	Date    time.Time `gorm:"type:timestamp;notnull" json:"date"`
	Type    string    `gorm:"type:varchar(64)" json:"type"`
	Content string    `gorm:"type:text" json:"content"`
}

func (Bill) TableName() string {
	return "bill"
}
