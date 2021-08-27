package server

import (
	"log"
	"net/http"
	"sabigo/config"
	"sabigo/routes"
	"sabigo/utils"
)

func Serve() {
	//register log file
	utils.Log()
	port := config.LoadEnvironmentalVariables("PORT")
	//Register Route Package here
	r, err := routes.RegisterRoutes()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Println("Started Server at port", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Println("Error Starting Server", err)
		panic(err)
	} else {
		log.Println("Server Started at", port)
	}

}
