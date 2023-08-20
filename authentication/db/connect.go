package db

import (
	"landate/authentication/models"
	"landate/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDBInstance() *gorm.DB {
	return DB
}

func PGConnect() {
	postgresURI := config.GetEnvConfig("POSTGRES_URI")
	db, err := gorm.Open(postgres.Open(postgresURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to Database.", err)
	}
	log.Println("Connected ✅")
	db.Logger = logger.Default.LogMode(logger.Info)

	// Migrate the schema to Database
	log.Println("Running migrations...")
	db.AutoMigrate(&models.User{})

	DB = db

}
