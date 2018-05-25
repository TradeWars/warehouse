package types

import (
	"time"

	"github.com/google/uuid"
)

// Player represents a player in the game and all their data
type Player struct {
	ID           uuid.UUID `validate:"required" json:"id"`
	Name         string    `validate:"required" json:"name"`
	Pass         string    `validate:"required" json:"pass"`
	Ipv4         uint32    `validate:"required" json:"ipv4"`
	Alive        *bool     `validate:"required" json:"alive"`
	Registration time.Time `validate:"required" json:"regdate"`
	LastLogin    time.Time `validate:"required" json:"lastlog"`
	LastSpawn    time.Time `validate:"omitempty" json:"spawntime,omitempty"`
	TotalSpawns  int32     `validate:"required" json:"spawns"`
	Warnings     int32     `validate:"required" json:"warnings"`
	Gpci         string    `validate:"required" json:"gpci"`
	Archived     bool      `validate:"omitempty" json:"archived,omitempty"`
}

// ExamplePlayer returns an example object of a player
func ExamplePlayer() Player {
	alive := true
	return Player{
		ID:           uuid.New(),
		Name:         "John",
		Pass:         "[whirlpool hash of password]",
		Ipv4:         1544996175,
		Alive:        &alive,
		Registration: time.Now().Add(-time.Hour * 24),
		LastLogin:    time.Now().Add(-time.Hour * 6),
		LastSpawn:    time.Now().Add(-time.Hour * 6),
		TotalSpawns:  3,
		Warnings:     1,
		Gpci:         "[gpci hash]",
		Archived:     false,
	}
}
