package BasicServer

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Employee struct {
	Name    string `json:"employee_name" xml:"emp_name"`
	Age     string `json:"age" xml:"age"`
	PinCode string `json:"pincode" xml:"pincode"`
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees := []Employee{
		{"Shivam", "27", "40071"},
		{"Anupriya", "27", "40071"},
	}

	if contentType := r.Header.Get("Content-Type"); contentType == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(employees)
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	w.WriteHeader(201)
	fmt.Fprintf(w, m["employee"])
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	w.WriteHeader(201)
	fmt.Fprintf(w, m["employee"]+" employee created")

}
