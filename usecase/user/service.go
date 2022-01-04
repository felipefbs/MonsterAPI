package user

import (
	"errors"
	"time"

	"github.com/felipefbs/MonsterAPI/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) UseCase {
	return &Service{repo: repo}
}

func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}

func (s *Service) GetAllUsers() ([]*entity.User, error) {
	return s.repo.GetAll()
}

func (s *Service) StoreUser(email, password, nickname string) (entity.ID, error) {
	e, err := entity.NewUser(email, password, nickname)
	if err != nil {
		return entity.NilID, err
	}

	return e.ID, s.repo.Store(e)
}

func (s *Service) UpdateUser(user *entity.User) error {
	if err := user.Validate(); err != nil {
		return errors.New("invalid entity")
	}

	user.UpdatedAt = time.Now()

	return s.repo.Update(user)
}

func (s *Service) DeleteUser(id entity.ID) error {
	u, err := s.GetUser(id)
	if u == nil {
		return errors.New("user not found")
	}
	if err != nil {
		return err
	}
	if len(u.Monsters) > 0 {
		return errors.New("cannot be deleted")
	}

	return s.repo.Delete(id)
}
