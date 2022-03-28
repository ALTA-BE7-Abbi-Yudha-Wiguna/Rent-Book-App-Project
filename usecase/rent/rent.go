package rent

import (
	"errors"
	_entities "rentBook/entities"
	bookRepository "rentBook/repository/book"
	_rentRepository "rentBook/repository/rent"
)

type RentUseCase struct {
	RentRepository _rentRepository.RentRepositoryInterface
	BookRepository bookRepository.BookRepositoryInterface
}

func NewRentUseCase(rentRepo _rentRepository.RentRepositoryInterface, bookRepo bookRepository.BookRepositoryInterface) RentUseCaseInterface {
	return &RentUseCase{
		RentRepository: rentRepo,
		BookRepository: bookRepo,
	}

}
func (ruc *RentUseCase) GetAll() ([]_entities.Rent, error) {
	rents, err := ruc.RentRepository.GetAll()
	return rents, err
}

func (ruc *RentUseCase) CreateRent(rent _entities.Rent) error {

	book, err := ruc.BookRepository.GetBookById(int(rent.BookID))
	if book.Status != "Available" {
		return errors.New("Buku Tidak Bisa Disewa")
	}

	book.Status = "Rent"
	ruc.BookRepository.UpdateBook(int(rent.BookID), book)

	err = ruc.RentRepository.CreateRent(rent)
	return err
}

func (ruc *RentUseCase) ReturnBook(rent _entities.Rent) error {

	var book _entities.Book

	rentbook, err := ruc.RentRepository.GetRentById(int(rent.Id))

	if rentbook.StatusRent == "Sudah Dikembalikan" || rentbook.UserID != rent.UserID {
		return errors.New("Buku Tidak Bisa Disewa")
	}

	rentbook.StatusRent = "Sudah Dikembalikan"
	ruc.RentRepository.ReturnBook(rentbook)

	book.Status = "Available"
	ruc.BookRepository.UpdateBook(int(rent.BookID), book)

	return err
}
