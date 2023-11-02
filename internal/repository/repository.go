package repository

import (
	"context"
	"errors"

	"github.com/afthaab/job-portal/internal/models"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=repository

type UserRepo interface {
	CreateUser(ctx context.Context, userData models.User) (models.User, error)
	CheckEmail(ctx context.Context, email string) (models.User, error)

	CreateCompany(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error)

	CreateJob(ctx context.Context, jobData models.Jobs) (models.Jobs, error)
	FindJob(ctx context.Context, cid uint64) ([]models.Jobs, error)
	FindAllJobs(ctx context.Context) ([]models.Jobs, error)
	ViewJobDetailsBy(ctx context.Context, jid uint64) (models.Jobs, error)
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
