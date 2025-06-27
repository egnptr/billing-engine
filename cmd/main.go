package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/egnptr/billing-engine/handlers"
	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8080"
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	}).Methods("GET")
	r.HandleFunc("/loan", handlers.CreateLoan).Methods("POST")
	r.HandleFunc("/loan/{id}", handlers.GetLoan).Methods("GET")
	r.HandleFunc("/loan/{id}/payment", handlers.MakePayment).Methods("POST")
	r.HandleFunc("/loan/{id}/outstanding", handlers.GetOutstanding).Methods("GET")
	r.HandleFunc("/loan/{id}/delinquent", handlers.IsDelinquent).Methods("GET")

	log.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
