package user_test

import (
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

	// Check if user is not created while passing only empty fields
	_, err = service.StoreUser("", "", "")
	assert.NotNil(t, err)
}
