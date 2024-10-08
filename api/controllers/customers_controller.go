package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/iLeon/gofiber-api/helpers/customers"
	"github.com/iLeon/gofiber-api/models"
)

func GetCustomers(service customers.CustomerService) fiber.Handler {
	return func(c fiber.Ctx) error {
		customers, err := service.FetchCustomers()

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.CustomersErrorResponse(err, "Couldn't fetch the data"))
		}

		return c.JSON(models.CustomersSuccessResponse(customers))
	}

}

func GetCustomerById(service customers.CustomerService) fiber.Handler {
	return func(c fiber.Ctx) error {
		customerId := c.Params("id")

		customer, err := service.FetchCustomer(customerId)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(models.CustomersErrorResponse(err, "Couldn't fetch the customer data"))
		}

		return c.JSON(models.CustomerSuccessResponse(customer))
	}
}

func CreateCustomer(service customers.CustomerService, validator *validator.Validate) fiber.Handler {
	return func(c fiber.Ctx) error {
		var customer models.Customer

		err := c.Bind().Body(&customer)

		if err := validator.Struct(customer); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(err.Error())
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.CustomersErrorResponse(err, "Couldn't create a new customer"))
		}

		data, err := service.InsertCustomer(&customer)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.CustomersErrorResponse(err, "Couldn't create a new customer"))
		}

		return c.JSON(models.CustomerSuccessResponse(data))
	}
}

func UpdateCustomer(service customers.CustomerService, validator *validator.Validate) fiber.Handler {
	return func(c fiber.Ctx) error {
		customerId := c.Params("id")

		customer := &models.Customer{
			CustomerID: customerId,
		}

		err := c.Bind().Body(&customer)

		if err := validator.Struct(customer); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(err.Error())
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.CustomersErrorResponse(err, "Couldn't update the specified customer"))
		}

		data, err := service.ChangeCustomer(customer, customerId)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.CustomersErrorResponse(err, "Couldn't update the specified customer"))
		}

		return c.JSON(&fiber.Map{
			"data":    data,
			"message": "updated the specified customer successfully",
		})
	}
}

func DeleteCustomer(service customers.CustomerService) fiber.Handler {
	return func(c fiber.Ctx) error {
		customerId := c.Params("id")

		err := service.RemoveCustomer(customerId)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(models.CustomersErrorResponse(err, "Couldn't delete the specified customer"))
		}

		return c.JSON(&fiber.Map{
			"message": "deleted the specified customer successfully",
		})
	}
}
