package BasicServer

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartBasicServer() {

	r := mux.NewRouter()
	//Register Routes
	r.HandleFunc("/employees", GetAllEmployees).Methods(http.MethodGet)
	r.HandleFunc("/employees/{employee:[0-9]+}", GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employees/{employee:[0-9]+}", CreateEmployee).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":2001", r))
}
