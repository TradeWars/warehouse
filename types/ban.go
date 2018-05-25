package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Ban contains all the information for a banned player
type Ban struct {
	ID       bson.ObjectId `validate:"omitempty" json:"_id"`
	Name     string        `validate:"required" json:"name"`
	Ipv4     string        `validate:"required" json:"ipv4"`
	Date     time.Time     `validate:"required" json:"date"`
	Reason   string        `validate:"required" json:"reason"`
	By       string        `validate:"required" json:"by"`
	Position Geo           `validate:"required" json:"position"`
	Duration time.Duration `validate:"required" json:"duration"`
	Active   bool          `validate:"required" json:"active"`
}
