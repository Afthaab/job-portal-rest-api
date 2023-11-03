package service

import (
	"context"
	"errors"
	"testing"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/models"
	"github.com/afthaab/job-portal/internal/repository"
	"go.uber.org/mock/gomock"
)

func TestService_UserSignIn(t *testing.T) {
	type args struct {
		ctx      context.Context
		userData models.NewUser
	}
	tests := []struct {
		name         string
		args         args
		want         string
		wantErr      bool
		mockResponse func() (models.User, error)
	}{
		{
			name: "error in the database",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "afthab",
					Email:    "afthab606@gmail.com",
					Password: "1234",
				},
			},
			want:    "",
			wantErr: true,
			mockResponse: func() (models.User, error) {
				return models.User{}, errors.New("test error from the  mock function")
			},
		},
		{
			name: "error in hashing the password",
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
					PasswordHash: "wrong password",
				}, nil
			},
		},
		{
			name: "error in claims",
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
					PasswordHash: "$2a$10$zbU8cCee6SUrxRVTb/GHG.92sshB2JULj1cZwcPqYNQNWwIKfPKzW",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().CheckEmail(tt.args.ctx, tt.args.userData.Email).Return(tt.mockResponse()).AnyTimes()
			svc, err := NewService(mockRepo, &auth.Auth{})
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
