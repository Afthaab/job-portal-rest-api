package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	auth "github.com/afthaab/job-portal/internal/auth/mockModels"
	"github.com/afthaab/job-portal/internal/models"
	"github.com/afthaab/job-portal/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestService_UserSignIn(t *testing.T) {
	type args struct {
		ctx      context.Context
		userData models.NewUser
	}
	tests := []struct {
		name             string
		args             args
		want             string
		wantErr          bool
		claims           jwt.RegisteredClaims
		mockResponse     func() (models.User, error)
		mockAuthResponse func() (string, error)
	}{

		{
			name: "wrong email",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "afthab",
					Email:    "afthab606@gmail.com",
					Password: "12345",
				},
			},
			want:    "",
			wantErr: true,
			mockResponse: func() (models.User, error) {
				return models.User{}, errors.New("test error from the mock function")
			},
			mockAuthResponse: func() (string, error) {
				return "", errors.New("test error from the mock function")
			},
		},
		{
			name: "token generation failed",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "afthab",
					Email:    "afthab606@gmail.com",
					Password: "12345678",
				},
			},
			want:    "jwt test string",
			wantErr: false,
			mockResponse: func() (models.User, error) {
				return models.User{
					Username:     "afthab",
					Email:        "afthab606@gmail.com",
					PasswordHash: "$2a$10$uS/GmX48bxvhGPS.IrujaefuktoqGuKz3HBeOOMH6MGrnDT1H4TEy",
					Model: gorm.Model{
						ID: 1,
					},
				}, nil
			},
			claims: jwt.RegisteredClaims{
				Issuer:  "job portal project",
				Subject: "1",
				Audience: jwt.ClaimStrings{
					"users",
				},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
			mockAuthResponse: func() (string, error) {
				return "jwt test string", nil
			},
		},
		{
			name: "success generate token",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "afthab",
					Email:    "afthab606@gmail.com",
					Password: "12345678",
				},
			},
			want:    "",
			wantErr: true,
			mockResponse: func() (models.User, error) {
				return models.User{
					Username:     "afthab",
					Email:        "afthab606@gmail.com",
					PasswordHash: "$2a$10$uS/GmX48bxvhGPS.IrujaefuktoqGuKz3HBeOOMH6MGrnDT1H4TEy",
					Model: gorm.Model{
						ID: 1,
					},
				}, nil
			},
			claims: jwt.RegisteredClaims{
				Issuer:  "job portal project",
				Subject: "1",
				Audience: jwt.ClaimStrings{
					"users",
				},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
			mockAuthResponse: func() (string, error) {
				return "", errors.New("test error from mock function")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockAuth := auth.NewMockAuthentication(mc)

			mockRepo.EXPECT().CheckEmail(tt.args.ctx, tt.args.userData.Email).Return(tt.mockResponse()).AnyTimes()

			mockAuth.EXPECT().GenerateAuthToken(tt.claims).Return(tt.mockAuthResponse()).AnyTimes()

			svc, err := NewService(mockRepo, mockAuth)
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}
			got, err := svc.UserSignIn(tt.args.ctx, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Service.UserSignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UserSignup(t *testing.T) {

	type args struct {
		ctx      context.Context
		userData models.NewUser
	}
	tests := []struct {
		name         string
		args         args
		want         models.User
		wantErr      bool
		mockResponse func() (models.User, error)
	}{
		{
			name: "error from the database",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "afthab",
					Email:    "afthab606@gmail.com",
					Password: "12345678",
				},
			},
			want:    models.User{}, // Change the expected result to an empty User since an error is expected.
			wantErr: true,          // Set wantErr to true since an error is expected.
			mockResponse: func() (models.User, error) {
				return models.User{}, errors.New("error while hashing the password")
			},
		},
		{
			name: "success from the database",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "afthab",
					Email:    "afthab606@gmail.com",
					Password: "12345678",
				},
			},
			want: models.User{
				Username:     "afthab",
				Email:        "afthab606@gmail.com",
				PasswordHash: "hashed password",
			}, // Change the expected result to an empty User since an error is expected.
			wantErr: false, // Set wantErr to true since an error is expected.
			mockResponse: func() (models.User, error) {
				return models.User{
					Username:     "afthab",
					Email:        "afthab606@gmail.com",
					PasswordHash: "hashed password",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRespo := repository.NewMockUserRepo(mc)
			mockRespo.EXPECT().CreateUser(tt.args.ctx, gomock.Any()).Return(tt.mockResponse()).AnyTimes()

			svc, err := NewService(mockRespo, &auth.MockAuthentication{})
			if err != nil {
				t.Errorf("error in initializing the repo layer")
				return
			}

			got, err := svc.UserSignup(tt.args.ctx, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}
