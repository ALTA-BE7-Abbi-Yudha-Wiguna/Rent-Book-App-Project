package rent

import _entities "rentBook/entities"

type RentUseCaseInterface interface {
	GetAll() ([]_entities.Rent, error)
	CreateRent(rent _entities.Rent) error
	ReturnBook(rent _entities.Rent) error
}
