package repository

import (
	"context"
	"errors"

	"github.com/afthaab/job-portal/internal/models"
	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(ctx context.Context, UserDetails models.User) (models.User, error) {
	result := r.db.Create(&UserDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("could not create the user")
	}
	return UserDetails, nil
}

func (r *Repo) CheckEmail(ctx context.Context, email string) (models.User, error) {
	var userDetails models.User
	result := r.db.Where("email = ?", email).First(&userDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("email not found")
	}
	return userDetails, nil

}
