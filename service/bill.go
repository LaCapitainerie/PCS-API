package service

import (
	"PCS-API/models"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func billGenerateContent(traveler models.Users, lessor models.Users, dto models.ReservationDTO) string {
	const strLayout = "%dx - %s de %s\n"
	content := fmt.Sprintf(
		strLayout,
		int(dto.BeginDate.Sub(dto.EndDate).Hours()/24),
	)
	return content
}

func billCreate(traveler models.Users, lessor models.Users, dto models.ReservationDTO) models.Bill {
	var bill models.Bill
	bill.ID = uuid.New()
	bill.Date = time.Now()
	bill.Statut = "success"
	bill.Content = billGenerateContent(traveler, lessor, dto)

	return bill
}

/*
type Bill struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Price   float64   `gorm:"type:numeric(10,2);notnull" json:"price"`
	Date    time.Time `gorm:"type:timestamp;notnull" json:"date"`
	Statut   string    `gorm:"type:varchar(64)" json:"type"`
	Content string    `gorm:"type:text" json:"content"`
}
*/
