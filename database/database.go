package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"main/models"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Database Dbinstance

func Connect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Database connected")
	db.Logger = db.Logger.LogMode(logger.Info)
	log.Println("Running migrations")

	db.AutoMigrate(&models.Task{})

	Database = Dbinstance{Db: db}
}
