package main

import (
	apigateway_svc "landate/api-gateway"
	document_svc "landate/document"
	storage_svc "landate/storage"
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
