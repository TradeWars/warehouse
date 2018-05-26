package storage

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (mgr *Manager) ensurePlayerCollection() (err error) {
	mgr.players = mgr.db.C("players")

	err = mgr.players.EnsureIndex(mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	})
	if err != nil {
		return
	}

	return
}

// PlayerCreate creates a new player account in the database
func (mgr *Manager) PlayerCreate(player types.Player) (id bson.ObjectId, err error) {
	player.ID = bson.NewObjectId()

	err = mgr.players.Insert(player)
	if err != nil {
		err = errors.Wrap(err, "failed to insert")
		return
	}

	return player.ID, err
}

// PlayerGetByName returns a player object by name
func (mgr *Manager) PlayerGetByName(name string) (player types.Player, err error) {
	err = mgr.players.Find(bson.M{"name": name}).One(&player)
	return
}

// PlayerGetByID returns a player object by ID
func (mgr *Manager) PlayerGetByID(id bson.ObjectId) (player types.Player, err error) {
	err = mgr.players.Find(bson.M{"_id": id}).One(&player)
	return
}

// PlayerUpdate updates a player in the database by their ID
func (mgr *Manager) PlayerUpdate(id bson.ObjectId, player types.Player) (err error) {
	return mgr.players.Update(bson.M{"_id": id}, player)
}

// PlayerRemove removes a player in the database by their ID
func (mgr *Manager) PlayerRemove(id bson.ObjectId) (err error) {
	return mgr.players.Remove(bson.M{"_id": id})
}
