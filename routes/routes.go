package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/harshcasper/golang-docker-template/app"
)

type Handler = func(writer http.ResponseWriter, request *http.Request)

func Register() {
	log.Println("Registering routes")
	routes := map[string]Handler{
		"/":      home,
		"/users": users,
	}

	for route, handler := range routes {
		http.HandleFunc(route, handler)
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "{\"message\": \"Hello World\"}")
}

func users(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Received request for %s\n", request.URL.Path)
	users, fetchError := app.GetUsers()
	if fetchError != nil {
		log.Printf("Error fetching users: %s\n", fetchError.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Found %d users\n", len(users))
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf("{\"users\": %s}", users)))
}
