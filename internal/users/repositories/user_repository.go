package repositories

import (
	"errors"
	"library/internal/users/models"
	"sync"
)

type UserRepository struct {
	users  map[int64]*models.User
	mu     sync.RWMutex
	nextId int64
}

func NewUserRepository() models.UserRepository {
	return &UserRepository{
		users:  make(map[int64]*models.User),
		nextId: 1,
	}
}

// CreateUser implements [models.UserRepository].
func (u *UserRepository) CreateUser(user *models.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	user.ID = u.nextId
	u.nextId++
	u.users[user.ID] = user
	return nil
}

// DeleteUser implements [models.UserRepository].
func (u *UserRepository) DeleteUser(id int64) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	_, exists := u.users[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(u.users, id)
	return nil
}

// GetAllUsers implements [models.UserRepository].
func (u *UserRepository) GetAllUsers() ([]*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	users := make([]*models.User, 0, len(u.users))
	for _, user := range u.users {
		users = append(users, user)
	}

	return users, nil
}

// GetUser implements [models.UserRepository].
func (u *UserRepository) GetUser(id int64) (*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exists := u.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser implements [models.UserRepository].
func (u *UserRepository) UpdateUser(id int64, user *models.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exists := u.users[id]
	if !exists {
		return errors.New("user not found")
	}

	u.users[user.ID] = user

	return nil
}
