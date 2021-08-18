package memory_test

import (
	"testing"
	"time"

	"github.com/felipefbs/MonsterAPI/src/entities"
	"github.com/felipefbs/MonsterAPI/src/repositories/memory"
	"github.com/google/uuid"
)

var (
	db   = make(map[uuid.UUID]*entities.Monster, 5)
	repo = memory.NewMemoryRepository(db)
)

func TestGetByID(t *testing.T) {
	monster := &entities.Monster{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		Name:      "Ankheg",
		Moves: []string{
			"Undermine the ground",
			"Burst from the earth",
			"Spray forth acid, eating away at metal and flesh",
		},

		Instinct:         "To undermine",
		Description:      "A hide like plate armor and great crushing mandibles are problematic. A stomach full of acid that can burn a hole through a stone wall makes them all the worse. They’d be bad enough if they were proper insect-sized, but these things have the gall to be as long as any given horse. It’s just not natural! Good thing they tend to stick to one place? Easy for you to say—you don’t have an ankheg living under your corn field. ",
		Attack:           "Bite",
		AttackTags:       []string{"Close", " Reach"},
		Damage:           "d8+1",
		MonsterTags:      []string{"Group", " Large"},
		HP:               10,
		Armor:            3,
		SpecialQualities: []string{"Burrowing"},
		Setting:          "Cavern Dwellers",
		Source:           "core rulebook",
	}

	err := repo.Store(monster)
	if err != nil {
		t.Error("Erro na criação do monstro")
	}

	want := monster
	got, err := repo.GetByID(monster.ID)
	if err != nil {
		t.Error("Erro get do monstro")
	}

	err = repo.DeleteByName(monster.Name)
	if err != nil {
		t.Error("Erro get do monstro")
	}

	if want != got {
		t.Error("monster fail")
	}
}
