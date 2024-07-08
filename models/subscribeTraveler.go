package models

import (
	"github.com/google/uuid"
	"time"
)

/*

CREATE TABLE subscribe_traveler (
id  UUID PRIMARY KEY,
begin_date TIMESTAMP WITH TIME ZONE NOT NULL,
end_date TIMESTAMP WITH TIME ZONE NOT NULL,
traveler_id  UUID NOT NULL,
subscribe_id  UUID NOT NULL,
FOREIGN KEY (traveler_id) REFERENCES traveler(id),
FOREIGN KEY (subscribe_id) REFERENCES subscribe(id)
);
*/

type SubscribeTraveler struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BeginDate   time.Time `gorm:"type:timestamp;notnull" json:"beginDate"`
	EndDate     time.Time `gorm:"type:timestamp;notnull" json:"endDate"`
	TravelerId  uuid.UUID `gorm:"type:uuid;notnull" json:"travelerId"`
	SubscribeId uuid.UUID `gorm:"type:uuid;notnull" json:"subscribeId"`
}

// TableName SubscribeTraveler Spécifie à gorm le nom de la table
func (SubscribeTraveler) TableName() string {
	return "subscribe_traveler"
}
