package usecase

import (
	// "errors"

	"errors"
	"fmt"

	"gitnub.com/premwut/wallet-service/domain"
)

type UserRepositoryInterface interface {
	GetUser(userId string) (domain.User, error)
}

type UserService struct {
	userRepository UserRepositoryInterface
}

func NewUserService(r UserRepositoryInterface) *UserService {
	s := UserService{}
	fmt.Println("[NewUserService] r:", r)
	s.userRepository = r
	return &s
}

func (userService *UserService) GetUser(userId string) (*domain.User, error) {
	userService.TestHelloFromUserService()
	user, err := userService.userRepository.GetUser(userId)

	if err != nil {
		return nil, errors.New("Internal server error")
	}

	return &user, nil
}

func (userService *UserService) TestHelloFromUserService() {
	fmt.Println("TestHelloFromUserService")
}
