package monster

import (
	"github.com/felipefbs/MonsterAPI/entity"
)

type Repository interface {
	Get(id entity.ID) (*entity.Monster, error)
	GetByName(name string) (*entity.Monster, error)
	GetBySetting(setting string) ([]*entity.Monster, error)
	GetByMonsterTags(tags []string) ([]*entity.Monster, error)
	GetByAttackTags(tags []string) ([]*entity.Monster, error)
	GetAll() ([]*entity.Monster, error)
	Store(user *entity.Monster) error
	Update(user *entity.Monster) error
	Delete(id entity.ID) error
}

type UseCase interface {
	GetMonsterByID(id entity.ID) (*entity.Monster, error)
	GetMonsterByName(name string) (*entity.Monster, error)
	GetMonsterBySetting(setting string) ([]*entity.Monster, error)
	GetMonsterByMonsterTags(tags []string) ([]*entity.Monster, error)
	GetMonsterByAttackTags(tags []string) ([]*entity.Monster, error)
	GetAllMonster() ([]*entity.Monster, error)
	UpdateMonster(monster *entity.Monster) error
	StoreMonster(name string, moves []string, instinct string, description string, attack string, attackTags []string, damage string, monsterTags []string, hp int32, armor int32, specialQualities []string, setting string, source string) (entity.ID, error)
	DeleteMonsterByID(id entity.ID) error
	DeleteMonsterByName(name string) error
}
