package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string
	Id   int
}

func customers(w http.ResponseWriter, r *http.Request) {
	c := Customer{Name: "Seyi Gay", Id: 1}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/customers", customers).Methods("GET").Name("customers")

	err := http.ListenAndServe(":8000", router)

	if err != nil {
		log.Fatal(err)
	}
}
