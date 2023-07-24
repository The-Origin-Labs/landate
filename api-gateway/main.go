package apigateway

import (
	// service "landate/api-gateway/discovery"

	api "landate/api-gateway/handlers"
	config "landate/config"
	consul "landate/consul"
	"log"
	"strconv"
)

const (
	svcId   = "apigateway_id"
	svcName = "apigateway_svc"
	envPORT = "API_GATEWAY_PORT"
	svcTag  = "api-gateway"
)

func ApiGatewaySVC() {
	// Register Service
	svc_port, err := strconv.Atoi(config.GetEnvConfig(envPORT))
	if err != nil {
		log.Fatal(err)
	}
	api_svc := consul.NewService(svcId, svcName, svcTag, svc_port)

	api.Gateway()
	api_svc.DeregisterService()

}
