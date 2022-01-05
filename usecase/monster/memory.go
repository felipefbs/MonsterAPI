package monster

import (
	"errors"

	"github.com/felipefbs/MonsterAPI/entity"
	"github.com/google/uuid"
)

type memoryRepository struct {
	database map[uuid.UUID]*entity.Monster
}

func NewMemoryRepository() Repository {
	database := make(map[entity.ID]*entity.Monster)
	return &memoryRepository{database}
}

func (r *memoryRepository) Get(id entity.ID) (*entity.Monster, error) {
	if r.database[id] == nil {
		return nil, errors.New("not found")
	}

	return r.database[id], nil
}

func (r *memoryRepository) GetAll() ([]*entity.Monster, error) {
	monsters := make([]*entity.Monster, 0)

	for _, monster := range r.database {
		monsters = append(monsters, monster)
	}

	return monsters, nil
}

func (r *memoryRepository) Store(monster *entity.Monster) error {
	r.database[monster.ID] = monster

	return nil
}

func (r *memoryRepository) Update(monster *entity.Monster) error {
	_, err := r.Get(monster.ID)
	if err != nil {
		return err
	}

	r.database[monster.ID] = monster

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
