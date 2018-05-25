package types

import (
	"time"

	"github.com/google/uuid"
)

// Report represents a behaviour report made against a player
type Report struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Reason   string    `json:"reason"`
	By       string    `json:"by"`
	Date     time.Time `json:"date"`
	Read     bool      `json:"read"`
	Type     string    `json:"type"`
	Position Geo       `json:"position"`
	Metadata string    `json:"metadata"`
	Archived bool      `json:"archived"`
}

// ExampleReport returns an example report
func ExampleReport() Report {
	return Report{
		ID:       uuid.New(),
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
