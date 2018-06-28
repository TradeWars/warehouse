package timeline

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	// PostgreSQL interface
	_ "github.com/lib/pq"
)

/*
CREATE database ssc;
\c ssc
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;

CREATE TABLE events (
  time TIMESTAMPTZ NOT NULL,
  data JSONB       NOT NULL,
  uuid UUID        NOT NULL
);

SELECT create_hypertable('events', 'time');
*/

// Manager provides access to collections and predefined CRUD functionality.
type Manager struct {
	config Config
	db     *sql.DB

	emit *sql.Stmt
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
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.Host, config.Port, config.Name, config.User, config.Pass))
	if err != nil {
		return
	}

	mgr.emit, err = mgr.db.Prepare(`INSERT INTO events VALUES($1, $2, $3)`)
	if err != nil {
		return
	}

	return
}

// Emit handles an event emitted from the client and stores it to the database
func (mgr *Manager) Emit(timestamp time.Time, data []byte) (err error) {
	_, err = mgr.emit.Exec(timestamp, data, uuid.New().String())
	return
}
