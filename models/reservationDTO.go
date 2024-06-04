package models

type ReservationDTO struct {
	Reservation
	Bill Bill `json:"bill"`
}
