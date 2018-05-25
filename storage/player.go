package storage

import (
	"github.com/Southclaws/ScavengeSurviveCore/types"
	"github.com/google/uuid"
)

// PlayerCreate creates a new player account in the database
func (mgr *Manager) PlayerCreate(player types.Player) (err error) {
	return
}

// PlayerGetByName returns a player object by name
func (mgr *Manager) PlayerGetByName(name string) (player types.Player, err error) {
	return
}

// PlayerGetByID returns a player object by ID
func (mgr *Manager) PlayerGetByID(id uuid.UUID) (player types.Player, err error) {
	return
}

// PlayerUpdate updates a player in the database by their ID
func (mgr *Manager) PlayerUpdate(id uuid.UUID, player types.Player) (err error) {
	return
}
