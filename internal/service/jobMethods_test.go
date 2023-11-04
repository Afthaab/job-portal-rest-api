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

func TestService_AddJobDetails(t *testing.T) {
	type args struct {
		ctx     context.Context
		jobData models.Jobs
		Cid     uint64
	}
	tests := []struct {
		name             string
		args             args
		want             models.Jobs
		wantErr          bool
		mockRepoResponse func() (models.Jobs, error)
	}{
		{
			name: "database success",
			args: args{
				ctx: context.Background(),
				jobData: models.Jobs{
					Cid:          1,
					Name:         "Junior web developer",
					NoticePeriod: "30",
					Salary:       "10000",
				},
				Cid: 1,
			},
			want: models.Jobs{
				Cid:          1,
				Name:         "Junior web developer",
				NoticePeriod: "30",
				Salary:       "10000",
			},
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Cid:          1,
					Name:         "Junior web developer",
					NoticePeriod: "30",
					Salary:       "10000",
				}, nil
			},
		},
		{
			name: "database error",
			args: args{
				ctx: context.Background(),
				jobData: models.Jobs{
					Cid:          1,
					Name:         "Junior web developer",
					NoticePeriod: "30",
					Salary:       "10000",
				},
				Cid: 1,
			},
			want:    models.Jobs{},
			wantErr: true,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, errors.New("could not create the jobs")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().CreateJob(tt.args.ctx, tt.args.jobData).Return(tt.mockRepoResponse()).AnyTimes()

			svc, err := NewService(mockRepo, &auth.Auth{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}
			got, err := svc.AddJobDetails(tt.args.ctx, tt.args.jobData, tt.args.Cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobDetailsBy(tt.args.ctx, tt.args.jid).Return(tt.mockRepoResponse()).AnyTimes()
			}
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
		name             string
		args             args
		want             []models.Jobs
		wantErr          bool
		mockRepoResponse func() ([]models.Jobs, error)
	}{
		{
			name: "database success",
			args: args{
				ctx: context.Background(),
			},
			want: []models.Jobs{
				{
					Cid:          01,
					Name:         "junio web developer",
					Salary:       "10000",
					NoticePeriod: "30",
				}, {
					Cid:          01,
					Name:         "senior web developer",
					Salary:       "100000",
					NoticePeriod: "50",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return []models.Jobs{
					{
						Cid:          01,
						Name:         "junio web developer",
						Salary:       "10000",
						NoticePeriod: "30",
					}, {
						Cid:          01,
						Name:         "senior web developer",
						Salary:       "100000",
						NoticePeriod: "50",
					},
				}, nil
			},
		},
		{
			name: "database error",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return nil, errors.New("could not find the records in the database")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().FindAllJobs(tt.args.ctx).Return(tt.mockRepoResponse()).AnyTimes()

			svc, err := NewService(mockRepo, &auth.Auth{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}

			got, err := svc.ViewAllJobs(tt.args.ctx)
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

func TestService_ViewJob(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name             string
		args             args
		want             []models.Jobs
		wantErr          bool
		mockRepoResponse func() ([]models.Jobs, error)
	}{
		{
			name: "error in database",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return nil, errors.New("could not view the jobs")
			},
		},
		{
			name: "success from database",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want: []models.Jobs{
				{
					Cid:          1,
					Name:         "Junior web developer",
					Salary:       "100000",
					NoticePeriod: "30",
				},
				{
					Cid:          1,
					Name:         "Senior web developer",
					Salary:       "200000",
					NoticePeriod: "10",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return []models.Jobs{
					{
						Cid:          1,
						Name:         "Junior web developer",
						Salary:       "100000",
						NoticePeriod: "30",
					},
					{
						Cid:          1,
						Name:         "Senior web developer",
						Salary:       "200000",
						NoticePeriod: "10",
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().FindJob(tt.args.ctx, tt.args.cid).Return(tt.mockRepoResponse()).AnyTimes()

			svc, err := NewService(mockRepo, &auth.Auth{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}
			got, err := svc.ViewJob(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
