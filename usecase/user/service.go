package user

import "github.com/felipefbs/MonsterAPI/entity"

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

func (s *Service) UpdateUser(user *entity.User) error { return nil }
func (s *Service) DeleteUser(id entity.ID) error      { return nil }
