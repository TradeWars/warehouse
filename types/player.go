package types

import (
	"time"

	"github.com/google/uuid"
)

// Player represents a player in the game and all their data
type Player struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Pass         string    `json:"pass"`
	Ipv4         uint32    `json:"ipv4"`
	Alive        bool      `json:"alive"`
	Registration time.Time `json:"regdate"`
	LastLogin    time.Time `json:"lastlog"`
	LastSpawn    time.Time `json:"spawntime,omitempty"`
	TotalSpawns  int32     `json:"spawns"`
	Warnings     int32     `json:"warnings"`
	Gpci         string    `json:"gpci"`
	Archived     bool      `json:"archived,omitempty"`
}

// ExamplePlayer returns an example object of a player
func ExamplePlayer() Player {
	return Player{
		ID:           uuid.New(),
		Name:         "John",
		Pass:         "[whirlpool hash of password]",
		Ipv4:         1544996175,
		Alive:        true,
		Registration: time.Now().Add(-time.Hour * 24),
		LastLogin:    time.Now().Add(-time.Hour * 6),
		LastSpawn:    time.Now().Add(-time.Hour * 6),
		TotalSpawns:  3,
		Warnings:     1,
		Gpci:         "[gpci hash]",
		Archived:     false,
	}
}
