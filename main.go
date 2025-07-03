package main

import (
	"flag"
	"fmt"
	"github.com/ibrahimelothmani/API/internal/app"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the application on")
	flag.Parse()
	application, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health", HealthCheck)

	server := http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	application.Logger.Println("we are running our application")
	application.Logger.Printf("Starting server on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		application.Logger.Fatal("Error starting server:", err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
