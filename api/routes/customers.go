package routes

import (
	"github.com/gofiber/fiber/v3"
	handlers "github.com/iLeon/gofiber-api/api/controllers"
	"github.com/iLeon/gofiber-api/helpers/customers"
)

func CustomersRoute(app *fiber.App, service customers.CustomerService) {
	app.Get("/customers", handlers.GetCustomers(service))
	app.Get("/customers/:id", handlers.GetCustomerById(service))
	app.Post("/customers/create", handlers.CreateCustomer(service))
	app.Patch("/customers/update/:id", handlers.UpdateCustomer(service))
	app.Delete("/customers/delete/:id", handlers.DeleteCustomer(service))
}
