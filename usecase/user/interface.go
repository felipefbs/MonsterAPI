package user

import "github.com/felipefbs/MonsterAPI/entity"

type Repository interface {
	Get(id entity.ID) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Store(user *entity.User) error
	Update(user *entity.User) error
	Delete(id entity.ID) error
	CheckEmail(email string) error
}

type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
	GetAllUsers() ([]*entity.User, error)
	StoreUser(email, password, nickname string) (entity.ID, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id entity.ID) error
}
