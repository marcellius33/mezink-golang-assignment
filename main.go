package main

import (
	"fmt"
	"mezink-golang-assignment/controllers"
	"mezink-golang-assignment/database"
	"mezink-golang-assignment/params/custom"
	"mezink-golang-assignment/repositories"
	"mezink-golang-assignment/routers"
	"mezink-golang-assignment/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Initialize database connection
func init() {
	database.Connect()
}

func customValidation() {
	validate := validator.New()

	if err := validate.RegisterValidation("customDate", custom.ValidateCustomDate); err != nil {
		fmt.Println("Failed to register custom date validation", err)
		return
	}
}

func main() {
	database.SeedDatabase()
	// customValidation()

	routes := gin.Default()

	recordRepository := repositories.NewRecordRepository(database.GetDB())
	recordService := services.NewRecordService(recordRepository)
	recordController := controllers.NewRecordController(recordService)
	routers.InitRecordRoutes(routes, recordController)

	routes.Run(":8080")
}
