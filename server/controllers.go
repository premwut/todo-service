package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitnub.com/premwut/wallet-service/database"
	"gitnub.com/premwut/wallet-service/domain"
	"gitnub.com/premwut/wallet-service/usecase"
)

var InternalServerError = errors.New("Internal server error")

type UserController struct {
	userService usecase.UserService
}

func NewUserController() UserController {
	r := database.NewUserRepo()
	service := usecase.NewUserService(r)
	fmt.Println("[NewUserController] service:", service)
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
