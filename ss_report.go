package goss

import "time"

// Report represents a report one player or the server has made against another
// player.
type Report struct {
	ID       string    `redis:"id,omitempty"`
	Name     string    `redis:"name"`
	Reason   string    `redis:"reason"`
	Date     time.Time `redis:"date"`
	Read     bool      `redis:"read"`
	Type     string    `redis:"type"`
	Posx     float32   `redis:"posx"`
	Posy     float32   `redis:"posy"`
	Posz     float32   `redis:"posz"`
	World    int       `redis:"world"`
	Interior int       `redis:"interior"`
	Info     string    `redis:"info"`
	By       string    `redis:"by"`
	Active   bool      `redis:"active"`
}
