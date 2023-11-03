package auth

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey int

const Key ctxKey = 1

type Auth struct {
	privateKey *rsa.PrivateKey
	publickey  *rsa.PublicKey
}

//go:generate mockgen -source=auth.go -destination=mockModels/auth_mock.go -package=auth

type Authentication interface {
	GenerateAuthToken(claims jwt.RegisteredClaims) (string, error)
	ValidateToken(token string) (jwt.RegisteredClaims, error)
}

func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (Authentication, error) {
	if privateKey == nil && publicKey == nil {
		return nil, errors.New("publickey and privatekey cannot be null")
	}
	return &Auth{
		privateKey: privateKey,
		publickey:  publicKey,
	}, nil
}
