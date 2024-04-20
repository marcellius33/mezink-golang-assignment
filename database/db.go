package database

import (
	"log"
	"mezink-golang-assignment/database/seed"
	"mezink-golang-assignment/models"
	"mezink-golang-assignment/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	// Connecting database
	db, err = gorm.Open(postgres.Open("host=db port=5432 user=postgres password=mezink dbname=mezink_golang_assignment sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database")
	}

	// Auto migrate models to database
	db.AutoMigrate(&models.Record{})
	log.Printf("Success connecting to database")
}

func GetDB() *gorm.DB {
	return db
}

func SeedDatabase() {
	recordRepository := repositories.NewRecordRepository(GetDB())
	recordSeed := seed.NewRecordSeeder(recordRepository)
	recordSeed.Execute()
}
