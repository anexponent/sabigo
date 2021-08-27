package routes

import (
	"sabigo/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() (r *mux.Router, err error) {
	r = mux.NewRouter()
	r.HandleFunc("/", controllers.HomeController).Methods("GET")

	return r, err
}
