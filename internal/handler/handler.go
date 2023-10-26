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
	user := r.Group("/user")
	{
		user.POST("/signup", h.SignUp)
		user.POST("/signin", h.Signin)
	}
	admin := r.Group("/admin")
	{
		company := admin.Group("company")
		{
			company.POST("/add", m.Authenticate(h.AddCompany))
			company.GET("/view/all", m.Authenticate(h.ViewAllCompanies))
			company.GET("/view/:id", m.Authenticate(h.ViewCompany))
		}
	}

	return r

}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "ok",
	})
}
