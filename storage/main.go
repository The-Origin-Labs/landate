package main

import (
	database "github.com/The-Origin-Labs/landate/storage/database"
	api "github.com/The-Origin-Labs/landate/storage/routes"
)

func main() {
	// Database connection
	database.DBConnect()
	// api entrypoint
	api.StorageService()
}
