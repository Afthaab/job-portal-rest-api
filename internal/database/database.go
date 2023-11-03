package database

import (
	"github.com/afthaab/job-portal/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=12345 dbname=jportal port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err = db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Jobs{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return nil, err
	}

	return db, nil
}
