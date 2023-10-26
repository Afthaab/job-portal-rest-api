package database

import (
	"github.com/afthaab/job-portal/internal/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	//if s.db.Migrator().HasTable(&User{}) {
	//	returngorm
	//}

	err := db.Migrator().DropTable(&models.User{})
	if err != nil {
		return err
	}

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	return nil
}
