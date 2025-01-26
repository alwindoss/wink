package service

import (
	"time"

	"github.com/alwindoss/wink/internal/repository"
)

type Goal struct {
	Name        string
	Description string
	TargetDate  *time.Time
}

type GoalsService interface {
	CreateGoal(g *Goal) error
	CreateGoals(g []*Goal) error
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
func (*goalsService) CreateGoal(g *Goal) error {
	panic("unimplemented")
}

// CreateGoals implements GoalsService.
func (*goalsService) CreateGoals(g []*Goal) error {
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
