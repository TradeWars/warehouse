package types

import (
	"time"
)

// Ban contains all the information for a banned player
type Ban struct {
	Name     string        `json:"name"`
	Ipv4     string        `json:"ipv4"`
	Date     time.Time     `json:"date"`
	Reason   string        `json:"reason"`
	By       string        `json:"by"`
	Duration time.Duration `json:"duration"`
	Active   bool          `json:"active"`
}
