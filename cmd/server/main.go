package main

import (
	transportHTTP "codymj/go-microservice/internal/transport/http"
	"fmt"
	"net/http"
)

// App contains things for the app to work
type App struct {}

// Run initializes the application
func (app *App) Run() error {
	fmt.Println("Initializing application...")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Error setting up server")
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting application")
		fmt.Println(err)
	}
}
