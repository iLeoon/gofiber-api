package customers

import "github.com/iLeon/gofiber-api/models"

type CustomerService interface {
	FetchCustomers() (*[]models.Customer, error)
	FetchCustomer(id string) (*models.Customer, error)
	InsertCustomer(body *models.Customer) (*models.Customer, error)
	ChangeCustomer(body *models.Customer, customerId string) (*models.Customer, error)
	RemoveCustomer(customerId string) error
}

type Service struct {
	repository CustomersRepository
}

func NewService(r CustomersRepository) CustomerService {
	return &Service{
		repository: r,
	}
}

func (s *Service) FetchCustomers() (*[]models.Customer, error) {
	return s.repository.FindAll()
}

func (s *Service) FetchCustomer(id string) (*models.Customer, error) {
	return s.repository.FindOne(id)
}

func (s *Service) InsertCustomer(body *models.Customer) (*models.Customer, error) {
	return s.repository.Create(body)
}

func (s *Service) ChangeCustomer(body *models.Customer, customerId string) (*models.Customer, error) {
	return s.repository.Update(body, customerId)
}

func (s *Service) RemoveCustomer(customerId string) error {
	return s.repository.Delete(customerId)
}
