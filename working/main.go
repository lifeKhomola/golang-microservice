package main

import (
	"net/http"
	"working/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.UserRoutes(router)
	http.ListenAndServe(":9090",router)
}
