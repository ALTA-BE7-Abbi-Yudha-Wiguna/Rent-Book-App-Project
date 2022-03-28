package routes

import (
	"github.com/labstack/echo/v4"
	_userHandler "rentBook/delivery/handler/user"
	_middlewares "rentBook/delivery/middlewares"
)

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.GET("/users", uh.GetAllHandler(), _middlewares.JWTMiddleware())
	e.GET("/users/profile", uh.GetUserById(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.DELETE("/users/:id", uh.DeleteUser(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdateUser(), _middlewares.JWTMiddleware())

}