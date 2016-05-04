package main

import (
	"log"
	"net/http"
	"github.com/antifragilesoftware/simple-gokit-microservice/microservice"
)

func main() {

	microservice.AddServices()

	log.Fatal(http.ListenAndServe(":8080", nil))
}


