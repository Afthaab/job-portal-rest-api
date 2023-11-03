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

func TestService_AddCompanyDetails(t *testing.T) {
	type args struct {
		ctx         context.Context
		companyData models.Company
	}
	tests := []struct {
		name         string
		args         args
		want         models.Company
		wantErr      bool
		mockResponse func() (models.Company, error)
	}{
		{
			name: "error in the database",
			args: args{
				ctx: context.Background(),
				companyData: models.Company{
					Name:     "Infosys",
					Location: "Bangalore",
					Field:    "IT",
				},
			},
			want:    models.Company{},
			wantErr: true,
			mockResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("could not add the data to the database")
			},
		},
		{
			name: "success from database",
			args: args{
				ctx: context.Background(),
				companyData: models.Company{
					Name:     "Infosys",
					Location: "Bangalore",
					Field:    "IT",
				},
			},
			want: models.Company{
				Name:     "Infosys",
				Location: "Bangalore",
				Field:    "IT",
			},
			wantErr: false,
			mockResponse: func() (models.Company, error) {
				return models.Company{
					Name:     "Infosys",
					Location: "Bangalore",
					Field:    "IT",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().CreateCompany(tt.args.ctx, tt.args.companyData).Return(tt.mockResponse()).AnyTimes()

			svc, err := NewService(mockRepo, &auth.Auth{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}

			got, err := svc.AddCompanyDetails(tt.args.ctx, tt.args.companyData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddCompanyDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddCompanyDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewAllCompanies(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		args         args
		want         []models.Company
		wantErr      bool
		mockResponse func() ([]models.Company, error)
	}{
		{
			name: "error from database",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mockResponse: func() ([]models.Company, error) {
				return nil, errors.New("test error from the  mock function")
			},
		},
		{
			name: "success from database",
			args: args{
				ctx: context.Background(),
			},
			want: []models.Company{
				{
					Name:     "Bosch",
					Location: "Whitefield",
					Field:    "IT",
				},
				{
					Name:     "Allegis",
					Location: "Koramangala",
					Field:    "IT",
				},
			},
			wantErr: false,
			mockResponse: func() ([]models.Company, error) {
				return []models.Company{
					{
						Name:     "Bosch",
						Location: "Whitefield",
						Field:    "IT",
					},
					{
						Name:     "Allegis",
						Location: "Koramangala",
						Field:    "IT",
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().ViewCompanies(tt.args.ctx).Return(tt.mockResponse()).AnyTimes()

			svc, err := NewService(mockRepo, &auth.Auth{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}
			got, err := svc.ViewAllCompanies(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewAllCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewAllCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewCompanyDetails(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name         string
		args         args
		want         models.Company
		wantErr      bool
		mockResponse func() (models.Company, error)
	}{
		{
			name: "error from the database",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want:    models.Company{},
			wantErr: true,
			mockResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("test error from the  mock function")
			},
		},
		{
			name: "success from the database",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want: models.Company{
				Name:     "Infosys",
				Location: "Chennai",
				Field:    "Web Development",
			},
			wantErr: false,
			mockResponse: func() (models.Company, error) {
				return models.Company{
					Name:     "Infosys",
					Location: "Chennai",
					Field:    "Web Development",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().ViewCompanyById(tt.args.ctx, tt.args.cid).Return(tt.mockResponse()).AnyTimes()

			svc, err := NewService(mockRepo, &auth.Auth{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}
			got, err := svc.ViewCompanyDetails(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewCompanyDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewCompanyDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
