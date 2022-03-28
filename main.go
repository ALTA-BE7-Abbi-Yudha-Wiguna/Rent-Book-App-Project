package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"rentBook/config"
	_userHandler "rentBook/delivery/handler/user"
	_userRepository "rentBook/repository/user"
	_userUseCase "rentBook/usecase/user"

	bookHandler "rentBook/delivery/handler/book"
	bookRepository "rentBook/repository/book"
	bookUseCase "rentBook/usecase/book"

	_middlewares "rentBook/delivery/middlewares"
	_routes "rentBook/delivery/routes"
	_utils "rentBook/utils"
)

func main() {
	configs := config.GetConfig()
	db := _utils.InitDB(configs)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	bookRepo := bookRepository.NewBookRepository(db)
	bookUse := bookUseCase.NewBookUseCase(bookRepo)
	bookHand := bookHandler.NewBookHandler(bookUse)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())

	_routes.RegisterPath(e, userHandler, bookHand)

	_routes.RegisterPath(e, userHandler, bookHand)
	log.Fatal(e.Start(fmt.Sprintf(":%d", configs.Port)))
}
