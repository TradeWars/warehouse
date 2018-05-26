package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Admin represents a player who has been assigned staff role
type Admin struct {
	ID       bson.ObjectId `validate:"omitempty,required,len=24" json:"_id" bson:"_id"`
	PlayerID bson.ObjectId `validate:"required,len=24" json:"player_id"`
	Level    int32         `validate:"required,max=5" json:"level"`
	Date     time.Time     `validate:"required" json:"date"`
}

// ExampleAdmin returns an example object of an admin description
func ExampleAdmin() Admin {
	return Admin{
		ID:       bson.NewObjectId(),
		PlayerID: bson.NewObjectId(),
		Date:     time.Now().Add(-time.Hour * 24),
		Level:    3,
	}
}
