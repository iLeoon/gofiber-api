package models

import "github.com/gofiber/fiber/v3"

type Customer struct {
	CustomerID  string `json:"customer_id" validate:"required,alpha,omitempty,max=4,min=4"`
	ContactName string `json:"contact_name" validate:"required"`
	City        string `json:"city" validate:"required"`
	Country     string `json:"country" validate:"required"`
}

func CustomersSuccessResponse(data *[]Customer) *fiber.Map {
	return &fiber.Map{
		"customers": data,
	}
}

func CustomerSuccessResponse(data *Customer) *fiber.Map {
	customer := Customer{
		CustomerID:  data.CustomerID,
		ContactName: data.ContactName,
		City:        data.City,
		Country:     data.Country,
	}

	return &fiber.Map{
		"customer": customer,
	}
}

func CustomersErrorResponse(err error, msg string) *fiber.Map {
	return &fiber.Map{
		"error":   err.Error(),
		"message": msg,
	}
}
