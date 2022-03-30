package rent

import (
	"gorm.io/gorm"
	_entities "rentBook/entities"
)

type RentRepository struct {
	database *gorm.DB
}

func NewRentRepository(db *gorm.DB) *RentRepository {
	return &RentRepository{
		database: db,
	}
}
func (rr *RentRepository) GetAll() ([]_entities.Rent, error) {
	var rents []_entities.Rent
	tx := rr.database.Find(&rents)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rents, nil

}

func (rr *RentRepository) CreateRent(rent _entities.Rent) error {
	tx := rr.database.Create(&rent)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (rr *RentRepository) ReturnBook(rent _entities.Rent) error {

	tx := rr.database.Where("id = ?", rent.Id).Updates(&rent)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (rr *RentRepository) GetRentById(id int) (_entities.Rent, error) {
	var rents _entities.Rent
	tx := rr.database.Where("id = ?", id).First(&rents)
	if tx.Error != nil {
		return rents, tx.Error
	}
	return rents, nil

}
