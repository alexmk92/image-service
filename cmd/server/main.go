package main

import (
	"net/http"

	"github.com/alexmk92/image-service/internal/database"
    "github.com/alexmk92/image-service/internal/image"
	transportHTTP "github.com/alexmk92/image-service/internal/transport/http"

    log "github.com/sirupsen/logrus"
)

// App - contain application information
type App struct {
    Name string
    Version string
}

func (app *App) Run() error {
    log.SetFormatter(&log.JSONFormatter{})
    log.WithFields(log.
            Fields{"AppName": app.Name, "AppVersion": app.Version}).
            Info("Setting up Application",
    )

    var err error
    db, err := database.NewDatabase()
    if err != nil {
        return err
    }

    database.MigrateDB(db)

    imageService := image.NewService(db)

    handler := transportHTTP.NewHandler(imageService)
    handler.SetupRoutes()

    if err := http.ListenAndServe(":8080", handler.Router); err != nil {
        log.Error("Failed to set up server")
        return err
    }
    return nil
}

func main() {
    app := App{
        Name: "Image Service",
        Version: "1.0.0",
    }
    if err := app.Run(); err != nil {
        log.Error("Error starting up our REST API")
        log.Fatal(err)
    }
}
