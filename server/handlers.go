package server

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Employee struct {
	Name    string `json:"employee_name" xml:"emp_name"`
	Age     string `json:"age" xml:"age"`
	PinCode string `json:"pincode" xml:"pincode"`
}

func Route1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Route1")
}

func Route2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Route2")
}
func Route3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Route3")
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
