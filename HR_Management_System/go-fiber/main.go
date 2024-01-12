package main

import (
	"fmt"
	"log"
	"os"

	controllers "github.com/AbhishekCS3459/HR_MS/controllers/Employee"
	"github.com/AbhishekCS3459/HR_MS/db"
	"github.com/AbhishekCS3459/HR_MS/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var mg = &models.Mg
var dbName, MONGOURI string

func init() {
	// logic to load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
		return
	}
	MONGOURI := os.Getenv("DB_URI")
	DB_NAME := os.Getenv(("DB_NAME"))
	// logic to connect to db
	db.ConnectDB(MONGOURI, DB_NAME)
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/employee", controllers.GetEmployeesAll)
	app.Post("/employee", controllers.CreateEmployee)
	app.Get("/employee/:id", controllers.GetEmployeeByID)
	app.Put("/employee/:id", controllers.CreateEmployee)
	app.Delete("/employee/:id", controllers.DeleteEmployee)
	fmt.Println("Server is running on port http://localhost:8080")
	app.Listen(":8080")
}
