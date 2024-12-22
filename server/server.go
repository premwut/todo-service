package server

import (
	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()

	c := NewUserController()
	e.GET("/hello", c.HelloWorld)
	e.GET("/user/:id", c.GetUser)
	e.Start(":8000")
}
