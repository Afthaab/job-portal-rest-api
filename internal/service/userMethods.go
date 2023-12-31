package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/afthaab/job-portal/internal/models"
	"github.com/afthaab/job-portal/internal/pkg"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) UserSignIn(ctx context.Context, userData models.NewUser) (string, error) {
	// checcking the email in the db
	userDetails, err := s.UserRepo.CheckEmail(ctx, userData.Email)
	if err != nil {
		return "", err
	}

	// comaparing the password and hashed password
	err = pkg.CheckHashedPassword(userData.Password, userDetails.PasswordHash)
	if err != nil {
		log.Info().Err(err).Send()
		return "", errors.New("entered password is not wrong")
	}

	// setting up the claims
	claims := jwt.RegisteredClaims{
		Issuer:    "job portal project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token, err := s.auth.GenerateAuthToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *Service) UserSignup(ctx context.Context, userData models.NewUser) (models.User, error) {
	hashedPass, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return models.User{}, err
	}
	userDetails := models.User{
		Username:     userData.Username,
		Email:        userData.Email,
		PasswordHash: hashedPass,
	}

	userDetails, err = s.UserRepo.CreateUser(ctx, userDetails)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, nil
}
