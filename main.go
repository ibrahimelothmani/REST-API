package main

import (
	"flag"
	"fmt"
	"github.com/ibrahimelothmani/API/internal/app"
	"github.com/ibrahimelothmani/API/internal/routes"
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

	defer application.DB.Close()

	r := routes.SetupRoutes(application)

	server := http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           r,
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
