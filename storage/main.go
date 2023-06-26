package storage

import (
	"log"
	"strconv"

	config "landate/config"
	consul "landate/consul"

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
	// storage_svc.StartDicovery()

	// api entrypoint
	api.StorageService()
	storage_svc.DeregisterService()

}
