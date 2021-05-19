package main

import (
	"fmt"
	"net/http"

	"github.com/alexmk92/image-service/internal/database"
    "github.com/alexmk92/image-service/internal/image"
	transportHTTP "github.com/alexmk92/image-service/internal/transport/http"
)

// App - pointers to db connections etc
type App struct {

}

func (app *App) Run() error {
    fmt.Println("Setting up App")
    var err error
    db, err := database.NewDatabase()
    if err != nil {
        return err
    }

    database.MigrateDB(db)

    imageService := image.NewService(db)

    handler := transportHTTP.NewHandler(imageService)
    handler.SetupRoutes()

    if err := http.ListenAndServe(":8089", handler.Router); err != nil {
        fmt.Println("Failed to set up server")
        return err
    }
    return nil
}

func main() {
    fmt.Println("Image service")
    app := App{}
    if err := app.Run(); err != nil {
        fmt.Println("Error starting up our REST API")
        fmt.Println(err)
    }
}
