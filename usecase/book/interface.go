package book

import _entities "rentBook/entities"

type BookUseCaseInterface interface {
	GetAll() ([]_entities.Book, error)
	GetBookById(id int) (_entities.Book, error)
	CreateBook(book _entities.Book) error
	DeleteBook(id int) error
	UpdateBook(id int, book _entities.Book) error
}
