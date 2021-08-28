package memory_test

import (
	"testing"
	"time"

	"github.com/felipefbs/MonsterAPI/src/entities"
	"github.com/felipefbs/MonsterAPI/src/repositories/memory"
	"github.com/google/uuid"
	"github.com/kylelemons/godebug/pretty"
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
		t.Error("Create monster failed", err.Error())
	}

	want := monster
	got, err := repo.GetByID(monster.ID)
	if err != nil {
		t.Error("Get monster failed", err.Error())
	}

	err = repo.DeleteByID(monster.ID)
	if err != nil {
		t.Error("Delete monster failed", err.Error())
	}

	if want != got {
		diff := pretty.Compare(want, got)
		t.Error("Test failed")
		pretty.Print(diff)
	}
}

func TestGetByName(t *testing.T) {
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
		t.Error("Create monster failed", err.Error())
	}

	want := monster
	got, err := repo.GetByName(monster.Name)
	if err != nil {
		t.Error("Get monster failed", err.Error())
	}

	err = repo.DeleteByName(monster.Name)
	if err != nil {
		t.Error("Delete monster failed", err.Error())
	}

	if want != got {
		diff := pretty.Compare(want, got)
		t.Error("Test failed")
		pretty.Print(diff)
	}
}

func TestGetBySetting(t *testing.T) {
	monsterA := &entities.Monster{
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

	monsterB := &entities.Monster{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		Name:             "Cave Rat",
		Moves:            []string{"Swarm", "Rip something (or someone) apart"},
		Instinct:         "To devour",
		Description:      "Who hasn’t seen a rat before? It’s like that, but nasty and big and not afraid of you anymore. Maybe this one was a cousin to that one you caught in a trap or the one you killed with a knife in that filthy tavern in Darrow. Maybe he’s looking for a little ratty revenge. ",
		Attack:           "Gnaw",
		AttackTags:       []string{"Close", " Messy"},
		Damage:           "d6  1 piercing",
		MonsterTags:      []string{"Horde", " Small"},
		HP:               7,
		Armor:            1,
		SpecialQualities: []string{""},
		Setting:          "Cavern Dwellers",
		Source:           "core rulebook",
	}

	setting := "Cavern Dwellers"

	err := repo.Store(monsterA)
	if err != nil {
		t.Error("Create monster A failed", err.Error())
	}

	err = repo.Store(monsterB)
	if err != nil {
		t.Error("Create monster B failed", err.Error())
	}

	got, err := repo.GetBySetting(setting)
	if err != nil {
		t.Error("Get monster by monster tags failed", err.Error())
	}

	if got[0].Setting != setting && got[1].Setting != setting {
		t.Errorf("Wanted monter settings: %s\\But received: %s", setting, got[0].Setting)
		t.Errorf("Wanted monter settings: %s\\But received: %s", setting, got[1].Setting)
	}

	repo.DeleteByID(monsterA.ID)
	repo.DeleteByID(monsterB.ID)
}

func TestGetByAttackTags(t *testing.T) {
	monsterA := &entities.Monster{
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
		AttackTags:       []string{"Close", "Reach"},
		Damage:           "d8+1",
		MonsterTags:      []string{"Group", " Large"},
		HP:               10,
		Armor:            3,
		SpecialQualities: []string{"Burrowing"},
		Setting:          "Cavern Dwellers",
		Source:           "core rulebook",
	}

	monsterB := &entities.Monster{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		Name:             "Cave Rat",
		Moves:            []string{"Swarm", "Rip something (or someone) apart"},
		Instinct:         "To devour",
		Description:      "Who hasn’t seen a rat before? It’s like that, but nasty and big and not afraid of you anymore. Maybe this one was a cousin to that one you caught in a trap or the one you killed with a knife in that filthy tavern in Darrow. Maybe he’s looking for a little ratty revenge. ",
		Attack:           "Gnaw",
		AttackTags:       []string{"Close", "Messy"},
		Damage:           "d6  1 piercing",
		MonsterTags:      []string{"Horde", " Small"},
		HP:               7,
		Armor:            1,
		SpecialQualities: []string{""},
		Setting:          "Cavern Dwellers",
		Source:           "core rulebook",
	}

	tags := []string{"Reach", "Messy"}

	err := repo.Store(monsterA)
	if err != nil {
		t.Error("Create monster A failed", err.Error())
	}

	err = repo.Store(monsterB)
	if err != nil {
		t.Error("Create monster B failed", err.Error())
	}

	got, err := repo.GetByAttackTags(tags)
	if err != nil {
		t.Error("Get monster by attack tags failed", err.Error())
	}

	if len(got) != 2 {
		t.Errorf("Wanted %d monsters, but got %d", 2, len(got))
	}

	repo.DeleteByID(monsterA.ID)
	repo.DeleteByID(monsterB.ID)
}

func TestGetByMonsterTags(t *testing.T) {
	monsterA := &entities.Monster{
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

	monsterB := &entities.Monster{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		Name:             "Cave Rat",
		Moves:            []string{"Swarm", "Rip something (or someone) apart"},
		Instinct:         "To devour",
		Description:      "Who hasn’t seen a rat before? It’s like that, but nasty and big and not afraid of you anymore. Maybe this one was a cousin to that one you caught in a trap or the one you killed with a knife in that filthy tavern in Darrow. Maybe he’s looking for a little ratty revenge. ",
		Attack:           "Gnaw",
		AttackTags:       []string{"Close", " Messy"},
		Damage:           "d6  1 piercing",
		MonsterTags:      []string{"Horde", " Small"},
		HP:               7,
		Armor:            1,
		SpecialQualities: []string{""},
		Setting:          "Cavern Dwellers",
		Source:           "core rulebook",
	}

	tags := []string{"Horde", "Group"}

	err := repo.Store(monsterA)
	if err != nil {
		t.Error("Create monster A failed", err.Error())
	}

	err = repo.Store(monsterB)
	if err != nil {
		t.Error("Create monster B failed", err.Error())
	}

	got, err := repo.GetByMonsterTags(tags)
	if err != nil {
		t.Error("Get monster by setting failed", err.Error())
	}

	if len(got) != 2 {
		t.Errorf("Wanted %d monsters, but got %d", 2, len(got))
	}

	repo.DeleteByID(monsterA.ID)
	repo.DeleteByID(monsterB.ID)
}

func TestGetAll(t *testing.T) {
	monsterA := &entities.Monster{
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

	monsterB := &entities.Monster{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		Name:             "Cave Rat",
		Moves:            []string{"Swarm", "Rip something (or someone) apart"},
		Instinct:         "To devour",
		Description:      "Who hasn’t seen a rat before? It’s like that, but nasty and big and not afraid of you anymore. Maybe this one was a cousin to that one you caught in a trap or the one you killed with a knife in that filthy tavern in Darrow. Maybe he’s looking for a little ratty revenge. ",
		Attack:           "Gnaw",
		AttackTags:       []string{"Close", " Messy"},
		Damage:           "d6  1 piercing",
		MonsterTags:      []string{"Horde", " Small"},
		HP:               7,
		Armor:            1,
		SpecialQualities: []string{""},
		Setting:          "Cavern Dwellers",
		Source:           "core rulebook",
	}

	err := repo.Store(monsterA)
	if err != nil {
		t.Error("Create monster A failed", err.Error())
	}

	err = repo.Store(monsterB)
	if err != nil {
		t.Error("Create monster B failed", err.Error())
	}

	got, err := repo.GetAll()
	if err != nil {
		t.Error("Get all monsters", err.Error())
	}

	if len(got) != 2 {
		t.Errorf("Wanted %d monsters, but got %d", 2, len(got))
	}

	repo.DeleteByID(monsterA.ID)
	repo.DeleteByID(monsterB.ID)
}

func TestUpdate(t *testing.T) {
	monster := &entities.Monster{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		Name:             "Cave Rat",
		Moves:            []string{"Swarm", "Rip something (or someone) apart"},
		Instinct:         "To devour",
		Description:      "Who hasn’t seen a rat before? It’s like that, but nasty and big and not afraid of you anymore. Maybe this one was a cousin to that one you caught in a trap or the one you killed with a knife in that filthy tavern in Darrow. Maybe he’s looking for a little ratty revenge. ",
		Attack:           "Gnaw",
		AttackTags:       []string{"Close", " Messy"},
		Damage:           "d6  1 piercing",
		MonsterTags:      []string{"Horde", " Small"},
		HP:               7,
		Armor:            1,
		SpecialQualities: []string{""},
		Setting:          "Cavern Dwellers",
		Source:           "core rulebook",
	}

	err := repo.Store(monster)
	if err != nil {
		t.Error("Create monster A failed", err.Error())
	}

	monsterUpdated := &entities.Monster{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		Name:             "Cave Big Rat",
		Moves:            []string{"Swarm", "Rip something (or someone) apart"},
		Instinct:         "To devour",
		Description:      "Who hasn’t seen a rat before? It’s like that, but nasty and big and not afraid of you anymore. Maybe this one was a cousin to that one you caught in a trap or the one you killed with a knife in that filthy tavern in Darrow. Maybe he’s looking for a little ratty revenge. ",
		Attack:           "Gnaw",
		AttackTags:       []string{"Close", " Messy"},
		Damage:           "d6  1 piercing",
		MonsterTags:      []string{"Horde", " Small"},
		HP:               7,
		Armor:            1,
		SpecialQualities: []string{""},
		Setting:          "Cavern Dwellers",
		Source:           "core rulebook",
	}

	err = repo.Update(monsterUpdated)
	if err != nil {
		t.Error("Create monster A failed", err.Error())
	}

	want := monsterUpdated.Name
	got, _ := repo.GetByID(monsterUpdated.ID)

	if got.Name != want {
		t.Errorf("Wanted monster name: %s, but got: %s", want, got.Name)
	}

}
