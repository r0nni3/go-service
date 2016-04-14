package main

import (
	"log"
	"net/http"

	"github.com/r0nni3/go-service/routes"
)

func main() {
	router := routes.NewRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server Started")
}
