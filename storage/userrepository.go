package storage

import (
	"fmt"
	"github.com/VBuligan/server-and-db/internal/app/models"
	"log"
)

// UserRepository * instance of repository (model interface)
type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

// Create * Create User in db
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

/*


 */

// FindByLogin * Find user by login
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	return nil, false, nil
}

// SelectAll * Select all users
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// * Куда будем читать
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}
