package types

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Report represents a behaviour report made against a player
type Report struct {
	ID       bson.ObjectId `validate:"omitempty,required" json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string        `validate:"required" json:"name" bson:"name"`
	Reason   string        `validate:"required" json:"reason" bson:"reason"`
	By       bson.ObjectId `validate:"omitempty,len=12" json:"by,omitempty" bson:"by,omitempty"`
	Date     time.Time     `validate:"required" json:"date" bson:"date"`
	Read     *bool         `validate:"required" json:"read" bson:"read"`
	Type     string        `validate:"required" json:"type" bson:"type"`
	Position Geo           `validate:"required" json:"position" bson:"position"`
	Metadata string        `validate:"omitempty" json:"metadata,omitempty" bson:"metadata,omitempty"`
	Archived *bool         `validate:"required" json:"archived" bson:"archived"`
}

// ExampleReport returns an example report
func ExampleReport() Report {
	f := false
	return Report{
		ID:       bson.NewObjectId(),
		Name:     "John",
		Reason:   "Health hack",
		By:       "Alice",
		Date:     time.Now().Add(-time.Hour),
		Read:     &f,
		Type:     "AC",
		Position: Geo{PosX: 800.0, PosY: 1200.0, PosZ: 16.0},
		Metadata: "135.00000",
		Archived: &f,
	}
}
