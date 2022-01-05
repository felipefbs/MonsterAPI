package monster

import (
	"github.com/felipefbs/MonsterAPI/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) UseCase {
	return &Service{repo: repo}
}

func (s *Service) GetMonsterByID(id string) (*entity.Monster, error) {
	return nil, nil
}

func (s *Service) GetMonsterByName(name string) (*entity.Monster, error) {
	return nil, nil
}

func (s *Service) GetMonsterBySetting(setting string) ([]*entity.Monster, error) {
	return nil, nil
}

func (s *Service) GetMonsterByMonsterTags(tags []string) ([]*entity.Monster, error) {
	return nil, nil
}

func (s *Service) GetMonsterByAttackTags(tags []string) ([]*entity.Monster, error) {
	return nil, nil
}

func (s *Service) GetAllMonster() ([]*entity.Monster, error) {
	return nil, nil
}

func (s *Service) UpdateMonster(monster *entity.Monster) error {
	return nil
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
	return nil
}

func (s *Service) DeleteMonsterByName(name string) error {
	return nil
}
