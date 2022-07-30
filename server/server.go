package server

import (
	"net/http"
	"sabigo/config"
	"sabigo/logger"
	"sabigo/routes"
)

func Serve() {
	//register log file
	logger.Init()

	port := config.LoadEnvironmentalVariables("PORT")
	//Register Route Package here
	r, err := routes.RegisterRoutes()

	if err != nil {
		logger.Error.Fatal(err)
		panic(err)
	}
	logger.Info.Println("Started Server at port", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		logger.Info.Fatal("Error Starting Server", err)
		panic(err)
	} else {
		logger.Info.Println("Server Started at", port)
	}

}
