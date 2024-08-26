package customers

import (
	"database/sql"
	"fmt"
	"github.com/iLeon/gofiber-api/models"
)

type CustomersRepository interface {
	FindAll() (*[]models.Customer, error)
	FindOne(id string) (*models.Customer, error)
	Create(body *models.Customer) (*models.Customer, error)
	Update(body *models.Customer, customerId string) (*models.Customer, error)
	Delete(customerId string) error
}

type Repository struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) CustomersRepository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindAll() (*[]models.Customer, error) {
	var customers []models.Customer
	rows, err := r.DB.Query("select customer_id, contact_name, city, country from customers")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		customer := models.Customer{}
		err := rows.Scan(&customer.CustomerID, &customer.ContactName, &customer.City, &customer.Country)
		fmt.Println(err)
		customers = append(customers, customer)
	}
	return &customers, nil

}

func (r *Repository) FindOne(customerId string) (*models.Customer, error) {
	customer := models.Customer{}

	data := r.DB.QueryRow("select customer_id, contact_name, city, country from customers where customer_id = $1", customerId)

	err := data.Scan(
		&customer.CustomerID,
		&customer.ContactName,
		&customer.City,
		&customer.Country,
	)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *Repository) Create(body *models.Customer) (*models.Customer, error) {

	_, err := r.DB.Query("insert into customers (customer_id, contact_name, city, country) values ($1, $2, $3, $4)",
		body.CustomerID,
		body.ContactName,
		body.City,
		body.Country,
	)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *Repository) Update(body *models.Customer, customerId string) (*models.Customer, error) {
	_, err := r.DB.Query("update customers set contact_name=$1, city=$2, country=$3 where customer_id=$4",
		body.ContactName,
		body.City,
		body.Country,
		customerId,
	)
	fmt.Println(customerId)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (r *Repository) Delete(customerId string) error {
	_, err := r.DB.Query("delete from customers where customer_id = $1", customerId)

	if err != nil {
		return err
	}

	return nil
}
