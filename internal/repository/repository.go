package repository

import (
	"errors"

	"github.com/afthaab/job-portal/internal/models"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type UserRepo interface {
	CreateUser(userData models.User) (models.User, error)
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
