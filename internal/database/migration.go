package database

import (
    "github.com/alexmk92/image-service/internal/image"
    "github.com/jinzhu/gorm"
)

// Given a gorm.Model, build the table.
func MigrateDB(db *gorm.DB) error {
    if result := db.AutoMigrate(&image.Image{}); result.Error != nil {
        return result.Error
    }

    return nil
}
