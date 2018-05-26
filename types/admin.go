package types

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Admin represents a player who has been assigned staff role
type Admin struct {
	ID       bson.ObjectId `validate:"omitempty,required,len=12" json:"_id" bson:"_id"`
	PlayerID bson.ObjectId `validate:"required,len=12" json:"player_id" bson:"player_id"`
	Level    *int32        `validate:"required,max=5" json:"level" bson:"level"`
	Date     time.Time     `validate:"omitempty,required" json:"date,omitempty" bson:"date"`
}

// ExampleAdmin returns an example object of an admin description
func ExampleAdmin() Admin {
	var level int32 = 3
	return Admin{
		ID:       bson.NewObjectId(),
		PlayerID: bson.NewObjectId(),
		Date:     time.Now().Add(-time.Hour * 24),
		Level:    &level,
	}
}
