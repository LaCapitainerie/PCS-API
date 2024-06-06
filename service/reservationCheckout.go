package service

import (
	"github.com/gin-gonic/gin"
)

func createCheckoutSession(c *gin.Context) {
	/*	domain := "http://localhost:4242"
		params := &stripe.CheckoutSessionParams{
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				&stripe.CheckoutSessionLineItemParams{
					// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
					Price:    stripe.String("{{PRICE_ID}}"),
					Quantity: stripe.Int64(1),
				},
			},
			Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
			SuccessURL: stripe.String(domain + "?success=true"),
			CancelURL:  stripe.String(domain + "?canceled=true"),
			AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
				Enabled: stripe.Bool(true),
			},
		}

		s, err := session.New(params)
		if err != nil {
			log.Printf("session.New: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout session"})
			return
		}

		c.Redirect(http.StatusSeeOther, s.URL)*/
}
