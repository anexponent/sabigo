package routes

import (
	"sabigo/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() (r *mux.Router, err error) {
	r = mux.NewRouter().StrictSlash(true)
	//Home Route
	r.HandleFunc("/", controllers.HomeController).Methods("GET")

	//Authentication Routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	r.HandleFunc("profile", controllers.Profile).Methods("POST")
	r.HandleFunc("/reset", controllers.Reset).Methods("POST")
	return r, err
}
