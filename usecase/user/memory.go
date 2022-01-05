package user

import (
	"errors"

	"github.com/felipefbs/MonsterAPI/entity"
	"github.com/google/uuid"
)

type memoryRepository struct {
	database map[uuid.UUID]*entity.User
}

func NewMemoryRepository() Repository {
	database := make(map[entity.ID]*entity.User)
	return &memoryRepository{database}
}

func (r *memoryRepository) Get(id entity.ID) (*entity.User, error) {
	if r.database[id] == nil {
		return nil, errors.New("not found")
	}

	return r.database[id], nil
}

func (r *memoryRepository) GetAll() ([]*entity.User, error) {
	users := make([]*entity.User, 0)

	for _, user := range r.database {
		users = append(users, user)
	}

	return users, nil
}

func (r *memoryRepository) Store(user *entity.User) error {
	r.database[user.ID] = user

	return nil
}

func (r *memoryRepository) Update(user *entity.User) error {
	_, err := r.Get(user.ID)
	if err != nil {
		return err
	}

	r.database[user.ID] = user

	return nil
}

func (r *memoryRepository) Delete(id entity.ID) error {
	_, err := r.Get(id)
	if err != nil {
		return err
	}

	delete(r.database, id)

	return nil
}

func (r *memoryRepository) CheckEmail(email string) error {
	for _, user := range r.database {
		if user.Email == email {
			return errors.New("invalid email")
		}
	}

	return nil
}
