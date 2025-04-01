package db

import (
	"fmt"
	"github.com/TiagoNora/GoCRUDV2/schemas"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Error loading file .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msgf("Failed to connect to the database: %v", err)
	}

	err = DB.AutoMigrate(&schemas.Product{})
	if err != nil {
		log.Fatal().Msg("Could not migrate the database")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal().Msgf("Failed to connect to the database: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Info().Msg("Connect to database!")
}
