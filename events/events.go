package events

import (
	"database/sql"
	"fmt"
	"time"

	// PostgreSQL interface
	_ "github.com/lib/pq"
)

/*
CREATE database tutorial;
\c tutorial
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;

CREATE TABLE events (
  time TIMESTAMPTZ NOT NULL,
  data JSON        NOT NULL,
  uuid UUID        NOT NULL
);

SELECT create_hypertable('events', 'time');
*/

// Manager provides access to collections and predefined CRUD functionality.
type Manager struct {
	config Config
	db     *sql.DB
}

// Config describes how to connect to the database
type Config struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

// New constructs a new events manager with the given configuration.
// This attempts to connect to a database endpoint
func New(config Config) (mgr *Manager, err error) {
	mgr = &Manager{
		config: config,
	}

	mgr.db, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s",
		config.Host, config.Port, config.User, config.Name))
	if err != nil {
		return
	}

	return
}

// Emit handles an event emitted from the client and stores it to the database
func (mgr *Manager) Emit(timestamp time.Time, event interface{}) (err error) {
	return
}
