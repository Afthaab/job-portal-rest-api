package handler

import (
	"errors"

	"github.com/afthaab/job-portal/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.UserService
}

type Handerfuncs interface {
	Signin(c *gin.Context)
	SignUp(c *gin.Context)

	ViewCompany(c *gin.Context)
	ViewAllCompanies(c *gin.Context)
	AddCompany(c *gin.Context)

	ViewJobByID(c *gin.Context)
	ViewAllJobs(c *gin.Context)
	ViewJob(c *gin.Context)
	AddJobs(c *gin.Context)
}

func NewHandler(svc service.UserService) (Handerfuncs, error) {
	if svc == nil {
		return nil, errors.New("service interface cannot be null")
	}
	return &handler{
		service: svc,
	}, nil
}
