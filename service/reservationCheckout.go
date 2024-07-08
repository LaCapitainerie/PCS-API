package service

import (
	"PCS-API/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	coupon2 "github.com/stripe/stripe-go/v78/coupon"
	"net/http"
	"strconv"
)

func reservationServiceExploratorCoupon(priceServicesAll int64) string {
	realReduc := int64(float64(priceServicesAll) * 0.05)
	params := &stripe.CouponParams{
		Duration:  stripe.String(string(stripe.CouponDurationForever)),
		AmountOff: stripe.Int64(realReduc),
		Currency:  stripe.String(string(stripe.CurrencyEUR)),
		Name:      stripe.String("Réduction abonnement explorator"),
	}

	coupon, _ := coupon2.New(params)
	return coupon.ID
}

func ReservationCheckoutCreateSession(c *gin.Context) {
	idUserStr, exist := c.Get("idUser")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "8"})
	}
	idUser, _ := uuid.Parse(idUserStr.(string))

	idStripe := c.Param("id")
	quantity, err := strconv.ParseInt(c.Param("quantity"), 10, 64)

	if err != nil || quantity < 1 || idStripe == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "28"})
		return
	}

	lineItems, priceServicesAll, idReservation := ReservationPropertyCreate(c, idUser)
	if idReservation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "28"})
		return
	}

	domain := "http://localhost:3000/stripe/success"
	params := &stripe.CheckoutSessionParams{
		LineItems:  lineItems,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?success=true&id_reservation=" + idReservation),
		CancelURL:  stripe.String(domain + "?canceled=true"),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
			Enabled: stripe.Bool(true),
		},
	}

	travelerId := repository.TravelerGetIdByUserId(idUser)
	subscribeTraveler := repository.SubscribeGetByTravelerId(travelerId)
	subscribe := repository.SubscribeTypeById(subscribeTraveler.SubscribeId)
	if subscribe.Type == "explorator" && subscribe.Annuel {
		couponId := reservationServiceExploratorCoupon(priceServicesAll)
		params.Discounts = []*stripe.CheckoutSessionDiscountParams{
			{
				Coupon: stripe.String(couponId),
			},
		}
	}

	s, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "28"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": s.URL})
}
