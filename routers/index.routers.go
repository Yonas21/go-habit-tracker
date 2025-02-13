package routers

import (
	"github.com/Yonas21/go-habit-tracker/controllers"
	"github.com/gorilla/mux"
	// Add controllers package import once you create it
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Commented out until controllers package is created
	// router.HandleFunc("/api/v1/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	return router // Add missing return statement
}
