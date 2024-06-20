package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TicketGetAll(c *gin.Context) {
	tickets, err := repository.TicketGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	chatDTO := make([]models.ChatDTO, len(tickets))
	for i, ticket := range tickets {
		result := repository.GetEverythingAboutAChat(ticket.ChatId.String())
		chatDTO[i] = createChatDTOWithAttribut(result.Chat, result.Tickets, result.Users, result.Messages)
	}
	c.JSON(http.StatusOK, gin.H{"chat": chatDTO})
}

/*
func GetEverythingAboutAChat(idChat string) struct {
	Chat      models.Chat
	ChatUsers []models.ChatUser
	Messages  []models.Message
	Tickets   models.Ticket
} {
*/
