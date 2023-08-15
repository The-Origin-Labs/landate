package consul

import (
	"log"
	"time"

	"github.com/hashicorp/consul/api"
)

type Service struct {
	ID           string
	Name         string
	Port         int
	Tags         []string
	checkId      string
	consulClient *api.Client
	serverId     string
	ttl          string
}

var (
	ttl            = time.Second * 8 // time to leave
	serviceAddress = "consul:8500"   // for production
	// serviceAddress = "localhost"
)

/*
@description:
Creates a new instance of the Service struct and
registers the service with Consul. Starts a background
goroutine to periodically update the health check status of the service.

@params:
serviceID: string - Unique identifier for the service.
serviceName: string - Name of the service.
serviceTag: string - Tag(s) associated with the service.
servicePort: int - Port number on which the service is running.

@returns:
*Service: A pointer to the created Service instance.
*/
func NewService(serviceID, serviceName, serviceTag string, servicePort int) *Service {
	client, err := api.NewClient(&api.Config{
		Address: serviceAddress,
	})
	if err != nil {
		log.Fatal(err)
	}

	micro_svc := &Service{
		ID:           serviceID,
		Name:         serviceName,
		Tags:         []string{serviceTag},
		Port:         servicePort,
		consulClient: client,
		checkId:      "check_health_" + serviceID,
		ttl:          ttl.String(),
	}

	if err := micro_svc.registerService(); err != nil {
		log.Fatal(err)
	}

	go micro_svc.updateHealthCheck()

	return micro_svc
}

/*
@description:
Periodically updates the health check status of the
service to Consul, indicating that the service is online.
*/
func (svc *Service) updateHealthCheck() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		err := svc.consulClient.Agent().UpdateTTL(svc.checkId, "service online", api.HealthPassing)
		if err != nil {
			log.Println("Failed to update health check:", err)
		}
		<-ticker.C
	}
}

/*
@description:
Deregisters the service from Consul, removing it from the service registry.
*/
func (svc *Service) DeregisterService() {
	if err := svc.consulClient.Agent().ServiceDeregister(svc.serverId); err != nil {
		log.Println("Failed to deregister service: ", err)
	}
}

/*
@description:
Registers the service with Consul, adding it to the service registry.

@returns:
*error: An error if registration fails, otherwise nil.
*/

func (svc *Service) registerService() error {
	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: ttl.String(),
		TLSSkipVerify:                  true,
		TTL:                            ttl.String(),
		CheckID:                        svc.checkId,
	}

	service := &api.AgentServiceRegistration{
		ID:      svc.serverId,
		Name:    svc.Name,
		Tags:    svc.Tags,
		Address: serviceAddress,
		Port:    svc.Port,
		Check:   check,
	}

	return svc.consulClient.Agent().ServiceRegister(service)
}
