package repository

import (
	"PCS-API/models"
	"PCS-API/utils"
)

func VerifyExistenceChat(users []string) (string, error) {
	var chatId string
	subReq := utils.DB.Model(&models.ChatUser{}).
		Select("user_id").
		Where("user_id IN ?", users)

	err := utils.DB.Model(&models.ChatUser{}).
		Select("chat_id").
		Where("user_id IN (?)", subReq).
		Order("chat_id").
		Limit(1).
		Pluck("chat_id", &chatId).
		Error
	return chatId, err
}

func CreateChat(chat models.Chat, users []models.ChatUser) (models.Chat, error) {
	result := utils.DB.Create(&chat)
	if result.Error != nil {
		return chat, result.Error
	}
	for i, _ := range users {
		result = utils.DB.Create(&users[i])
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

func VerifyExistenceUserInAChat(idUser string, idChat string) bool {
	var count int64
	result := utils.DB.Model(&models.ChatUser{}).
		Where("chat_id = ? AND user_id = ?", idChat, idUser).
		Count(&count)
	if result.Error != nil {
		return false
	}
	return count > 0
}

func GetChat(idChat string) (models.Chat, error) {
	var chat models.Chat
	err := utils.DB.First(&chat, idChat).Error
	return chat, err
}

func GetAllChatUserOfAChat(idChat string) []string {
	var chatUsers []models.ChatUser

	utils.DB.Where("chat_id = ?", idChat).Find(&chatUsers)
	userId := make([]string, len(chatUsers))
	for i, _ := range chatUsers {
		userId[i] = chatUsers[i].UserID.String()
	}
	return userId
}

func GetAllMessageOfAChat(idChat string) ([]models.Message, error) {
	var message []models.Message
	err := utils.DB.Where("chat_id", idChat).Find(&message).Error
	return message, err
}

func GetTicketOfAChat(idChat string) (models.Ticket, error) {
	var ticket models.Ticket
	err := utils.DB.Where("chat_id", idChat).First(&ticket).Error
	return ticket, err
}

func GetEverythingAboutAChat(idChat string) struct {
	Chat      models.Chat
	ChatUsers []models.ChatUser
	Messages  []models.Message
	Tickets   models.Ticket
} {
	var result struct {
		Chat      models.Chat
		ChatUsers []models.ChatUser
		Messages  []models.Message
		Tickets   models.Ticket
	}

	utils.DB.Where("id = ?", idChat).First(&result.Chat)
	utils.DB.Where("chat_id = ?", idChat).Find(&result.ChatUsers)
	utils.DB.Where("chat_id = ?", idChat).Find(&result.Messages)
	utils.DB.Where("chat_id = ?", idChat).Find(&result.Tickets)

	return result
}

func GetAllChatByUser(id string) []models.ChatUser {
	var chatUsers []models.ChatUser
	utils.DB.Where("user_id = ?", id).Find(&chatUsers)
	return chatUsers
}
