package main

import (
	// service "github.com/The-Origin-Labs/landate/api-gateway/discovery"
	api "github.com/The-Origin-Labs/landate/api-gateway/handlers"
	consul "github.com/The-Origin-Labs/landate/consul"
)

func main() {
	// Register Service
	// service.ServiceDicovery()
	// service.ServiceRegistry()
	consul.ServiceDicovery("apigateway-service")
	api.Gateway()
}
