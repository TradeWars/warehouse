package types

import (
	"time"

	"github.com/google/uuid"
)

// Event represents a thing that happened at a point in time.
type Event struct {
	Time time.Time `json:"time"`
	Data []byte    `json:"data"` // ssc never needs to unmarshal this data so it stays as a byte slice
	UUID uuid.UUID
}
