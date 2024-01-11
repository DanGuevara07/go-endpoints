package server

import (
	"go-endpoints/internal/users"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Server *echo.Echo

func init() {
	Server = echo.New()

	Server.Use(middleware.Logger())
	prefixGroup := Server.Group(os.Getenv("GLOBAL_PREFIX"))

	prefixGroup.GET("/users", users.GetUsers)
	prefixGroup.GET("/users/:id", users.GetUser)
	prefixGroup.POST("/users", users.CreateUser)
	prefixGroup.PATCH("/users/:id", users.UpdateUser)
	prefixGroup.DELETE("/users/:id", users.DeleteUser)
}
