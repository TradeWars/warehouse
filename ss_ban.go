package goss

import "time"

// Ban contains all the information for a banned player.
type Ban struct {
	Name     string        `redis:"name,omitempty"`
	Ipv4     string        `redis:"ipv4"`
	Date     time.Time     `redis:"date"`
	Reason   string        `redis:"reason"`
	By       string        `redis:"by"`
	Duration time.Duration `redis:"duration"`
	Active   bool          `redis:"active"`
}
