package entities

import (
	"time"

	"github.com/google/uuid"
)

// Monster definition
type Monster struct {
	ID               uuid.UUID `bson:"_id" json:"_id"`
	CreatedAt        time.Time `bson:"createdAt" json:"createdAt"`
	Name             string    `bson:"name" json:"name"`
	Moves            []string  `bson:"moves" json:"moves"`
	Instinct         string    `bson:"instinct" json:"instinct"`
	Description      string    `bson:"description" json:"description"`
	Attack           string    `bson:"attack" json:"attack"`
	AttackTags       []string  `bson:"attack_tags" json:"attack_tags"`
	Damage           string    `bson:"damage" json:"damage"`
	MonsterTags      []string  `bson:"monster_tags" json:"monster_tags"`
	HP               int32     `bson:"hp" json:"hp"`
	Armor            int32     `bson:"armor" json:"armor"`
	SpecialQualities []string  `bson:"special_qualities" json:"special_qualities"`
	Setting          string    `bson:"setting" json:"setting"`
	Source           string    `bson:"source" json:"source"`
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
	Delete(name string) error
}

type MonsterRepositoryInterface interface {
	GetByID(id string) (*Monster, error)
	GetByName(name string) (*Monster, error)
	GetBySetting(setting string) ([]*Monster, error)
	GetByMonsterTags(tags []string) ([]*Monster, error)
	GetByAttackTags(tags []string) ([]*Monster, error)
	GetAll() ([]*Monster, error)

	Update(monster *Monster) error
	Store(monster *Monster) error
	Delete(name string) error
}
