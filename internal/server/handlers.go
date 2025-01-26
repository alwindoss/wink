package server

import (
	"net/http"

	"github.com/alwindoss/wink/internal/service"
)

type GoalsController interface {
	CreateGoals(w http.ResponseWriter, r *http.Request)
	FetchGoal(w http.ResponseWriter, r *http.Request)
	FetchGoals(w http.ResponseWriter, r *http.Request)
	UpdateGoals(w http.ResponseWriter, r *http.Request)
	DeleteGoals(w http.ResponseWriter, r *http.Request)
}

func NewGoalsController(svc service.GoalsService) GoalsController {
	return &goalsController{
		svc: svc,
	}
}

type goalsController struct {
	svc service.GoalsService
}

// CreateGoals implements GoalsController.
func (g *goalsController) CreateGoals(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// DeleteGoal implements GoalsController.
func (g *goalsController) DeleteGoals(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// FetchGoal implements GoalsController.
func (g *goalsController) FetchGoal(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// FetchGoals implements GoalsController.
func (g *goalsController) FetchGoals(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// UpdateGoal implements GoalsController.
func (g *goalsController) UpdateGoals(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
