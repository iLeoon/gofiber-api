package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	handlers "github.com/iLeon/gofiber-api/api/controllers"
	"github.com/iLeon/gofiber-api/helpers/customers"
)

func CustomersRoute(app *fiber.App, service customers.CustomerService, validator *validator.Validate) {
	app.Get("/customers", handlers.GetCustomers(service))
	app.Get("/customers/:id", handlers.GetCustomerById(service))
	app.Post("/customers/create", handlers.CreateCustomer(service, validator))
	app.Patch("/customers/update/:id", handlers.UpdateCustomer(service, validator))
	app.Delete("/customers/delete/:id", handlers.DeleteCustomer(service))
}
