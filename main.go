package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/iLeon/gofiber-api/api/routes"
	"github.com/iLeon/gofiber-api/database"
	"github.com/iLeon/gofiber-api/helpers/customers"
)

func main() {
	db, _ := database.Connect()
	customerRepo := customers.NewRepo(db)
	customerServ := customers.NewService(customerRepo)

	app := fiber.New()
	routes.CustomersRoute(app, customerServ)
	app.Listen(":8080")

}
