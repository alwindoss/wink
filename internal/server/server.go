package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alwindoss/wink/internal/controllers"
	"github.com/alwindoss/wink/internal/repository"
	"github.com/alwindoss/wink/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Addr string
	DSN  string
}

func Run(cfg *Config) error {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return err
	}
	goalsRepo := repository.NewGormGoalsRepository(db)
	goalsSvc := service.NewGoalsService(goalsRepo)
	goalsCtrl := controllers.NewGoalsController(goalsSvc)

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	setupGoalsRoutes(r, goalsCtrl)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Pong @ %s", time.Now().String())
	})

	log.Printf("Listening on %s", cfg.Addr)
	err = http.ListenAndServe(cfg.Addr, r)
	if err != nil {
		return err
	}
	return nil
}

func setupGoalsRoutes(m *chi.Mux, ctrl controllers.GoalsController) {
	m.Route("/wink/v1", func(r chi.Router) {
		r.Post("/goals", ctrl.CreateGoals)
		r.Get("/goals", ctrl.FetchGoals)
		r.Get("/goals/{goal_id}", ctrl.FetchGoal)
		r.Patch("/goals/{goal_id}", ctrl.UpdateGoals)
		r.Put("/goals/{goal_id}", ctrl.UpdateGoals)
		r.Delete("/goals/{goal_id}", ctrl.DeleteGoals)
	})
}
