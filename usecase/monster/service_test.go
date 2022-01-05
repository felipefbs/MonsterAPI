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

	m := newFixtureMonster()

	_, err := service.StoreMonster(m.Name, m.Moves, m.Instinct, m.Description, m.Attack, m.AttackTags, m.Damage, m.MonsterTags, m.HP, m.Armor, m.SpecialQualities, m.Setting, m.Source)
	assert.Nil(t, err)

	m = &entity.Monster{}
	_, err = service.StoreMonster(m.Name, m.Moves, m.Instinct, m.Description, m.Attack, m.AttackTags, m.Damage, m.MonsterTags, m.HP, m.Armor, m.SpecialQualities, m.Setting, m.Source)
	assert.NotNil(t, err)
}

func Test_SearchAndFind(t *testing.T) {
	repo := monster.NewMemoryRepository()
	service := monster.NewService(repo)

	t.Run("Get all monsters with empty database", func(t *testing.T) {
		allMonsters, err := service.GetAllMonster()

		assert.NotNil(t, err)
		assert.Equal(t, len(allMonsters), 0)
	})

	m := newFixtureMonster()

	savedID, err := service.StoreMonster(m.Name, m.Moves, m.Instinct, m.Description, m.Attack, m.AttackTags, m.Damage, m.MonsterTags, m.HP, m.Armor, m.SpecialQualities, m.Setting, m.Source)
	assert.Nil(t, err)
	_, err = service.StoreMonster(m.Name, m.Moves, m.Instinct, m.Description, m.Attack, m.AttackTags, m.Damage, m.MonsterTags, m.HP, m.Armor, m.SpecialQualities, m.Setting, m.Source)
	assert.Nil(t, err)

	t.Run("Get monster by id", func(t *testing.T) {
		savedMonster, err := service.GetMonsterByID(savedID)

		assert.Nil(t, err)
		assert.Equal(t, m.Name, savedMonster.Name)
	})

	t.Run("Get monster by id testing do not exist", func(t *testing.T) {
		_, err := service.GetMonsterByID(entity.NilID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Get monster by name", func(t *testing.T) {
		_, err := service.GetMonsterByName(m.Name)

		assert.Nil(t, err)
	})

	t.Run("Get monster by name that do not exist", func(t *testing.T) {
		_, err := service.GetMonsterByName("None")

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Get monsters by setting", func(t *testing.T) {
		allMonsters, err := service.GetMonsterBySetting(m.Setting)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(allMonsters))
	})

	t.Run("Get monsters by setting that do not exist", func(t *testing.T) {
		_, err := service.GetMonsterBySetting("none")

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Get monster by tags", func(t *testing.T) {
		allMonsters, err := service.GetMonsterByMonsterTags(m.MonsterTags)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(allMonsters))
	})

	t.Run("Get monster by tags that do not exist", func(t *testing.T) {
		_, err := service.GetMonsterByMonsterTags([]string{})

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Get monster by attack tags", func(t *testing.T) {
		allMonsters, err := service.GetMonsterByAttackTags(m.AttackTags)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(allMonsters))
	})

	t.Run("Get monster by attack tags that do not exist", func(t *testing.T) {
		_, err := service.GetMonsterByAttackTags([]string{})

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Get all monsters", func(t *testing.T) {
		allMonsters, err := service.GetAllMonster()

		assert.Nil(t, err)
		assert.Equal(t, len(allMonsters), 2)
	})
}

func Test_Update(t *testing.T) {
	repo := monster.NewMemoryRepository()
	service := monster.NewService(repo)

	m := newFixtureMonster()
	id, err := service.StoreMonster(m.Name, m.Moves, m.Instinct, m.Description, m.Attack, m.AttackTags, m.Damage, m.MonsterTags, m.HP, m.Armor, m.SpecialQualities, m.Setting, m.Source)
	assert.Nil(t, err)

	savedMonster, _ := service.GetMonsterByID(id)
	savedMonster.Armor = -1

	err = service.UpdateMonster(savedMonster)
	assert.NotNil(t, err)
	assert.Equal(t, err, entity.ErrInvalidEnt)

	savedMonster.Armor = 5

	err = service.UpdateMonster(savedMonster)
	assert.Nil(t, err)

	updateMonster, err := service.GetMonsterByID(id)
	assert.Nil(t, err)
	assert.Equal(t, updateMonster.Armor, int32(5))

	err = service.UpdateMonster(&entity.Monster{})
	assert.NotNil(t, err)
}

func Test_Delete(t *testing.T) {
	repo := monster.NewMemoryRepository()
	service := monster.NewService(repo)

	t.Run("Delete monster by id", func(t *testing.T) {
		m := newFixtureMonster()
		err := service.DeleteMonsterByID(m.ID)
		assert.Equal(t, entity.ErrNotFound, err)

		id, err := service.StoreMonster(m.Name, m.Moves, m.Instinct, m.Description, m.Attack, m.AttackTags, m.Damage, m.MonsterTags, m.HP, m.Armor, m.SpecialQualities, m.Setting, m.Source)
		assert.Nil(t, err)

		err = service.DeleteMonsterByID(id)
		assert.Nil(t, err)

		_, err = service.GetMonsterByID(id)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Delete monster by name", func(t *testing.T) {
		m := newFixtureMonster()
		err := service.DeleteMonsterByName(m.Name)
		assert.Equal(t, entity.ErrNotFound, err)

		id, err := service.StoreMonster(m.Name, m.Moves, m.Instinct, m.Description, m.Attack, m.AttackTags, m.Damage, m.MonsterTags, m.HP, m.Armor, m.SpecialQualities, m.Setting, m.Source)
		assert.Nil(t, err)

		err = service.DeleteMonsterByName(m.Name)
		assert.Nil(t, err)

		_, err = service.GetMonsterByID(id)
		assert.Equal(t, entity.ErrNotFound, err)
	})
}
