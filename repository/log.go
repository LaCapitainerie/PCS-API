package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"fmt"
)

// CreateLogEntry
// Crée une nouvelle entrée dans la table "log"
func CreateLogEntry(logEntry models.Log) {
	fmt.Printf("Creating log entry for user %s\n", logEntry.UserID)
	if err := utils.DB.Create(&logEntry); err.Error != nil {
		panic("Unable to create log entry " + err.Error.Error())
	}
}
