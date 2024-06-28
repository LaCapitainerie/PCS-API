package service

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"net/http"
	"strconv"
)

func CheckoutCreateSession(c *gin.Context) {
	idStripe := c.Param("id")
	quantity, err := strconv.ParseInt(c.Param("quantity"), 10, 64)

	if err != nil || quantity < 1 || idStripe == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "28"})
		return
	}

	idReservation := ReservationPropertyCreate(c)
	if idReservation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "28"})
		return
	}

	domain := "http://localhost:3000/stripe/success"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(idStripe),
				Quantity: stripe.Int64(quantity),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?success=true&id_reservation=" + idReservation),
		CancelURL:  stripe.String(domain + "?canceled=true"),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
			Enabled: stripe.Bool(true),
		},
	}

	s, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "28"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": s.URL})
}
