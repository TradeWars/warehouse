package types

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Player represents a player in the game and all their data
type Player struct {
	ID           bson.ObjectId `validate:"omitempty,required" json:"_id" bson:"_id"`
	Name         string        `validate:"required,max=24" json:"name" bson:"name"`
	Pass         string        `validate:"required,len=128" json:"pass" bson:"pass"`
	Ipv4         uint32        `validate:"required" json:"ipv4" bson:"ipv4"`
	Alive        *bool         `validate:"required" json:"alive" bson:"alive"`
	Registration time.Time     `validate:"required" json:"regdate" bson:"regdate"`
	LastLogin    time.Time     `validate:"required" json:"lastlog" bson:"lastlog"`
	LastSpawn    time.Time     `validate:"omitempty" json:"spawntime,omitempty" bson:"spawntime,omitempty"`
	TotalSpawns  *int32        `validate:"required" json:"spawns" bson:"spawns"`
	Warnings     *int32        `validate:"required" json:"warnings" bson:"warnings"`
	Gpci         string        `validate:"required,len=40" json:"gpci" bson:"gpci"`
	Archived     bool          `validate:"omitempty" json:"archived,omitempty" bson:"archived,omitempty"`
}

// ExamplePlayer returns an example object of a player
func ExamplePlayer() Player {
	alive := true
	return Player{
		ID:           bson.NewObjectId(),
		Name:         "John",
		Pass:         "[whirlpool hash of password]",
		Ipv4:         1544996175,
		Alive:        &alive,
		Registration: time.Now().Add(-time.Hour * 24),
		LastLogin:    time.Now().Add(-time.Hour * 6),
		LastSpawn:    time.Now().Add(-time.Hour * 6),
		TotalSpawns:  &[]int32{3}[0],
		Warnings:     &[]int32{1}[0],
		Gpci:         "[gpci hash]",
		Archived:     false,
	}
}
