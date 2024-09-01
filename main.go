package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/iLeon/gofiber-api/api/routes"
	"github.com/iLeon/gofiber-api/database"
	"github.com/iLeon/gofiber-api/helpers/customers"
)

func main() {
	db, _ := database.Connect()
	customerRepo := customers.NewRepo(db)
	customerServ := customers.NewService(customerRepo)
	validator := validator.New()

	app := fiber.New()
	routes.CustomersRoute(app, customerServ, validator)
	app.Listen(":8080")

}
