package main

import (
	"log"
	"net/http"

	"github.com/Yonas21/go-habit-tracker/routers"
)

func main() {

	r := routers.SetupRoutes()
	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
