package document

import (
	"fmt"
	"log"
	"strconv"

	// db "github.com/The-Origin-Labs/landate/document/db"
	config "github.com/The-Origin-Labs/landate/config"
	consul "github.com/The-Origin-Labs/landate/consul"
	api "github.com/The-Origin-Labs/landate/document/routes"
)

const (
	svcId   = "doc_id"
	svcName = "doc_svc"
	envPORT = "DOCUMENT_SERVICE_PORT"
	svcTag  = "doc_svc"
)

func DocumentSVC() {

	svc_port, err := strconv.Atoi(config.GetEnvConfig(envPORT))
	if err != nil {
		log.Fatal("Unable to load svc port: ", err)
	}
	doc_svc := consul.NewService(svcId, svcName, svcTag, svc_port)

	// Entry Point to API
	err = api.Init()
	if err != nil {
		log.Fatal("Unable to start Document Microserivces")
	} else {
		fmt.Println("Welcome to Document Microserivces")
	}

	doc_svc.DeregisterService()
}
