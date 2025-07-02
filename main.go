package main

import (
	"github.com/ibrahimelothmani/API/internal/app"
	"net/http"
	"time"
)

func main() {

	application, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	application.Logger.Println("we are running our application")

	http.HandleFunc("/health", HealthCheck)

	server := http.Server{
		Addr:              ":8080",
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		application.Logger.Fatal("Error starting server:", err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
