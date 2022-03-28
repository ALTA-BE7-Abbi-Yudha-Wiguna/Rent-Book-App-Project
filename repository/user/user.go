package user

import (
	"fmt"
	"gorm.io/gorm"
	_entities "rentBook/entities"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetAll() ([]_entities.User, error) {
	var users []_entities.User
	tx := ur.database.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil

}

func (ur *UserRepository) GetUserById(id int) (_entities.User, error) {
	var users _entities.User
	tx := ur.database.Where("id = ?", id).First(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil

}

func (ur *UserRepository) CreateUser(user _entities.User) error {

	tx := ur.database.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (ur *UserRepository) DeleteUser(id int) error {
	var users _entities.User
	tx := ur.database.Where("id = ?", id).Delete(&users)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (ur *UserRepository) UpdateUser(id int, user _entities.User) error {

	tx := ur.database.Where("id = ?", id).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}
