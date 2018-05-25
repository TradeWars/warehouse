package types

import (
	"github.com/google/uuid"
)

// Storer declares a set of CRUD functions for persisting and accessing data
type Storer interface {
	// Player account interface
	PlayerCreate(player Player) (err error)
	PlayerGetByName(name string) (player Player, err error)
	PlayerGetByID(id uuid.UUID) (player Player, err error)
	PlayerUpdate(id uuid.UUID, player Player) (err error)
}
