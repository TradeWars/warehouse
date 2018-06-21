// Package cache provides a temporary persistence layer for development purposes
package cache

import (
	"errors"
	"time"

	"github.com/Southclaws/ScavengeSurviveCore/types"
	"github.com/globalsign/mgo/bson"
)

// Cache provides access to collections and predefined CRUD functionality.
type Cache struct {
	players map[bson.ObjectId]types.Player
	admins  map[bson.ObjectId]*types.Admin
	reports map[bson.ObjectId]*types.Report
	bans    map[bson.ObjectId]*types.Ban
}

// New creates a new cache for mock DB calls
func New() (cache *Cache) {
	return &Cache{
		players: make(map[bson.ObjectId]types.Player),
		admins:  make(map[bson.ObjectId]*types.Admin),
		reports: make(map[bson.ObjectId]*types.Report),
		bans:    make(map[bson.ObjectId]*types.Ban),
	}
}

// Player interface

// PlayerCreate creates a new player account in the database
func (c *Cache) PlayerCreate(player types.Player) (id bson.ObjectId, err error) {
	for _, v := range c.players {
		if v.Account.Name == player.Account.Name {
			err = errors.New("player name already registered")
			return
		}
	}
	player.ID = bson.NewObjectId()
	c.players[player.ID] = player
	return player.ID, nil
}

// PlayerGetByName returns a player object by name
func (c *Cache) PlayerGetByName(name string) (player types.Player, err error) {
	for _, v := range c.players {
		if v.Account.Name == name {
			player = v
			return
		}
	}
	err = errors.New("not found")
	return
}

// PlayerGetByID returns a player object by ID
func (c *Cache) PlayerGetByID(id bson.ObjectId) (player types.Player, err error) {
	player, ok := c.players[id]
	if !ok {
		err = errors.New("not found")
	}
	return
}

// PlayerUpdate updates a player in the database by their ID
func (c *Cache) PlayerUpdate(id bson.ObjectId, player types.Player) (err error) {
	if _, ok := c.players[player.ID]; ok {
		c.players[player.ID] = player
	} else {
		err = errors.New("not found")
	}
	return
}

// PlayerRemove removes a player in the database by their ID
func (c *Cache) PlayerRemove(id bson.ObjectId) (err error) {
	delete(c.players, id)
	return
}

// Admin interface

// AdminSetLevel creates, updates or removes an admin record based on level
func (c *Cache) AdminSetLevel(id bson.ObjectId, level int32) (err error) {
	_, ok := c.admins[id]
	if !ok {
		if level == 0 {
			return
		}

		c.admins[id] = &types.Admin{
			ID:       bson.NewObjectId(),
			PlayerID: id,
			Level:    &level,
			Date:     time.Now(),
		}
	} else {
		if level == 0 {
			for k, v := range c.admins {
				if v.PlayerID == id {
					delete(c.admins, k)
					break
				}
			}
		} else {
			for k, v := range c.admins {
				if v.PlayerID == id {
					c.admins[k].Level = &level
				}
			}
		}
	}
	return
}

// AdminGetList returns a list of all admins
func (c *Cache) AdminGetList() (result []types.Admin, err error) {
	for _, v := range c.admins {
		result = append(result, *v)
	}
	return
}

// Report interface

// ReportCreate creates a report in the database
func (c *Cache) ReportCreate(report types.Report) (id bson.ObjectId, err error) {
	id = bson.NewObjectId()
	c.reports[id] = &report
	return
}

// ReportArchive sets archive status on a report
func (c *Cache) ReportArchive(id bson.ObjectId, archived bool) (err error) {
	c.reports[id].Archived = &archived
	return
}

// ReportGetList returns a list of reports based on search parameters
func (c *Cache) ReportGetList(pageSize, page int, archived, noRead bool, by, of bson.ObjectId, from, to *time.Time) (result []types.Report, err error) {
	for _, v := range c.reports {
		result = append(result, *v)
	}
	return
}

// ReportGet returns a specific report given an id
func (c *Cache) ReportGet(id bson.ObjectId) (result types.Report, err error) {
	r, ok := c.reports[id]
	if !ok {
		err = errors.New("not found")
		return
	}
	result = *r
	return
}

// misc

// DeleteEverythingPermanently should only be used during testing!
func (c *Cache) DeleteEverythingPermanently() (err error) {
	c.players = make(map[bson.ObjectId]types.Player)
	c.admins = make(map[bson.ObjectId]*types.Admin)
	c.reports = make(map[bson.ObjectId]*types.Report)
	c.bans = make(map[bson.ObjectId]*types.Ban)
	return
}
