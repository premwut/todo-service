package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitnub.com/premwut/todo-service/domain"
	"gitnub.com/premwut/todo-service/model"
	"gitnub.com/premwut/todo-service/usecase"
	"gorm.io/gorm"
)

var InternalServerError = errors.New("internal server error")
var NotFoundError = errors.New("resource not found")

type UserController struct {
	userService usecase.UserService
}

func NewUserController() UserController {
	r := model.NewUserRepo()
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

func NewProjectController(db *gorm.DB) ProjectController {
	service := usecase.NewProjectService(db)
	return ProjectController{
		projectService: *service,
	}
}

func (controller *ProjectController) createProject(c echo.Context) error {
	var err error
	// TODO: refactor jsonBody logic to a helper
	jsonBody := make(map[string]string)
	err = json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	projectName := jsonBody["name"]
	fmt.Printf("[pc.createProject] projectName: %v", projectName)

	project, err := controller.projectService.CreateProject(projectName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, project)
}

func (controller *ProjectController) getProject(c echo.Context) error {
	var err error
	id := c.Param("projectId")

	project, err := controller.projectService.GetProject(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, project)
}

type CreateTaskRequestParams struct {
	projectId string
}

type CreateTaskRequestBody struct {
	taskName string
}

// * route: POST /projects/:projectId/tasks
func (controller *ProjectController) createTask(c echo.Context) error {
	var err error
	projectId := c.Param("projectId")
	jsonBody := make(map[string]string)
	err = json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	taskName := jsonBody["taskName"]
	task, err := controller.projectService.CreateTask(projectId, taskName)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, task)
}
