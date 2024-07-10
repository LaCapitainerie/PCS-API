package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"PCS-API/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

// Vérifie si on est éligible à la réduction
func subscribeReductionVerify(idUser uuid.UUID) bool {
	travelerId := repository.TravelerGetIdByUserId(idUser)
	subscribeTraveler := repository.SubscribeGetByTravelerId(travelerId)
	subscribeType := repository.SubscribeTypeById(subscribeTraveler.SubscribeId)
	return subscribeType.Annuel && subscribeType.Type == "explorator"
}

// SubscribeGetAll Renvoie tous les types d'abonnements
func SubscribeGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"subscribe": repository.SubscribeGetAll()})
}

// Vérifie et complète toutes les informations nécessaires pour créer la commande (type d'abonnement, annuel ou non...)
func subscribeVerify(c *gin.Context) (string, map[string]string) {
	var subscribeDTO models.SubscribeDTO
	var err error
	metadata := make(map[string]string)

	IdUser, exist := c.Get("idUser")
	if !exist {
		return "", metadata
	}

	if err = c.BindJSON(&subscribeDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "", metadata
	}

	subscribe := repository.SubscribeTypeGetByTypeAndAnnuel(subscribeDTO.Type, subscribeDTO.Annuel)
	if subscribe.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Abonnement invalide"})
		return "", metadata
	}

	// Crée les metadonnées utilisé par la session (permettant d'être réutiliser lors de l'application de la commande)

	metadata["id_user"] = IdUser.(string)
	metadata["type_subscribe"] = subscribe.Type

	if subscribe.Annuel {
		metadata["annuel"] = "true"
	} else {
		metadata["annuel"] = "false"
	}

	return subscribe.IdStripe, metadata
}

func SubscribeCreateSession(c *gin.Context) {
	idStripe, metadata := subscribeVerify(c)
	if idStripe == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Resultat invalide"})
		return
	}

	// Paramètre abonnement
	domain := "http://localhost:3000/subscribe"

	stripe.Key = "sk_test_51PNwOpRrur5y60cs5Yv2aKu9v6SrJHigo2cLgmxevvozEfzSDWFnaQhMwVH02RLc8R2xHdjkJ6QagZ7KDyYTVxZt00gadizteA"

	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(idStripe),
				Quantity: stripe.Int64(1),
			},
		},

		SuccessURL: stripe.String(domain + "?success=true&session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(domain + "?canceled=true"),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
			Enabled: stripe.Bool(true),
		},
		Metadata: metadata,
	}

	// Gestion des coupons
	idUser, _ := uuid.Parse(metadata["id_user"])
	if metadata["type_subscribe"] == "explorator" && metadata["annuel"] == "true" && subscribeReductionVerify(idUser) {
		params.Discounts = []*stripe.CheckoutSessionDiscountParams{
			{
				Coupon: stripe.String(utils.CouponSubscribeId),
			},
		}
	}

	// Génère la session
	s, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "28"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": s.URL})
}

// SubscribeSessionCheck Une fois le paiement effectuée, application de la commande
func SubscribeSessionCheck(c *gin.Context) {
	// Récupération des informations de session
	sessionID := c.Query("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id manquant"})
		return
	}
	s, err := session.Get(sessionID, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	annuel := s.Metadata["annuel"] == "true"
	idUser, _ := uuid.Parse(s.Metadata["id_user"])
	// Crée l'abonnement et renvoie l'utilisateur un subscribeDTO
	c.JSON(http.StatusOK, gin.H{"subscribe": createSubscribe(idUser, s.Metadata["type_subscribe"], annuel)})
}

// createSubscribe création de l'abonnement
func createSubscribe(idUser uuid.UUID, typeSubscribe string, annuel bool) models.SubscribeDTO {
	var subscribe models.SubscribeTraveler
	subscribe.ID = uuid.New()

	// Génération d'un subscribeTraveler
	travelerId := repository.TravelerGetIdByUserId(idUser)
	subscribeTraveler := repository.SubscribeGetByTravelerId(travelerId)
	timeNow := time.Now()

	// Vérifie si un abonnement est déjà existant, si oui et si c'est le même, alors on ajoute par dessus l'abonnement en cours
	if subscribeTraveler.ID != uuid.Nil && !subscribeTraveler.EndDate.After(timeNow) {
		subscribeLastType := repository.SubscribeTypeById(subscribeTraveler.SubscribeId)
		if subscribeLastType.Type == typeSubscribe {
			subscribe.BeginDate = subscribeTraveler.EndDate
		} else {
			repository.SubscribeDeleteDateNow(travelerId)
		}
	}
	subscribe.BeginDate = timeNow

	// Applique la date de fin
	if annuel {
		subscribe.EndDate = subscribe.BeginDate.AddDate(1, 0, 0)
	} else {
		subscribe.EndDate = subscribe.BeginDate.AddDate(0, 1, 0)
	}

	// Mets à jour dans la base de donnée
	subscribe.TravelerId = travelerId
	subscribe.SubscribeId = repository.SubscribeTypeGetByTypeAndAnnuel(typeSubscribe, annuel).ID
	repository.SubscribeCreateNewTraveler(subscribe)
	return subscribeDTOCreateWithSubscribe(subscribe, idUser, typeSubscribe, annuel)
}

// subscribeDTOCreateWithSubscribe Crée un subscribeDTO à partir d'un Subscribe.
func subscribeDTOCreateWithSubscribe(subscribe models.SubscribeTraveler, userId uuid.UUID, typeSub string, annuel bool) models.SubscribeDTO {
	return models.SubscribeDTO{
		ID:        subscribe.ID,
		BeginDate: subscribe.BeginDate.String(),
		EndDate:   subscribe.EndDate.String(),
		UserId:    userId,
		Type:      typeSub,
		Annuel:    annuel,
	}
}
