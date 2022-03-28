package routes

import (
	"github.com/labstack/echo/v4"
	"rentBook/delivery/handler/book"
	"rentBook/delivery/handler/rent"

	_userHandler "rentBook/delivery/handler/user"
	_middlewares "rentBook/delivery/middlewares"
)

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler, bh *book.BookHandler, rh *rent.RentHandler) {
	e.GET("/users", uh.GetAllHandler(), _middlewares.JWTMiddleware())
	e.GET("/users/profile", uh.GetUserById(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.DELETE("/users/:id", uh.DeleteUser(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdateUser(), _middlewares.JWTMiddleware())

	e.GET("/books", bh.GetAllHandler())
	e.GET("/books/:id", bh.GetBookById())
	e.POST("/books", bh.CreateBook(), _middlewares.JWTMiddleware())
	e.DELETE("/books/:id", bh.DeleteBook(), _middlewares.JWTMiddleware())
	e.PUT("/books/:id", bh.UpdateBook(), _middlewares.JWTMiddleware())

	e.GET("/rents", rh.GetAllHandler())
	e.POST("/rents", rh.CreateRentHandler(), _middlewares.JWTMiddleware())
	e.POST("/return", rh.ReturnBookHandler(), _middlewares.JWTMiddleware())

}
