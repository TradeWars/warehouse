// Package storage provides a persistence layer for package server to use for
// storing, accessing and processing data.
package storage

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/Southclaws/ScavengeSurviveCore/util"
)

// Manager provides access to collections and predefined CRUD functionality.
type Manager struct {
	config  Config
	session *mgo.Session
	db      *mgo.Database

	// Collections
	players *mgo.Collection
	admins  *mgo.Collection
	reports *mgo.Collection
	bans    *mgo.Collection
}

// Config describes how to connect to the database
type Config struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

// New constructs a new storage manager with the given configuration.
// This attempts to connect to a database endpoint
func New(config Config) (mgr *Manager, err error) {
	mgr = &Manager{
		config: config,
	}

	mgr.session, err = mgo.Dial(config.Host + ":" + config.Port)
	if err != nil {
		return
	}
	if config.User != "" {
		err = mgr.session.Login(&mgo.Credential{
			Username: config.User,
			Password: config.Pass,
		})
		if err != nil {
			return
		}
	}

	mgr.db = mgr.session.DB("ss")

	err = util.ErrSeq(
		mgr.ensurePlayerCollection(),
		mgr.ensureAdminCollection(),
		mgr.ensureReportCollection(),
		mgr.ensureBanCollection(),
	)

	return
}

// DeleteEverythingPermanently should only be used during testing!
func (mgr *Manager) DeleteEverythingPermanently() error {
	mgr.players.RemoveAll(bson.M{})
	mgr.admins.RemoveAll(bson.M{})
	mgr.reports.RemoveAll(bson.M{})
	mgr.bans.RemoveAll(bson.M{})
	return nil
}
