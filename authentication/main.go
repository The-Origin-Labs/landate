package main

import (
	db "landate/authentication/db"
	auth_api "landate/authentication/routes"

	config "landate/config"
	consul "landate/consul"
	"log"
	"strconv"
)

var (
	svcId   = "auth_id"
	svcName = "auth_svc"
	envPORT = "AUTH_SERVICE_PORT"
	svcTag  = "auth"
)

// This Service contain User Authentication information
// and Its Profile Information
func main() {

	if err := db.PGConnect(); err != nil {
		log.Fatal("Unable to connect auth-svc to database.")
	}

	svc_port, err := strconv.Atoi(config.GetEnvConfig(envPORT))
	if err != nil {
		log.Fatal(err)
	}
	auth_svc := consul.NewService(svcId, svcName, svcTag, svc_port)

	if err := auth_api.Init(); err != nil {
		log.Fatalln("Unable to start auth service.")
	}
	auth_svc.DeregisterService()

}
