package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RoomID     primitive.ObjectID `bson:"roomID,omitempty" json:"roomID,omitempty"`
	UserID     primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	NumPersons int                `bson:"numPerson,omitempty" json:"numPerson,omitempty"`
	FromDate   time.Time          `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate   time.Time          `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
}

type BookRoomParams struct {
	FromDate   string `json:"fromDate"`
	TillDate   string `json:"tillDate"`
	NumPersons int    `json:"numPersons"`
}

func ParseStringToTime(timestr string) (time.Time, error) {
	formatTime, err := time.Parse("2006-01-02 15:04:05", timestr)
	if err != nil {
		return time.Now(), err
	}
	return formatTime, nil
}
