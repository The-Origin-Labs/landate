package main

import (
	"log"
	"strconv"

	config "github.com/The-Origin-Labs/landate/config"
	consul "github.com/The-Origin-Labs/landate/consul"

	database "github.com/The-Origin-Labs/landate/storage/database"
	api "github.com/The-Origin-Labs/landate/storage/routes"
)

const (
	svcId   = "storage_svc_id"
	svcName = "storage_svc"
	envPORT = "STORAGE_SERVICE_PORT"
	svcTag  = "storage"
)

func main() {
	// Database connection
	database.DBConnect()

	svc_port, err := strconv.Atoi(config.GetEnvConfig(envPORT))
	if err != nil {
		log.Fatal(err)
	}

	storage_svc := consul.NewService(svcId, svcName, svcTag, svc_port)
	// storage_svc.StartDicovery()

	// api entrypoint
	api.StorageService()
	storage_svc.DeregisterService()

}
