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

const (
	ttl            = time.Second * 8 // time to leave
	serviceAddress = "127.0.0.1"
)

func NewService(serviceID, serviceName, serviceTag string, servicePort int) *Service {
	client, err := api.NewClient(&api.Config{})
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

func (svc *Service) DeregisterService() {
	if err := svc.consulClient.Agent().ServiceDeregister(svc.serverId); err != nil {
		log.Println("Failed to deregister service: ", err)
	}
}

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

// 	check := &api.AgentServiceCheck{
// 		DeregisterCriticalServiceAfter: ttl.String(),
// 		TLSSkipVerify:                  true,
// 		TTL:                            ttl.String(),
// 		CheckID:                        checkID,
// 	}
// 	iport, err := strconv.Atoi(port)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	register := &api.AgentServiceRegistration{
// 		ID:      svcId,
// 		Name:    svcName,
// 		Tags:    []string{svcId},
// 		Address: serviceAddress,
// 		Port:    iport,
// 		Check:   check,
// 	}
// 	if err := svc.consulClient.Agent().ServiceRegister(register); err != nil {
// 		log.Fatal(err)
// 	}
// }
