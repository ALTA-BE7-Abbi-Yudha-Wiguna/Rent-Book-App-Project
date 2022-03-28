package book

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rentBook/delivery/helper"
	"rentBook/delivery/middlewares"
	_entities "rentBook/entities"
	_bookUseCase "rentBook/usecase/book"
	"strconv"
)

type BookHandler struct {
	bookUseCase _bookUseCase.BookUseCaseInterface
}

func NewBookHandler(bookUseCase _bookUseCase.BookUseCaseInterface) *BookHandler {
	return &BookHandler{
		bookUseCase: bookUseCase,
	}
}

func (bh *BookHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := bh.bookUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all books", books))
	}
}

func (bh *BookHandler) GetBookById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))

		books, err := bh.bookUseCase.GetBookById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get book by id", books))
	}
}

func (bh *BookHandler) CreateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var book _entities.Book
		c.Bind(&book)
		id := middlewares.ExtractToken(c)
		book.UserID = uint(id)
		err := bh.bookUseCase.CreateBook(book)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create book"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success Create books"))
	}
}

func (bh *BookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))

		err := bh.bookUseCase.DeleteBook(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete book by id"))
	}
}

func (bh *BookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var books _entities.Book
		c.Bind(&books)

		err := bh.bookUseCase.UpdateBook(id, books)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update book by id", books))
	}
}
