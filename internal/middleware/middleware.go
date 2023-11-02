package middleware

import (
	"fmt"

	"github.com/afthaab/job-portal/internal/auth"
)

type Mid struct {
	auth auth.Authentication
}

func NewMiddleware(a auth.Authentication) (Mid, error) {
	if a == nil {
		return Mid{}, fmt.Errorf("auth cant be null")
	}
	return Mid{
		auth: a,
	}, nil
}
