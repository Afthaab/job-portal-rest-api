package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/database"
	"github.com/afthaab/job-portal/internal/handler"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {

}

func StartApp() error {
	// =========================================================================
	// initializing the authentication support
	log.Info().Msg("main started : initializing the authentication support")

	//reading the private key file
	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		return fmt.Errorf("error in reading auth private key : %w", err) // %w is used for error wraping
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("error in parsing auth private key : %w", err) // %w is used for error wraping
	}
	publicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		return fmt.Errorf("error in reading auth public key : %w", err) // %w is used for error wraping
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("error in parsing auth public key : %w", err) // %w is used for error wraping
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("error in constructing auth %w", err)
	}

	// =========================================================================
	// start the database

	log.Info().Msg("main started : initializing the data")

	db, err := database.ConnectToDatabase()
	if err != nil {
		return fmt.Errorf("error in opening the database connection : %w", err)
	}

	pg, err := db.DB()
	if err != nil {
		return fmt.Errorf("error in getting the database instance")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("database is not connected: %w", err)
	}

	// initializing the http server
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      handler.SetupApi(a),
	}

}
