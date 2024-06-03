package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// TODO: Mettre un champ "name" dans property notamment pour ce cas ci-dessous
func billGenerateContent(property models.Property, dto models.ReservationDTO) string {
	propertyName := repository.LessorGetByUserId(property.LessorId)
	content := fmt.Sprintf(
		"%d x  - %s de %s %s\n",
		int(dto.BeginDate.Sub(dto.EndDate).Hours()/24),
		property.Address,
		propertyName.FirstName,
		propertyName.LastName)
	return content
}

func billCreate(property models.Property, dto models.ReservationDTO) (models.Bill, error) {
	var bill models.Bill
	bill.ID = uuid.New()
	bill.Date = time.Now()
	bill.Statut = "success"
	bill.Content = billGenerateContent(property, dto)

	bill, err := repository.BillCreate(bill)

	return bill, err
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
