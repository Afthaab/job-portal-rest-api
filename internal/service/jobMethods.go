package service

import (
	"context"

	"github.com/afthaab/job-portal/internal/models"
)

func (s *Service) ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error) {
	jobData, err := s.UserRepo.ViewJobDetailsBy(ctx, jid)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) ViewAllJobs(ctx context.Context) ([]models.Jobs, error) {
	jobDatas, err := s.UserRepo.FindAllJobs(ctx)
	if err != nil {
		return nil, err
	}
	return jobDatas, nil

}

func (s *Service) AddJobDetails(ctx context.Context, jobData models.Jobs, cid uint64) (models.Jobs, error) {
	jobData.Cid = uint(cid)
	jobData, err := s.UserRepo.CreateJob(ctx, jobData)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) ViewJob(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	jobData, err := s.UserRepo.FindJob(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
