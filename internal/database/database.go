package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
    log "github.com/sirupsen/logrus"
)

// Returns a pointer to a DB object
func NewDatabase() (*gorm.DB, error) {
    log.Info("Setting up new db connection")
    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbTable := os.Getenv("DB_TABLE")
    dbPort := os.Getenv("DB_PORT")
    sslMode := os.Getenv("DB_SSL_MODE")

    connectionString := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        dbHost, dbPort, dbUsername, dbPassword, dbTable, sslMode,
    )

    db, err := gorm.Open("postgres", connectionString)
    if err != nil {
        return db, err
    }
    // Ensure we can reach the DB :)
    if err := db.DB().Ping(); err != nil {
        return db, err
    }

    return db, nil
}
