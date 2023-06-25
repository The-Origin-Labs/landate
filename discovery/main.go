package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/The-Origin-Labs/landate/config"
	"github.com/hashicorp/consul/api"
)

func ServiceRegistry(svcId string, svc string, addr string, port string) {
	config := api.DefaultConfig()
	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Failed to create Consul client:", err)
	}

	// Registering the Service
	registration := new(api.AgentServiceRegistration)
	registration.ID = svcId // "ocr-service-id"
	registration.Name = svc // "ocr-service"
	registration.Address = addr
	svc_port, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	registration.Port = svc_port
	registration.Tags = []string{"http"}

	// Enable health monitoring for the service,
	// HTTP health-check endpoint is provided while
	// registering the service with the consul

	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = "http://" + addr + ":" + port // "http://127.0.0.1:8002"
	registration.Check.Interval = "30s"
	registration.Check.Timeout = "1s"

	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("Failed to register service:", err)
	}

	log.Printf("%s service registered successfully", svc)
}

func ServiceDicovery(svc string) {
	config := api.DefaultConfig()
	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Failed to create Consul client:", err)
	}

	services, _, err := consulClient.Catalog().Service(svc, "", nil)
	if err != nil {
		log.Fatal("Failed to query services:", err)
	}

	for _, service := range services {
		log.Printf("Found service: ID=%s, Address=%s, Port=%d\n", service.ServiceID, service.ServiceAddress, service.ServicePort)
	}
}

func main() {

	addr := "127.0.0.1"

	// Register API Gateway Service
	ServiceRegistry("api-gateway-id", "apigateway-service", addr, config.GetEnvConfig("API_GATEWAY_PORT"))

	// Register Storage Service
	ServiceRegistry("storage-service-id", "storage-service", addr, config.GetEnvConfig("STORAGE_SERVICE_PORT"))

	// ServiceDicovery("api-service")
	// ServiceDicovery("storage-service")

}
