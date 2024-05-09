package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ChatPostMessage(c *gin.Context) {
	var chatDTO models.ChatDTO
	var err error
	var chat models.Chat
	var message models.Message

	idC, exist := c.Get("idUser")
	id := idC.(string)

	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	if err = c.BindJSON(&chatDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(chatDTO.Message) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "12"})
		return
	}
	message = chatDTO.Message[0]

	if !utils.IsInArrayString(id, chatDTO.UserId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "9"})
		return
	}

	if (message.Type != "text" && message.Type != "image") || message.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "12"})
		return
	}

	idChats := repository.VerifyExistenceChat(chatDTO.UserId)
	if len(idChats) < 1 {
		chat.ID = uuid.New()
		chatUser := make([]models.ChatUser, len(chatDTO.UserId))

		for i, _ := range chatUser {
			uuidUser, err := uuid.Parse(chatDTO.UserId[i])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "10"})
				return
			}
			chatUser[i] = models.ChatUser{
				ChatID: chat.ID,
				UserID: uuidUser,
			}
		}
		chat, err = repository.CreateChat(chat, chatUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "11"})
			return
		}
	} else {
		chat.ID = idChats[0]
	}

	message.ID = uuid.New()

	uuidUser, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "10"})
		return
	}
	message.UserId = uuidUser
	message.ChatId = chat.ID
	message, err = repository.CreateMessage(message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "13"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}
