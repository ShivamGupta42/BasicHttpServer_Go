package TimeApiMuxServer

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartTimeServer() {

	//Register Routes
	m := mux.NewRouter()
	m.HandleFunc("/api/time", BaseHandler)
	//Start Servers
	log.Fatal(http.ListenAndServe(":3001", m))

}
