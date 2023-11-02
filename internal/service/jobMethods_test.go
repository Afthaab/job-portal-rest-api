package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/models"
	"github.com/afthaab/job-portal/internal/repository"
	"go.uber.org/mock/gomock"
)

func TestService_ViewJobById(t *testing.T) {
	type args struct {
		ctx context.Context
		jid uint64
	}
	tests := []struct {
		name             string
		args             args
		want             models.Jobs
		wantErr          bool
		mockRepoResponse func() (models.Jobs, error)
	}{
		{
			name: "error from db",
			want: models.Jobs{},
			args: args{
				ctx: context.Background(),
				jid: 15,
			},
			wantErr: true,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, errors.New("test error")
			},
		},
		{
			name: "success",
			want: models.Jobs{
				Company: models.Company{
					Name: "TCS",
				},
				Cid:  1,
				Name: "SDE",
			},
			args: args{
				ctx: context.Background(),
				jid: 15,
			},
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Company: models.Company{
						Name: "TCS",
					},
					Cid:  1,
					Name: "SDE",
				}, nil
			},
		},
		{
			name: "invalid job id",
			want: models.Jobs{},
			args: args{
				ctx: context.Background(),
				jid: 5,
			},
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, nil
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobDetailsBy(tt.args.ctx, tt.args.jid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			// mockAuth:=
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewJobById(tt.args.ctx, tt.args.jid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJobById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJobById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewAllJobs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    []models.Jobs
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ViewAllJobs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
