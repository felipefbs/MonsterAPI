package user

import (
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
	if err := s.repo.CheckEmail(email); err != nil {
		return entity.NilID, err
	}

	e, err := entity.NewUser(email, password, nickname)
	if err != nil {
		return entity.NilID, err
	}

	return e.ID, s.repo.Store(e)
}

func (s *Service) UpdateUser(user *entity.User) error {
	if err := user.Validate(); err != nil {
		return entity.ErrInvalidEnt
	}

	user.UpdatedAt = time.Now()

	return s.repo.Update(user)
}

func (s *Service) DeleteUser(id entity.ID) error {
	u, err := s.GetUser(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	if len(u.Monsters) > 0 {
		return entity.ErrCantDelete
	}

	return s.repo.Delete(id)
}
