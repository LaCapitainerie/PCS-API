package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ChatPostMessage(c *gin.Context) {
	var chatDTO models.ChatDTO
	var err error
	var chat models.Chat
	var message models.Message

	idC, exist := c.Get("idUser")
	id := idC.(string)

	if !exist {
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

	idChatStr, err := repository.VerifyExistenceChat(chatDTO.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "10"})
		return
	}

	idChat, err := uuid.Parse(idChatStr)
	if err != nil {
		chat.ID = uuid.New()
		chatUser := make([]models.ChatUser, len(chatDTO.UserId))

		for i := range chatUser {
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
		chat.ID = idChat
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

func createChatDTOWithAttribut(chat models.Chat, ticket models.Ticket, chatUser []models.ChatUser, message []models.Message) models.ChatDTO {
	chatUserStr := make([]string, len(chatUser))
	for i, v := range chatUser {
		chatUserStr[i] = v.UserID.String()
	}

	return models.ChatDTO{
		ID:      chat.ID,
		View:    chat.View,
		Ticket:  ticket,
		UserId:  chatUserStr,
		Message: message,
	}
}

func ChatGetAllMessages(c *gin.Context) {
	IDUSER, exist := c.Get("idUser")
	idUser := IDUSER.(string)
	idChat := c.Param("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}
	if verify := repository.VerifyExistenceUserInAChat(idUser, idChat); !verify {
		c.JSON(http.StatusBadRequest, gin.H{"error": "10"})
		return
	}

	chatFetch := repository.GetEverythingAboutAChat(idChat)
	if chatFetch.Chat.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "10"})
		return
	}

	chatDTO := createChatDTOWithAttribut(
		chatFetch.Chat,
		chatFetch.Tickets,
		chatFetch.ChatUsers,
		chatFetch.Messages)

	c.JSON(http.StatusOK, gin.H{"chat": chatDTO})
}

func GetAllChatByUser(c *gin.Context) {
	idBrut, exist := c.Get("idUser")
	id := idBrut.(string)

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
		return
	}

	chatsId := repository.GetAllChatByUser(id)
	chats := make([]models.ChatDTO, len(chatsId))
	for i := range chats {
		users := repository.GetAllChatUserOfAChat(chatsId[i].ChatID.String())
		chats[i].ID = chatsId[i].ChatID
		chats[i].UserId = users
	}
	c.JSON(http.StatusOK, gin.H{"chat": chats})
}

func GetChatConnect(c *gin.Context) {
	/*	idBrut, exist := c.Get("idUser")
		id := idBrut.(string)

		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
			return
		}

		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close()

		for {

		}
	*/
}