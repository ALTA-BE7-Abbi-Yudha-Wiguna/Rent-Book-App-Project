package rent

import _entities "rentBook/entities"

type RentRepositoryInterface interface {
	GetAll() ([]_entities.Rent, error)
	CreateRent(rent _entities.Rent) error
	ReturnBook(rent _entities.Rent) error
	GetRentById(id int) (_entities.Rent, error)
}
