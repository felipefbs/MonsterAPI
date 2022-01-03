package entity

import (
	"errors"
	"time"
)

// Monster definition
type Monster struct {
	ID               ID        `json:"_id"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
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

func NewMonster(
	name string,
	moves []string,
	instinct string,
	description string,
	attack string,
	attackTags []string,
	damage string,
	monsterTags []string,
	hp int32,
	armor int32,
	specialQualities []string,
	setting string,
	source string,
) (*Monster, error) {
	m := &Monster{
		ID:               NewID(),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		Name:             name,
		Moves:            moves,
		Instinct:         instinct,
		Description:      description,
		Attack:           attack,
		AttackTags:       attackTags,
		Damage:           damage,
		MonsterTags:      monsterTags,
		HP:               hp,
		Armor:            armor,
		SpecialQualities: specialQualities,
		Setting:          setting,
		Source:           source,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m Monster) Validate() error {
	if m.Name == "" ||
		m.Moves == nil ||
		m.Instinct == "" ||
		m.Description == "" ||
		m.Attack == "" ||
		m.AttackTags == nil ||
		m.Damage == "" ||
		m.MonsterTags == nil ||
		m.HP < 0 ||
		m.Armor < 0 ||
		m.SpecialQualities == nil ||
		m.Setting == "" ||
		m.Source == "" {
		return errors.New("invalid entity")
	}
	return nil
}
