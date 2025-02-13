package db

import (
	"log"
	"time"

	"github.com/Yonas21/go-habit-tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUrl := "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"

	var db *gorm.DB
	var err error

	// Retry connecting to the database
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Attempt %d: Failed to connect to database: %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Habit{})
	db.AutoMigrate(&models.HabitLogs{})
	db.AutoMigrate(&models.Streak{})
	db.AutoMigrate(&models.Rewards{})
	db.AutoMigrate(&models.Leaderboard{})

	return db
}
