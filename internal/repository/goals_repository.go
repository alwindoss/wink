package repository

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	Name         string
	Description  string
	Achieved     bool
	AchievedDate time.Time
	TargetDate   time.Time
}

type GoalsRepository interface {
	CreateGoal(g *Goal) (uint, error)
	CreateGoals(g []*Goal) (uint, error)
	UpdateGoal(g *Goal) error
	UpdateGoals(g []*Goal) error
	FetchGoal(id string) (*Goal, error)
	FetchGoals() ([]*Goal, error)
	DeleteGoal(g *Goal) error
}

func NewGormGoalsRepository(db *gorm.DB) GoalsRepository {
	db.AutoMigrate(&Goal{})
	r := &gormGoalsRepository{
		db: db,
	}

	return r
}

type gormGoalsRepository struct {
	db *gorm.DB
}

// CreateGoal implements GoalsRepository.
func (r *gormGoalsRepository) CreateGoal(g *Goal) (uint, error) {
	tx := r.db.Create(g)
	if tx.Error != nil {
		err := fmt.Errorf("unable to create the goal:s %w", tx.Error)
		return 0, err
	}
	log.Printf("Created goal with ID %d", g.ID)
	return g.ID, nil
}

// CreateGoals implements GoalsRepository.
func (*gormGoalsRepository) CreateGoals(g []*Goal) (uint, error) {
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
