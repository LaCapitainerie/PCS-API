package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func TicketUpdateById(c *gin.Context) {
	var err error

	var ticket models.Ticket
	if err = c.BindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket, err = repository.TicketUpdateById(ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ticket": ticket})
}

func TicketCreate(c *gin.Context) {

	// JSON send is only

	var err error

	// Parse json from the body to a ticket struct
	var ticketDTO models.IssueMakerDTO
	if err = c.BindJSON(&ticketDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a chat between the user and the admin

	var chat models.Chat
	chat.ID = uuid.New()

	adminUuid, err := uuid.Parse("efc6adf0-dbc2-46b2-b56b-78a76ccb08b7")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	repository.CreateChat(chat, []models.ChatUser{
		{
			ChatID: chat.ID,
			UserID: adminUuid,
		},
		{
			ChatID: chat.ID,
			UserID: ticketDTO.UserID,
		},
	})

	// Create a ticket
	ticket := models.Ticket{
		ID:          uuid.New(),
		Type:        ticketDTO.Type,
		State:       models.TICKET_STATE_OPEN,
		Description: ticketDTO.Description,
		ChatId:      chat.ID,
	}

	// Create the ticket in the database

	ticket, err = repository.TicketCreate(ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Return the created ticket
	c.JSON(http.StatusOK, gin.H{"ticket": ticket})
}

/*
func GetEverythingAboutAChat(idChat string) struct {
	Chat      models.Chat
	ChatUsers []models.ChatUser
	Messages  []models.Message
	Tickets   models.Ticket
} {
*/
