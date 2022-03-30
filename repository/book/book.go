package book

import (
	"fmt"
	"gorm.io/gorm"
	_entities "rentBook/entities"
)

type BookRepository struct {
	database *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		database: db,
	}
}

func (br *BookRepository) GetAll() ([]_entities.Book, error) {
	var books []_entities.Book
	tx := br.database.Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return books, nil

}

func (br *BookRepository) GetBookById(id int) (_entities.Book, error) {
	var books _entities.Book
	tx := br.database.Where("id = ?", id).First(&books)
	if tx.Error != nil {
		return books, tx.Error
	}
	return books, nil

}

func (br *BookRepository) CreateBook(book _entities.Book) error {

	tx := br.database.Save(&book)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (br *BookRepository) DeleteBook(id, userID int) error {
	var books _entities.Book
	err := br.database.Where("id =? and user_id = ?", id, userID).First(&books).Error
	if err != nil {
		return err
	}
	br.database.Delete(&books)
	
	return nil

}

func (br *BookRepository) UpdateBook(id int, book _entities.Book) error {

	tx := br.database.Where("id = ?", id).Updates(&book)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}
