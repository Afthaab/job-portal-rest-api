package repository

import "github.com/afthaab/job-portal/internal/models"

func (r *Repo) CreateUser(UserDetails models.User) (models.User, error) {
	result := r.DB.Create(&UserDetails)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return UserDetails, nil
}
