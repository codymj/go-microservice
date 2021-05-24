package main

import (
	"codymj/go-microservice/internal/database"
	transportHTTP "codymj/go-microservice/internal/transport/http"
	"log"
	"net/http"
)

// App contains things for the app to work
type App struct {}

// Run initializes the application
func (app *App) Run() error {
	log.Println("Initializing application...")

	_, err := database.Init()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	err = http.ListenAndServe(":8080", handler.Router)
	if err != nil {
		log.Println("Error setting up server")
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		log.Println("Error starting application")
		log.Println(err)
	}
}
