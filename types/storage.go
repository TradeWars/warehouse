package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Storer declares a set of CRUD functions for persisting and accessing data
type Storer interface {
	// Player account interface
	PlayerCreate(player Player) (id bson.ObjectId, err error)
	PlayerGetByName(name string) (player Player, err error)
	PlayerGetByID(id bson.ObjectId) (player Player, err error)
	PlayerUpdate(id bson.ObjectId, player Player) (err error)
}
