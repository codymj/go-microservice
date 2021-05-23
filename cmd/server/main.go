package main

import "fmt"

// App contains things for the app to work
type App struct {}

// Run initializes the application
func (app *App) Run() error {
	fmt.Println("Initializing application...")
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting application")
		fmt.Println(err)
	}
}
