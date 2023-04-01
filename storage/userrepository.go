package storage

import "github.com/VBuligan/server-and-db/internal/app/models"

// UserRepository * instance of repository (model interface)
type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

// Create * Create User in db
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	return nil, nil
}

// FindByLogin * Find user by login
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	return nil, false, nil
}

// * Select all users
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	return nil, nil
}
