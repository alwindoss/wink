package controllers

import (
	"encoding/json"
	"net/http"
	"time"

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

type errorResp struct {
	ErrCode    string `json:"err_code,omitempty"`
	ErrMessage string `json:"err_message,omitempty"`
}

type createGoalsRequest struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	TargetDate   time.Time `json:"target_date"`
	Achieved     bool      `json:"achieved"`
	AchievedDate time.Time `json:"achieved_date"`
}

type createGoalsResponse struct {
	ID uint `json:"id,omitempty"`
	// Name         string    `json:"name,omitempty"`
	// Description  string    `json:"description,omitempty"`
	// TargetDate   time.Time `json:"target_date,omitempty"`
	// Achieved     bool      `json:"achieved,omitempty"`
	// AchievedDate time.Time `json:"achieved_date,omitempty"`
}

// CreateGoals implements GoalsController.
func (g *goalsController) CreateGoals(w http.ResponseWriter, r *http.Request) {
	req := createGoalsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ERR00001", "Error decoding the request")
		return
	}
	defer r.Body.Close()
	goal := service.Goal{
		Name:         req.Name,
		Description:  req.Description,
		Achieved:     req.Achieved,
		AchievedDate: req.AchievedDate,
		TargetDate:   req.TargetDate,
	}
	id, err := g.svc.CreateGoal(&goal)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "ERR00002", "Error in creating the goal")
		return
	}
	resp := createGoalsResponse{
		ID: id,
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	return
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
