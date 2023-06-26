package main

import (
	apigateway_svc "github.com/The-Origin-Labs/landate/api-gateway"
	document_svc "github.com/The-Origin-Labs/landate/document"
	storage_svc "github.com/The-Origin-Labs/landate/storage"
)

func main() {

	// api-gateway service
	go apigateway_svc.ApiGatewaySVC()

	// storage service
	go storage_svc.StorageSVC()

	// document service
	go document_svc.DocumentSVC()

	select {}
}
