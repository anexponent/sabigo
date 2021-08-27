package controllers

import (
	"fmt"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We have hit the home endpoint")
}
