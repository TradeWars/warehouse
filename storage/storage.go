// Package storage provides a persistence layer for package server to use for
// storing, accessing and processing data.
package storage

// Manager provides access to collections and predefined CRUD functionality.
type Manager struct {
	config Config
}

// Config describes how to connect to the database
type Config struct {
}

// New constructs a new storage manager with the given configuration.
// This attempts to connect to a database endpoint
func New(config Config) (mgr *Manager) {
	mgr = &Manager{
		config: config,
	}
	return
}
