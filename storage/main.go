package main

import (
	consul "github.com/The-Origin-Labs/landate/consul"

	database "github.com/The-Origin-Labs/landate/storage/database"
	api "github.com/The-Origin-Labs/landate/storage/routes"
)

func main() {
	// Database connection
	database.DBConnect()
	consul.ServiceDicovery("storage-service")

	// api entrypoint
	api.StorageService()
}
