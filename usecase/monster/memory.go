package monster

import (
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
		return nil, entity.ErrNotFound
	}

	return r.database[id], nil
}

func (r *memoryRepository) GetByName(name string) (*entity.Monster, error) {
	for _, m := range r.database {
		if m.Name == name {
			return m, nil
		}
	}

	return nil, entity.ErrNotFound
}

func (r *memoryRepository) GetBySetting(setting string) ([]*entity.Monster, error) {
	var monsters []*entity.Monster

	for _, m := range r.database {
		if m.Setting == setting {
			monsters = append(monsters, m)
		}
	}

	if len(monsters) == 0 {
		return nil, entity.ErrNotFound
	}

	return monsters, nil
}

func (r *memoryRepository) GetByMonsterTags(tags []string) ([]*entity.Monster, error) {
	monsters := make(map[entity.ID]*entity.Monster)

	for _, tag := range tags {
		for _, m := range r.database {
			if containsTag(tag, m.MonsterTags) {
				monsters[m.ID] = m
			}
		}
	}

	if len(monsters) == 0 {
		return nil, entity.ErrNotFound
	}

	return mapToSlice(monsters), nil
}

func (r *memoryRepository) GetByAttackTags(tags []string) ([]*entity.Monster, error) {
	monsters := make(map[entity.ID]*entity.Monster)

	for _, tag := range tags {
		for _, m := range r.database {
			if containsTag(tag, m.AttackTags) {
				monsters[m.ID] = m
			}
		}
	}

	if len(monsters) == 0 {
		return nil, entity.ErrNotFound
	}

	return mapToSlice(monsters), nil
}

func (r *memoryRepository) GetAll() ([]*entity.Monster, error) {
	return mapToSlice(r.database), nil
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

func containsTag(tag string, tags []string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}

	return false
}

func mapToSlice(m map[entity.ID]*entity.Monster) []*entity.Monster {
	slice := make([]*entity.Monster, 0)

	for _, monster := range m {
		slice = append(slice, monster)
	}

	return slice
}
