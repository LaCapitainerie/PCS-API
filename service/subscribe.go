package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"net/http"
)

func subscribeCreate(c *gin.Context) string {
	var subscribeDTO models.SubscribeDTO
	var err error

	if err = c.BindJSON(&subscribeDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return ""
	}

	abonnement := repository.SubscribeGetByTypeAndAnnuel(subscribeDTO.Type, subscribeDTO.Annuel)
	if abonnement.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Abonnement invalide"})
		return ""
	}

	return abonnement.IdStripe
}

func SubscribeCreateSession(c *gin.Context) {
	idStripe := subscribeCreate(c)
	if idStripe != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Resultat invalide"})
		return
	}

	domain := "http://localhost:3000/subscribe/success"
	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(idStripe),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(domain + "?success=true&session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(domain + "?canceled=true"),
	}

	s, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "28"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": s.URL})
}
