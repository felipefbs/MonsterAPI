package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        ID        `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Nickname  string    `json:"nickname"`
	Monsters  []ID      `json:"monsters"`
}

func NewUser(email, password, nickname string) (*User, error) {
	u := &User{
		ID:        NewID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     email,
		Nickname:  nickname,
	}

	pwd, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	u.Password = pwd

	if err = u.Validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Validate() error {
	if u.Email == "" || u.Password == "" || u.Nickname == "" {
		return errors.New("invalid entity")
	}

	return nil
}

func hashPassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (u *User) AddMonster(id ID) error {
	u.Monsters = append(u.Monsters, id)

	return nil
}

func (u *User) GetMonster(id ID) (ID, error) {
	for _, m := range u.Monsters {
		if m == id {
			return m, nil
		}
	}

	return id, errors.New("not found")
}

func (u *User) RemoveMonster(id ID) error {
	for i, m := range u.Monsters {
		if m == id {
			u.Monsters = append(u.Monsters[:i], u.Monsters[i+1:]...)

			return nil
		}
	}

	return errors.New("not found")
}
