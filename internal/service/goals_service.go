package service

import (
	"time"

	"github.com/alwindoss/wink/internal/repository"
)

type Goal struct {
	Name         string
	Description  string
	Achieved     bool
	AchievedDate time.Time
	TargetDate   time.Time
}

type GoalsService interface {
	CreateGoal(g *Goal) (uint, error)
	CreateGoals(g []*Goal) (uint, error)
	UpdateGoal(g *Goal) error
	UpdateGoals(g []*Goal) error
	FetchGoal(id string) (*Goal, error)
	FetchGoals() ([]*Goal, error)
	DeleteGoal(g *Goal) error
}

func NewGoalsService(repo repository.GoalsRepository) GoalsService {
	s := &goalsService{
		repo: repo,
	}

	return s
}

type goalsService struct {
	repo repository.GoalsRepository
}

// CreateGoal implements GoalsService.
func (s *goalsService) CreateGoal(g *Goal) (uint, error) {
	repoGoal := repository.Goal{
		Name:         g.Name,
		Description:  g.Description,
		Achieved:     g.Achieved,
		AchievedDate: g.AchievedDate,
		TargetDate:   g.TargetDate,
	}
	id, err := s.repo.CreateGoal(&repoGoal)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// CreateGoals implements GoalsService.
func (*goalsService) CreateGoals(g []*Goal) (uint, error) {
	panic("unimplemented")
}

// DeleteGoal implements GoalsService.
func (*goalsService) DeleteGoal(g *Goal) error {
	panic("unimplemented")
}

// FetchGoal implements GoalsService.
func (g *goalsService) FetchGoal(id string) (*Goal, error) {
	panic("unimplemented")
}

// FetchGoals implements GoalsService.
func (g *goalsService) FetchGoals() ([]*Goal, error) {
	panic("unimplemented")
}

// UpdateGoal implements GoalsService.
func (*goalsService) UpdateGoal(g *Goal) error {
	panic("unimplemented")
}

// UpdateGoals implements GoalsService.
func (*goalsService) UpdateGoals(g []*Goal) error {
	panic("unimplemented")
}
