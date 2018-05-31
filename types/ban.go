package types

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Ban contains all the information for a banned player
type Ban struct {
	ID       bson.ObjectId `validate:"omitempty,required,len=12" json:"_id,omitempty" bson:"_id,omitempty"`
	Of       bson.ObjectId `validate:"required,len=12" json:"of_player_id" bson:"of_player_id"`
	By       bson.ObjectId `validate:"omitempty,len=12" json:"by_player_id,omitempty" bson:"by_player_id,omitempty"`
	Ipv4     uint32        `validate:"required" json:"ipv4" bson:"ipv4"`
	Date     time.Time     `validate:"required" json:"date" bson:"date"`
	Reason   string        `validate:"required" json:"reason" bson:"reason"`
	Position Geo           `validate:"required" json:"position" bson:"position"`
	Duration time.Duration `validate:"required" json:"duration" bson:"duration"`
	Archived bool          `validate:"required" json:"archived" bson:"archived"`
}

func ExampleBan() Ban {
	return Ban{
		ID:       bson.NewObjectId(),
		Of:       bson.NewObjectId(),
		By:       bson.NewObjectId(),
		Ipv4:     1544996175,
		Date:     time.Now().Add(-time.Hour * 6),
		Reason:   "hacking",
		Position: Geo{PosX: 300.0, PosY: 1200.0, PosZ: 16.0},
		Duration: time.Hour * 24,
		Archived: false,
	}
}
