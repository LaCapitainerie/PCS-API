package service

import (
	"PCS-API/models"
	"github.com/google/uuid"
	"time"
)

func billCreate() {
	var bill models.Bill
	bill.ID = uuid.New()
	bill.Date = time.Now()

}

/*
type Bill struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Price   float64   `gorm:"type:numeric(10,2);notnull" json:"price"`
	Date    time.Time `gorm:"type:timestamp;notnull" json:"date"`
	Type    string    `gorm:"type:varchar(64)" json:"type"`
	Content string    `gorm:"type:text" json:"content"`
}
*/
