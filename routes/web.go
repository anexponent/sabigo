package routes

import (
	"sabigo/controllers"
	"sabigo/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes() (r *mux.Router, err error) {
	r = mux.NewRouter().StrictSlash(true)
	auth := r.PathPrefix("/api").Subrouter()
	unauth := r.PathPrefix("/api").Subrouter()
	auth.Use(middleware.AuthMiddleware)

	unauth.HandleFunc("/", controllers.HomeController).Methods("GET")
	unauth.HandleFunc("/register", controllers.Register).Methods("POST")
	unauth.HandleFunc("/login", controllers.Login).Methods("POST")

	//Authenticated Routes
	auth.HandleFunc("profile", controllers.Profile).Methods("POST")
	auth.HandleFunc("/reset", controllers.Reset).Methods("POST")
	auth.HandleFunc("/logout", controllers.Logout).Methods("POST")
	return r, err
}
