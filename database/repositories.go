package database

import (
	"time"

	"gitnub.com/premwut/wallet-service/domain"
)

type UserRepository struct {
	// may have db instance here in real use case
}

func NewUserRepo() UserRepository {
	return UserRepository{}
}

var mockUsers = []domain.User{
	{
		ID:   "1",
		Name: "Jimi Hendrix",
		Age:  999,
	},
}

func (r UserRepository) GetUser(userId string) (domain.User, error) {
	var err error
	time.Sleep(1 * time.Second)
	mockUser := mockUsers[0]
	return mockUser, err
}
