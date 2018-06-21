package types

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Player represents a player in the game and all their data
//
// Each field (aside from ID) is an object that corresponds to a package in the
// gamemode that requires persistent storage for a player.
// When a new package is added, its data structure must first be added here. The
// structure at first may just be a `map[string]interface{}` for flexibility,
// however eventually the schema _should_ be defined in some way.
//
// Because this application is used by both Scavenge and Survive and the
// Sandbox gamemode, there may be fields here that don't apply to both gamemodes
// and will be left blank if unused.
type Player struct {
	ID      bson.ObjectId `validate:"omitempty,required" json:"_id"     bson:"_id"`
	Account Account       `validate:"required"           json:"account" bson:"account"`
	Spawn   Geo
}

// Account represents player account information such as password hash
type Account struct {
	Name         string    `validate:"required,max=24" json:"name"               bson:"name"`
	Pass         string    `validate:"required"        json:"pass"               bson:"pass"`
	Ipv4         string    `validate:"required"        json:"ipv4"               bson:"ipv4"`
	Registration time.Time `validate:"required"        json:"regdate"            bson:"regdate"`
	Gpci         string    `validate:"required,len=40" json:"gpci"               bson:"gpci"`
	Archived     bool      `validate:"omitempty"       json:"archived,omitempty" bson:"archived"`
}

// ExamplePlayer returns an example object of a player
func ExamplePlayer() Player {
	return Player{
		ID: bson.NewObjectId(),
		Account: Account{
			Name:         "John",
			Pass:         "[whirlpool hash of password]",
			Ipv4:         "191.24.25.16",
			Registration: time.Now().Add(-time.Hour * 24),
			Gpci:         "[gpci hash]",
			Archived:     false,
		},
	}
}
