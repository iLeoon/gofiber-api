package main

import (
	"log"
	"github.com/gofiber/fiber/v3"
	"github.com/iLeon/gofiber-api/api/routes"
	"github.com/iLeon/gofiber-api/database"
	"github.com/iLeon/gofiber-api/helpers/customers"
	"github.com/joho/godotenv"
)

func main() {
	db, _ := database.Connect()
	customerRepo := customers.NewRepo(db)
	customerServ := customers.NewService(customerRepo)
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Can't reload env variables")
	}

	app := fiber.New()
	routes.CustomersRoute(app, customerServ)
	app.Listen(":8080")

}
