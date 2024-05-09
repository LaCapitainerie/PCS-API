package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
	"github.com/google/uuid"
)

func VerifyExistenceChat(id []string) []uuid.UUID {
	var chatIds []uuid.UUID
	utils.DB.Model(&models.ChatUser{}).
		Select("chat_id").
		Where("user_id IN ?", id).
		Group("chat_id").
		Having("COUNT(DISTINCT user_id = ?", len(id)).
		Find(&chatIds)
	return chatIds
}

func CreateChat(chat models.Chat, users []models.ChatUser) (models.Chat, error) {
	result := utils.DB.Create(&chat)
	if result.Error != nil {
		return chat, result.Error
	}
	for _, item := range users {
		utils.DB.Create(&item)
		if result.Error != nil {
			return chat, result.Error
		}
	}
	return chat, nil
}

func CreateMessage(message models.Message) (models.Message, error) {
	result := utils.DB.Create(&message)
	if result.Error != nil {
		return message, result.Error
	}
	return message, nil
}
