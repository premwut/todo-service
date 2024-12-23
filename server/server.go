package server

import (
	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()

	c := NewUserController()
	e.GET("/hello", c.HelloWorld)
	e.GET("/user/:id", c.GetUser)

	// projects
	pc := NewProjectController()
	e.GET("/projects/:projectId", pc.getProject)
	e.Start(":8000")

}
