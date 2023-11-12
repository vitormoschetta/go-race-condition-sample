package main

import (
	"fmt"
	"log"
	"net/http"

	"testserver/internal"

	"github.com/gorilla/mux"
	"github.com/vitormoschetta/go/pkg/middlewares"
)

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.AcceptJSON)

	router.HandleFunc("/api/execute1", internal.Execute1).Methods("GET")
	router.HandleFunc("/api/execute2", internal.Execute2).Methods("GET")

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
