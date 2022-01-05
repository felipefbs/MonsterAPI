package monster

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

func (s *Service) GetMonsterByID(id entity.ID) (*entity.Monster, error) {
	return s.repo.Get(id)
}

func (s *Service) GetMonsterByName(name string) (*entity.Monster, error) {
	return s.repo.GetByName(name)
}

func (s *Service) GetMonsterBySetting(setting string) ([]*entity.Monster, error) {
	monsters, err := s.repo.GetBySetting(setting)
	if err != nil {
		return nil, err
	}

	return monsters, nil
}

func (s *Service) GetMonsterByMonsterTags(tags []string) ([]*entity.Monster, error) {
	monsters, err := s.repo.GetByMonsterTags(tags)
	if err != nil {
		return nil, err
	}

	return monsters, nil
}

func (s *Service) GetMonsterByAttackTags(tags []string) ([]*entity.Monster, error) {
	monsters, err := s.repo.GetByAttackTags(tags)
	if err != nil {
		return nil, err
	}

	return monsters, nil
}

func (s *Service) GetAllMonster() ([]*entity.Monster, error) {
	all, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	if len(all) == 0 {
		return nil, entity.ErrNotFound
	}

	return all, nil
}

func (s *Service) UpdateMonster(monster *entity.Monster) error {
	if err := monster.Validate(); err != nil {
		return err
	}

	monster.UpdatedAt = time.Now()

	return s.repo.Update(monster)
}

func (s *Service) StoreMonster(
	name string,
	moves []string,
	instinct string,
	description string,
	attack string,
	attackTags []string,
	damage string,
	monsterTags []string,
	hp int32,
	armor int32,
	specialQualities []string,
	setting string,
	source string,
) (entity.ID, error) {
	m, err := entity.NewMonster(name,
		moves,
		instinct,
		description,
		attack,
		attackTags,
		damage,
		monsterTags,
		hp,
		armor,
		specialQualities,
		setting,
		source)
	if err != nil {
		return entity.NilID, err
	}
	s.repo.Store(m)

	return m.ID, err
}

func (s *Service) DeleteMonsterByID(id entity.ID) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *Service) DeleteMonsterByName(name string) error {
	m, err := s.repo.GetByName(name)
	if err != nil {
		return err
	}

	return s.repo.Delete(m.ID)
}
