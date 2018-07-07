package storage

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/TradeWars/warehouse/types"
)

func (mgr *Manager) ensurePlayerCollection() (err error) {
	mgr.players = mgr.db.C("players")

	err = mgr.players.EnsureIndex(mgo.Index{
		Key:    []string{"account.name"},
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
	return player.ID, mgr.players.Insert(player)
}

// PlayerGetByName returns a player object by name
func (mgr *Manager) PlayerGetByName(name string) (player types.Player, err error) {
	err = mgr.players.Find(bson.M{"account.name": name}).One(&player)
	return
}

// PlayerGetByID returns a player object by ID
func (mgr *Manager) PlayerGetByID(id bson.ObjectId) (player types.Player, err error) {
	err = mgr.players.Find(bson.M{"_id": id}).One(&player)
	return
}

// PlayerUpdate updates a player in the database by their ID
func (mgr *Manager) PlayerUpdate(id bson.ObjectId, player types.Player) (err error) {
	return mgr.players.Update(bson.M{"_id": id}, bson.M{"$set": player})
}

// PlayerRemove removes a player in the database by their ID
func (mgr *Manager) PlayerRemove(id bson.ObjectId) (err error) {
	return mgr.players.Remove(bson.M{"_id": id})
}
