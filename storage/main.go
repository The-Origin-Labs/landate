package storage

import (
	config "landate/config"
	consul "landate/consul"
	"log"
	"strconv"

	database "landate/storage/database"
	api "landate/storage/routes"
)

const (
	svcId   = "storage_svc_id"
	svcName = "storage_svc"
	envPORT = "STORAGE_SERVICE_PORT"
	svcTag  = "storage"
)

func StorageSVC() {
	// Database connection
	database.DBConnect()

	svc_port, err := strconv.Atoi(config.GetEnvConfig(envPORT))
	if err != nil {
		log.Fatal(err)
	}
	storage_svc := consul.NewService(svcId, svcName, svcTag, svc_port)

	// api entrypoint
	api.StorageService()
	storage_svc.DeregisterService()

}
