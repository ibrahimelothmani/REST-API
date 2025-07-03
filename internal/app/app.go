package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"github.com/ibrahimelothmani/API/internal/api"
	"github.com/ibrahimelothmani/API/internal/store"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {

	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	workoutHandler := api.NewWorkoutHandler()
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}
	return app, nil
}
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
