package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Report represents a behaviour report made against a player
type Report struct {
	ID       bson.ObjectId `validate:"omitempty" json:"_id"`
	Name     string        `validate:"required" json:"name"`
	Reason   string        `validate:"required" json:"reason"`
	By       string        `validate:"required" json:"by"`
	Date     time.Time     `validate:"required" json:"date"`
	Read     bool          `validate:"required" json:"read"`
	Type     string        `validate:"required" json:"type"`
	Position Geo           `validate:"required" json:"position"`
	Metadata string        `validate:"required" json:"metadata"`
	Archived bool          `validate:"required" json:"archived"`
}

// ExampleReport returns an example report
func ExampleReport() Report {
	return Report{
		ID:       bson.NewObjectId(),
		Name:     "John",
		Reason:   "Health hack",
		By:       "Alice",
		Date:     time.Now().Add(-time.Hour),
		Read:     false,
		Type:     "AC",
		Position: Geo{PosX: 800.0, PosY: 1200.0, PosZ: 16.0},
		Metadata: "135.00000",
		Archived: false,
	}
}
