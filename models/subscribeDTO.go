package models

import "github.com/google/uuid"

/*
   id  UUID PRIMARY KEY,
   begin_date TIMESTAMP WITH TIME ZONE NOT NULL,
   end_date TIMESTAMP WITH TIME ZONE NOT NULL,
   traveler_id  UUID NOT NULL,
   subscribe_id  UUID NOT NULL,
   FOREIGN KEY (traveler_id) REFERENCES traveler(id),
   FOREIGN KEY (subscribe_id) REFERENCES subscribe(id)
*/

type SubscribeDTO struct {
	ID        uuid.UUID `json:"id"`
	BeginDate string    `json:"beginDate"`
	EndDate   string    `json:"endDate"`
	UserId    uuid.UUID `json:"user"`
	Type      string    `json:"type"` // 'bagpacker' ou 'explorator'
	Annuel    bool      `json:"annuel"`
}
