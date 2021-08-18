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
	return nil, nil
}

func (r *memoryRepository) GetByName(name string) (*entities.Monster, error) {
	return nil, nil
}

func (r *memoryRepository) GetBySetting(setting string) ([]*entities.Monster, error) {
	return nil, nil
}

func (r *memoryRepository) GetByMonsterTags(tags []string) ([]*entities.Monster, error) {
	return nil, nil
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
	return nil
}

func (r *memoryRepository) DeleteByID(id uuid.UUID) error {
	return nil
}

func (r *memoryRepository) DeleteByName(name string) error {
	return nil
}
