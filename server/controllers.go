package server

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitnub.com/premwut/todo-service/database"
	"gitnub.com/premwut/todo-service/domain"
	"gitnub.com/premwut/todo-service/usecase"
)

var InternalServerError = errors.New("internal server error")
var NotFoundError = errors.New("resource not found")

type UserController struct {
	userService usecase.UserService
}

func NewUserController() UserController {
	r := database.NewUserRepo()
	service := usecase.NewUserService(r)
	return UserController{
		userService: *service,
	}
}

func (controller *UserController) HelloWorld(c echo.Context) error {
	msg := domain.Message{Text: "Hello world from wallet service!"}
	return c.JSON(http.StatusOK, msg)
}

func (controller *UserController) GetUser(c echo.Context) error {
	var err error
	userId := c.Param("id")

	user, err := controller.userService.GetUser(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, InternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}

type ProjectController struct {
	projectService usecase.ProjectService
}

func NewProjectController() ProjectController {
	r := database.NewProjectRepository(database.Database{})
	service := usecase.NewProjectService(r)
	return ProjectController{
		projectService: *service,
	}
}

func (controller *ProjectController) getProject(c echo.Context) error {
	var err error
	id := c.Param("projectId")

	project, err := controller.projectService.GetProject(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, NotFoundError)
	}

	return c.JSON(http.StatusOK, project)
}
