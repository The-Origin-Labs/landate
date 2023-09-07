package db

import (
	"fmt"
	"landate/authentication/models"
	"landate/config"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDBInstance() *gorm.DB {
	return DB
}

func PGConnect() error {
	dbport := config.GetEnvConfig("DB_PORT")

	// parsing string to int
	DB_PORT, err := strconv.ParseInt(dbport, 10, 32)
	if err != nil {
		log.Fatal("Unable to parse port string to int")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.GetEnvConfig("DB_HOST"),
		config.GetEnvConfig("DB_USER"),
		config.GetEnvConfig("DB_PASSWORD"),
		config.GetEnvConfig("DB_NAME"),
		DB_PORT)
	// postgresURI := config.GetEnvConfig("POSTGRES_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to Database.", err)
	}
	log.Println("Connected ✅")
	// db.Logger = logger.Default.LogMode(logger.Info)

	// Migrate the schema to Database
	log.Println("Running migrations...")
	err = db.AutoMigrate(&models.User{})

	DB = db
	return err
}
