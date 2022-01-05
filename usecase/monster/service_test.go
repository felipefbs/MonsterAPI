package monster_test

import (
	"testing"
	"time"

	"github.com/felipefbs/MonsterAPI/entity"
	"github.com/felipefbs/MonsterAPI/usecase/monster"
	"github.com/stretchr/testify/assert"
)

func newFixtureMonster() *entity.Monster {
	return &entity.Monster{
		ID:        entity.NewID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
}

func Test_Store(t *testing.T) {
	repo := monster.NewMemoryRepository()
	service := monster.NewService(repo)

	u := newFixtureMonster()

	_, err := service.StoreMonster(u.Name, u.Moves, u.Instinct, u.Description, u.Attack, u.AttackTags, u.Damage, u.MonsterTags, u.HP, u.Armor, u.SpecialQualities, u.Setting, u.Source)
	assert.Nil(t, err)

	u = &entity.Monster{}
	_, err = service.StoreMonster(u.Name, u.Moves, u.Instinct, u.Description, u.Attack, u.AttackTags, u.Damage, u.MonsterTags, u.HP, u.Armor, u.SpecialQualities, u.Setting, u.Source)
	assert.NotNil(t, err)
}

func Test_SearchAndFind(t *testing.T) {
	repo := monster.NewMemoryRepository()
	service := monster.NewService(repo)

	u := newFixtureMonster()

	savedID, err := service.StoreMonster(u.Name, u.Moves, u.Instinct, u.Description, u.Attack, u.AttackTags, u.Damage, u.MonsterTags, u.HP, u.Armor, u.SpecialQualities, u.Setting, u.Source)
	assert.Nil(t, err)

	t.Run("Get one monster", func(t *testing.T) {
		savedMonster, err := service.GetMonsterByID(savedID)

		assert.Nil(t, err)
		assert.Equal(t, u.Name, savedMonster.Name)
	})
}
