package server

import (
	"net/http"
	"os"
	"time"

	"github.com/alwindoss/wink/internal/repository"
	"github.com/alwindoss/wink/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Addr string
}

func Run(cfg *Config) error {
	os.Getenv("DB_URL")
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	goalsRepo := repository.NewGormGoalsRepository(db)
	goalsSvc := service.NewGoalsService(goalsRepo)

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

	setupGoalsRoutes(r, goalsSvc)

	err = http.ListenAndServe(cfg.Addr, r)
	if err != nil {
		return err
	}
	return nil
}

func setupGoalsRoutes(m *chi.Mux, svc service.GoalsService) {
	m.Route("/wink/api/v1", func(r chi.Router) {
		ctrl := NewGoalsController(svc)
		r.Post("/goals", ctrl.CreateGoals)
		r.Get("/goals", ctrl.FetchGoals)
		r.Get("/goals/{goal_id}", ctrl.FetchGoal)
		r.Patch("/goals/{goal_id}", ctrl.UpdateGoals)
		r.Put("/goals/{goal_id}", ctrl.UpdateGoals)
		r.Delete("/goals/{goal_id}", ctrl.DeleteGoals)
	})
}
