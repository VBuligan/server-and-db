package storage

// Storage * Instance of storage
type Storage struct {
	config *Config
}

// New * Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open * Open connection method
func (storage *Storage) Open() error {
	return nil
}

// Close * Close connection
func (storage *Storage) Close() {

}