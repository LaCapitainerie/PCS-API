package models

import (
	"github.com/google/uuid"
	"time"
)

// Constantes permettant le typage d'un UsersDTO
const (
	AdminType    string = "admin"
	TravelerType string = "traveler"
	LessorType   string = "lessor"
	ProviderType string = "provider"
)

// UsersDTO structure spécifiant un Data Transfer Object un objet de donnée de transfert avec le front
// sert notamment à l'inscription et à la connexion
type UsersDTO struct {
	ID                 uuid.UUID `json:"id"`
	TypeUser           string    `json:"type"`
	Mail               string    `json:"mail"`
	Password           string    `json:"password"`
	RegisterDate       time.Time `json:"registerDate"`
	LastConnectionDate time.Time `json:"lastConnectionDate"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	PhoneNumber        string    `json:"phoneNumber"`
	Nickname           string    `json:"nickname"`
}