package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Booking struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID       primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	RoomID       primitive.ObjectID `bson:"roomID,omitempty" json:"roomID,omitempty"`
	NumOfPersons int                `bson:"numOfPersons,omitempty" json:"numOfPersons,omitempty"`
	FromDate     time.Time          `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate     time.Time          `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
	Canceled     bool               `bson:"canceled" json:"canceled"`
}
