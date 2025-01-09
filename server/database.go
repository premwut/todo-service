package server

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=<YOUR_PASSWORD> database= todo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.New("error while connecting to database")
	}

	return db, err
}
