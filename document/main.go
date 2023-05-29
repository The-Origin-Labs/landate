package main

import (
	"fmt"
	"log"

	// db "github.com/The-Origin-Labs/landate/document/db"
	api "github.com/The-Origin-Labs/landate/document/routes"
)

func main() {

	// Entry Point to API
	err := api.Init()
	if err != nil {
		log.Fatal("Unable to start Document Microserivces")
	} else {
		fmt.Println("Welcome to Document Microserivces")
	}
}
