package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/api/assets", GetAssets).Methods("GET")
	router.HandleFunc("/api/assets", CreateAsset).Methods("POST")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
