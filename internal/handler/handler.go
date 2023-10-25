package handler

import (
	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupApi(a *auth.Auth) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Log())

}
