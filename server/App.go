package server

import (
	"log"
	"net/http"
)

func StartServer() {

	mux := http.NewServeMux()

	//Register Routes
	mux.HandleFunc("/Route1", Route1)
	mux.HandleFunc("/Route2", Route2)
	mux.HandleFunc("/Route3", Route3)
	mux.HandleFunc("/employees", GetAllEmployees)

	log.Fatal(http.ListenAndServe(":2001", mux))
}
