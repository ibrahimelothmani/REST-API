package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type WorkoutHandler struct {
}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkID := chi.URLParam(r, "workID")
	if paramsWorkID == "" {
		http.NotFound(w, r)
		return
	}

	workID, err := strconv.ParseInt(paramsWorkID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Work ID: %d\n", workID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	paramsWorkID := chi.URLParam(r, "workID")
	if paramsWorkID == "" {
		http.NotFound(w, r)
		return
	}
	workID, err := strconv.ParseInt(paramsWorkID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Work ID: %d\n", workID)
}
