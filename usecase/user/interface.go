package user

import _entities "rentBook/entities"

type UserUseCaseInterface interface {
	GetAll() ([]_entities.User, error)
	GetUserById(id int) (_entities.User, error)
	CreateUser(user _entities.User) error
	DeleteUser(id int) error
	UpdateUser(id int, user _entities.User) error
}
