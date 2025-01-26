package repository

import "gorm.io/gorm"

type Goal struct {
}

type GoalsRepository interface {
	CreateGoal(g *Goal) error
	CreateGoals(g []*Goal) error
	UpdateGoal(g *Goal) error
	UpdateGoals(g []*Goal) error
	FetchGoal(id string) (*Goal, error)
	FetchGoals() ([]*Goal, error)
	DeleteGoal(g *Goal) error
}

func NewGormGoalsRepository(db *gorm.DB) GoalsRepository {
	r := &gormGoalsRepository{
		db: db,
	}

	return r
}

type gormGoalsRepository struct {
	db *gorm.DB
}

// CreateGoal implements GoalsRepository.
func (*gormGoalsRepository) CreateGoal(g *Goal) error {
	panic("unimplemented")
}

// CreateGoals implements GoalsRepository.
func (*gormGoalsRepository) CreateGoals(g []*Goal) error {
	panic("unimplemented")
}

// DeleteGoal implements GoalsRepository.
func (*gormGoalsRepository) DeleteGoal(g *Goal) error {
	panic("unimplemented")
}

// FetchGoal implements GoalsRepository.
func (g *gormGoalsRepository) FetchGoal(id string) (*Goal, error) {
	panic("unimplemented")
}

// FetchGoals implements GoalsRepository.
func (g *gormGoalsRepository) FetchGoals() ([]*Goal, error) {
	panic("unimplemented")
}

// UpdateGoal implements GoalsRepository.
func (*gormGoalsRepository) UpdateGoal(g *Goal) error {
	panic("unimplemented")
}

// UpdateGoals implements GoalsRepository.
func (*gormGoalsRepository) UpdateGoals(g []*Goal) error {
	panic("unimplemented")
}
