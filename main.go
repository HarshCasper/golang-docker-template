package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/harshcasper/golang-docker-template/routes"
)

func main() {
	port := env("PORT", "3000")
	handler(fmt.Sprintf(":%s", port))
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func handler(address string) error {
	routes.Register()
	log.Printf("Listening on %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
