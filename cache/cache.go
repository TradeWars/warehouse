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

func New() (cache *Cache) {
	return &Cache{
		players: make(map[bson.ObjectId]types.Player),
		admins:  make(map[bson.ObjectId]*types.Admin),
		reports: make(map[bson.ObjectId]*types.Report),
		bans:    make(map[bson.ObjectId]*types.Ban),
	}
}

// Player interface
func (c *Cache) PlayerCreate(player types.Player) (id bson.ObjectId, err error) {
	for _, v := range c.players {
		if v.Name == player.Name {
			err = errors.New("player name already registered")
			return
		}
	}
	player.ID = bson.NewObjectId()
	c.players[player.ID] = player
	return player.ID, nil
}
func (c *Cache) PlayerGetByName(name string) (player types.Player, err error) {
	for _, v := range c.players {
		if v.Name == name {
			player = v
			return
		}
	}
	err = errors.New("not found")
	return
}
func (c *Cache) PlayerGetByID(id bson.ObjectId) (player types.Player, err error) {
	player, ok := c.players[id]
	if !ok {
		err = errors.New("not found")
	}
	return
}
func (c *Cache) PlayerUpdate(id bson.ObjectId, player types.Player) (err error) {
	if _, ok := c.players[player.ID]; ok {
		c.players[player.ID] = player
	} else {
		err = errors.New("not found")
	}
	return
}
func (c *Cache) PlayerRemove(id bson.ObjectId) (err error) {
	delete(c.players, id)
	return
}

// Admin interface
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
func (c *Cache) AdminGetList() (result []types.Admin, err error) {
	for _, v := range c.admins {
		result = append(result, *v)
	}
	return
}

// Report interface
func (c *Cache) ReportCreate(report types.Report) (id bson.ObjectId, err error) {
	id = bson.NewObjectId()
	c.reports[id] = &report
	return
}
func (c *Cache) ReportArchive(id bson.ObjectId, archived bool) (err error) {
	c.reports[id].Archived = &archived
	return
}
func (c *Cache) ReportGetList(pageSize, page int, archived, noRead bool, by, of bson.ObjectId, from, to *time.Time) (result []types.Report, err error) {
	for _, v := range c.reports {
		result = append(result, *v)
	}
	return
}
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
func (c *Cache) DeleteEverythingPermanently() (err error) {
	c.players = make(map[bson.ObjectId]types.Player)
	c.admins = make(map[bson.ObjectId]*types.Admin)
	c.reports = make(map[bson.ObjectId]*types.Report)
	c.bans = make(map[bson.ObjectId]*types.Ban)
	return
}
