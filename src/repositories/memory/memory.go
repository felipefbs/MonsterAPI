package memory

import (
	"github.com/felipefbs/MonsterAPI/src/entities"
	"github.com/google/uuid"
)

type memoryRepository struct {
	database map[uuid.UUID]*entities.Monster
}

func NewMemoryRepository(database map[uuid.UUID]*entities.Monster) entities.MonsterRepositoryInterface {
	return &memoryRepository{database}
}

func (r *memoryRepository) GetByID(id uuid.UUID) (*entities.Monster, error) {
	return r.database[id], nil
}

func (r *memoryRepository) GetByName(name string) (*entities.Monster, error) {
	for _, v := range r.database {
		if v.Name == name {
			return v, nil
		}
	}

	return nil, nil
}

func (r *memoryRepository) GetBySetting(setting string) ([]*entities.Monster, error) {
	var data []*entities.Monster

	for _, v := range r.database {
		if v.Setting == setting {
			data = append(data, v)
		}
	}

	return data, nil
}

func (r *memoryRepository) GetByMonsterTags(tags []string) ([]*entities.Monster, error) {
	var data []*entities.Monster

	for _, t := range tags {
		for _, v := range r.database {
			for _, mt := range v.MonsterTags {
				if mt == t {
					data = append(data, v)
				}
			}
		}
	}

	return data, nil
}

func (r *memoryRepository) GetByAttackTags(tags []string) ([]*entities.Monster, error) {
	return nil, nil
}

func (r *memoryRepository) GetAll() ([]*entities.Monster, error) {
	return nil, nil
}

func (r *memoryRepository) Update(monster *entities.Monster) error {
	return nil
}

func (r *memoryRepository) Store(monster *entities.Monster) error {
	r.database[monster.ID] = monster

	return nil
}

func (r *memoryRepository) DeleteByID(id uuid.UUID) error {
	delete(r.database, id)

	return nil
}

func (r *memoryRepository) DeleteByName(name string) error {
	for _, v := range r.database {
		if v.Name == name {
			r.DeleteByID(v.ID)
		}
	}

	return nil
}
