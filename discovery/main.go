package main

import (
	"log"

	"github.com/hashicorp/consul/api"
)

func ServiceRegistry() {
	config := api.DefaultConfig()
	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Failed to create Consul client:", err)
	}

	// Registering the Service
	registration := new(api.AgentServiceRegistration)
	registration.ID = "ocr-service-id"
	registration.Name = "ocr-service"
	registration.Address = "127.0.0.1"
	registration.Port = 8002
	registration.Tags = []string{"http"}

	// Enable health monitoring for the service,
	// HTTP health-check endpoint is provided while
	// registering the service with the consul
	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = "http://127.0.0.1:8002"
	registration.Check.Interval = "30s"
	registration.Check.Timeout = "1s"

	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("Failed to register service:", err)
	}

	log.Println("Service registered successfully")
}

func ServiceDicovery() {
	config := api.DefaultConfig()
	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Failed to create Consul client:", err)
	}

	services, _, err := consulClient.Catalog().Service("ocr-service", "", nil)
	if err != nil {
		log.Fatal("Failed to query services:", err)
	}

	for _, service := range services {
		log.Printf("Found service: ID=%s, Address=%s, Port=%d\n", service.ServiceID, service.ServiceAddress, service.ServicePort)
	}
}

func main() {

}
