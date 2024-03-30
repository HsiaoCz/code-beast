package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RoomID   primitive.ObjectID `bson:"roomID,omitempty" json:"roomID,omitempty"`
	UserID   primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	FromDate time.Time          `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate time.Time          `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
}
