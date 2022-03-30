package book

import (
	_entities "rentBook/entities"
	_bookRepository "rentBook/repository/book"
)

type BookUseCase struct {
	BookRepository _bookRepository.BookRepositoryInterface
}

func NewBookUseCase(bookRepo _bookRepository.BookRepositoryInterface) BookUseCaseInterface {
	return &BookUseCase{
		BookRepository: bookRepo,
	}

}

func (buc *BookUseCase) GetAll() ([]_entities.Book, error) {
	books, err := buc.BookRepository.GetAll()
	return books, err
}

func (buc *BookUseCase) GetBookById(id int) (_entities.Book, error) {
	book, err := buc.BookRepository.GetBookById(id)
	return book, err
}

func (buc *BookUseCase) CreateBook(book _entities.Book) error {
	err := buc.BookRepository.CreateBook(book)
	return err
}

func (buc *BookUseCase) DeleteBook(id, userID int) error {
	err := buc.BookRepository.DeleteBook(id, userID)
	return err
}

func (buc *BookUseCase) UpdateBook(id int, book _entities.Book) error {
	err := buc.BookRepository.UpdateBook(id, book)
	return err
}
