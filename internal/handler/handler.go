package handler

import (
	"log"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/middleware"
	"github.com/afthaab/job-portal/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupApi(a *auth.Auth, svc service.UserService) *gin.Engine {
	r := gin.New()

	m, err := middleware.NewMiddleware(a)
	if err != nil {
		log.Panic("middlewares not setup")
	}
	h := handler{
		service: svc,
	}

	r.Use(m.Log(), gin.Recovery())

	r.GET("/check", m.Authenticate(Check))
	r.POST("/user/signup", h.SignUp)

	return r

}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "ok",
	})
}
