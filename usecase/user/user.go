package user

import (
	_entities "rentBook/entities"
	_userRepository "rentBook/repository/user"
)

type UserUseCase struct {
	UserRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepo,
	}

}

func (uuc *UserUseCase) GetAll() ([]_entities.User, error) {
	users, err := uuc.UserRepository.GetAll()
	return users, err
}

func (uuc *UserUseCase) GetUserById(id int) (_entities.User, error) {
	user, err := uuc.UserRepository.GetUserById(id)
	return user, err
}

func (uuc *UserUseCase) CreateUser(user _entities.User) error {
	err := uuc.UserRepository.CreateUser(user)
	return err
}

func (uuc *UserUseCase) DeleteUser(id int) error {
	err := uuc.UserRepository.DeleteUser(id)
	return err
}

func (uuc *UserUseCase) UpdateUser(id int, user _entities.User) error {
	err := uuc.UserRepository.UpdateUser(id, user)
	return err
}
