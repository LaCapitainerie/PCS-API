package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

func TicketGetAll() ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := utils.DB.Find(&tickets).Error
	return tickets, err
}
