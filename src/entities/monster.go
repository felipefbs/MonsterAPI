package entities

import (
	"time"

	"github.com/google/uuid"
)

// Monster definition
type Monster struct {
	ID               uuid.UUID `json:"_id"`
	CreatedAt        time.Time `json:"createdAt"`
	Name             string    `json:"name"`
	Moves            []string  `json:"moves"`
	Instinct         string    `json:"instinct"`
	Description      string    `json:"description"`
	Attack           string    `json:"attack"`
	AttackTags       []string  `json:"attack_tags"`
	Damage           string    `json:"damage"`
	MonsterTags      []string  `json:"monster_tags"`
	HP               int32     `json:"hp"`
	Armor            int32     `json:"armor"`
	SpecialQualities []string  `json:"special_qualities"`
	Setting          string    `json:"setting"`
	Source           string    `json:"source"`
}

type MonsterUseCaseInterface interface {
	GetByID(id string) (*Monster, error)
	GetByName(name string) (*Monster, error)
	GetBySetting(setting string) ([]*Monster, error)
	GetByMonsterTags(tags []string) ([]*Monster, error)
	GetByAttackTags(tags []string) ([]*Monster, error)
	GetAll() ([]*Monster, error)

	Update(monster *Monster) error
	Store(monster *Monster) error

	DeleteByID(id uuid.UUID) error
	DeleteByName(name string) error
}

type MonsterRepositoryInterface interface {
	GetByID(id uuid.UUID) (*Monster, error)
	GetByName(name string) (*Monster, error)
	GetBySetting(setting string) ([]*Monster, error)
	GetByMonsterTags(tags []string) ([]*Monster, error)
	GetByAttackTags(tags []string) ([]*Monster, error)
	GetAll() ([]*Monster, error)

	Update(monster *Monster) error
	Store(monster *Monster) error

	DeleteByID(id uuid.UUID) error
	DeleteByName(name string) error
}
