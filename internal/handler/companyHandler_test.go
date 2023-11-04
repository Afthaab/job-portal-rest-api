package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/middleware"
	"github.com/afthaab/job-portal/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Test_handler_ViewCompany(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "misssing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httprequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httprequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "misssing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "invalid job id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})

				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		// {
		// 	name: "error while fetching jobs from service",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
		// 		ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
		// 		mc := gomock.NewController(t)
		// 		ms := service.NewMockUserService(mc)

		// 		ms.EXPECT().ViewJobById(c.Request.Context(), gomock.Any()).Return(models.Jobs{}, errors.New("test service error")).AnyTimes()

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusInternalServerError,
		// 	expectedResponse:   `{"error":"test service error"}`,
		// },
		// {
		// 	name: "success",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
		// 		ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
		// 		mc := gomock.NewController(t)
		// 		ms := service.NewMockUserService(mc)

		// 		ms.EXPECT().ViewJobById(c.Request.Context(), gomock.Any()).Return(models.Jobs{}, nil).AnyTimes()

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusOK,
		// 	expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"cid":0,"name":"","salary":"","notice_period":""}`,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			c, rr, _ := tt.setup()

			h, err := NewHandler(&service.Service{})
			if err != nil {
				t.Errorf("error is initializing the handler layer")
				return
			}
			h.ViewCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())

		})
	}
}
