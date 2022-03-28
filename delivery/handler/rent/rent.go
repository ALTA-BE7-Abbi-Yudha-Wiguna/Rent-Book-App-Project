package rent

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rentBook/delivery/helper"
	"rentBook/delivery/middlewares"
	"rentBook/entities"
	_rentUseCase "rentBook/usecase/rent"
)

type RentHandler struct {
	rentUseCase _rentUseCase.RentUseCaseInterface
}

func NewRentHandler(rentUseCase _rentUseCase.RentUseCaseInterface) *RentHandler {
	return &RentHandler{
		rentUseCase: rentUseCase,
	}
}

func (rh *RentHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		rents, err := rh.rentUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all rents", rents))
	}
}

func (rh *RentHandler) CreateRentHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var rent entities.Rent

		c.Bind(&rent)
		id := middlewares.ExtractToken(c)
		rent.UserID = uint(id)
		err := rh.rentUseCase.CreateRent(rent)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to rent book"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create rents", rent))
	}
}
func (rh *RentHandler) ReturnBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var rent entities.Rent

		c.Bind(&rent)
		id := middlewares.ExtractToken(c)
		rent.UserID = uint(id)
		err := rh.rentUseCase.ReturnBook(rent)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to return book"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success return book", rent))
	}
}
