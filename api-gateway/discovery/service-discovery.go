package discovery

import (
	"log"

	"github.com/hashicorp/consul/api"
)

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
