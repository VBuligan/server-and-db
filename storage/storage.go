package storage

import "database/sql"

// Storage * Instance of storage
type Storage struct {
	config *Config
	// * Database FileDescriptor
	db *sql.DB
}

// New * Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open * Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	return nil
}

// Close * Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}
