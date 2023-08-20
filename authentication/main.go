package main

import (
	auth_api "landate/authentication/routes"
	"log"
)

// This Service contain User Authentication information
// and Its Profile Information
func main() {

	if err := auth_api.Init(); err != nil {
		log.Fatalln("Unable to start auth service.")
	}

}
