package service

import (
	"context"
	"errors"

	"github.com/afthaab/job-portal/internal/models"
	"github.com/afthaab/job-portal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
}

type UserService interface {
	UserSignup(ctx context.Context, userData models.NewUser) (models.User, error)
}

func NewService(userRepo repository.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		UserRepo: userRepo,
	}, nil
}
