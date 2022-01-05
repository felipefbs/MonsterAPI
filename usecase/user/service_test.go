package user_test

import (
	"errors"
	"testing"
	"time"

	"github.com/felipefbs/MonsterAPI/entity"
	"github.com/felipefbs/MonsterAPI/usecase/user"
	"github.com/stretchr/testify/assert"
)

func newFixtureUser() *entity.User {
	return &entity.User{
		ID:        entity.NewID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Nickname:  "Ghamir",
		Email:     "ghamirih@gmail.com",
		Password:  "safepassword",
	}
}

func Test_Store(t *testing.T) {
	repo := user.NewMemoryRepository()
	service := user.NewService(repo)

	u := newFixtureUser()

	// Check if user is created if there is no error
	_, err := service.StoreUser(u.Email, u.Password, u.Nickname)
	assert.Nil(t, err)

	// Checking duplicated email
	_, err = service.StoreUser(u.Email, u.Password, u.Nickname)
	assert.NotNil(t, err)

	// Check if user is not created while passing only empty fields
	_, err = service.StoreUser("", "", "")
	assert.NotNil(t, err)
}

func Test_SearchAndFind(t *testing.T) {
	repo := user.NewMemoryRepository()
	service := user.NewService(repo)

	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2.Email = "bono.vox@email.com"

	uID, _ := service.StoreUser(u1.Email, u1.Password, u1.Nickname)
	_, _ = service.StoreUser(u2.Email, u2.Password, u2.Nickname)

	t.Run("Get an user by id", func(t *testing.T) {
		savedUser, err := service.GetUser(uID)

		assert.Nil(t, err)
		assert.Equal(t, u1.Nickname, savedUser.Nickname)
	})

	t.Run("Get an user that id do not exist", func(t *testing.T) {
		_, err := service.GetUser(entity.NilID)

		assert.NotNil(t, err)
	})

	t.Run("Get all users", func(t *testing.T) {
		all, err := service.GetAllUsers()

		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})
}

func Test_Update(t *testing.T) {
	repo := user.NewMemoryRepository()
	service := user.NewService(repo)

	user := newFixtureUser()
	id, err := service.StoreUser(user.Email, user.Password, user.Nickname)
	assert.Nil(t, err)

	savedUser, _ := service.GetUser(id)
	savedUser.Nickname = "felipefbs"

	err = service.UpdateUser(savedUser)
	assert.Nil(t, err)

	updatedUser, err := service.GetUser(id)
	assert.Nil(t, err)
	assert.Equal(t, "felipefbs", updatedUser.Nickname)

	_, err = service.GetUser(entity.NilID)
	assert.NotNil(t, err)
}

func Test_Delete(t *testing.T) {
	repo := user.NewMemoryRepository()
	service := user.NewService(repo)

	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2ID, _ := service.StoreUser(u2.Email, u2.Password, u2.Nickname)

	err := service.DeleteUser(u1.ID)
	assert.Equal(t, errors.New("user not found"), err)

	err = service.DeleteUser(u2ID)
	assert.Nil(t, err)
	_, err = service.GetUser(u2ID)
	assert.Equal(t, errors.New("not found"), err)

	u3 := newFixtureUser()
	id, _ := service.StoreUser(u3.Email, u3.Password, u3.Nickname)
	savedUser, _ := service.GetUser(id)
	savedUser.Monsters = []entity.ID{entity.NewID()}
	_ = service.UpdateUser(savedUser)
	err = service.DeleteUser(id)
	assert.Equal(t, errors.New("cannot be deleted"), err)
}
