package repository

import (
	"context"
	"errors"

	"github.com/afthaab/job-portal/internal/models"
	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateCompany(ctx context.Context, companyData models.Company) (models.Company, error) {
	result := r.db.Create(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not create the company")
	}
	return companyData, nil
}

func (r *Repo) ViewCompanies(ctx context.Context) ([]models.Company, error) {
	var userDetails []models.Company
	result := r.db.Find(&userDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find the companies")
	}
	return userDetails, nil
}

func (r *Repo) ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error) {
	var companyData models.Company
	result := r.db.Where("id = ?", cid).First(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not find the company")
	}
	return companyData, nil
}
