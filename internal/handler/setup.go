package handler

import (
	"log"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/middleware"
	"github.com/afthaab/job-portal/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupApi(a auth.Authentication, svc service.UserService) *gin.Engine {
	r := gin.New()

	m, err := middleware.NewMiddleware(a)
	if err != nil {
		log.Panic("middlewares not setup")
	}

	h, err := NewHandler(svc)
	if err != nil {
		log.Panic("handlers not setup")
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
			company.GET("/job/view/:id", m.Authenticate(h.ViewJob))
		}

		jobs := admin.Group("jobs")
		{
			jobs.POST("/add/:cid", m.Authenticate(h.AddJobs))
			jobs.GET("/view/all", m.Authenticate(h.ViewAllJobs))
			jobs.GET("/view/:id", m.Authenticate(h.ViewJobByID))
		}
	}

	return r

}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "ok",
	})
}
