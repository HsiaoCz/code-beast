package types

import (
	"time"
)

type Booking struct {
	ID         string    `bson:"_id,omitempty" json:"id,omitempty"`
	RoomID     string    `bson:"roomID,omitempty" json:"roomID,omitempty"`
	UserID     string    `bson:"userID,omitempty" json:"userID,omitempty"`
	NumPersons int       `bson:"numPerson,omitempty" json:"numPerson,omitempty"`
	FromDate   time.Time `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate   time.Time `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
	Canceled   bool      `bson:"canceled,omitempty" json:"canceled,omitempty"`
}
