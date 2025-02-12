package initializers

import (
	// "os"


	"gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error;
	dsn := "host=localhost user=postgres password=3055 dbname=golang port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}
}