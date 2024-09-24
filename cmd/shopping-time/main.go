package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Awes0meEman/shopping-time/internal/handlers"
)

func main () {
	router := mux.NewRouter()

	router.HandleFunc("/games", handlers.GetAllGames).Methods("GET")
	router.HandleFunc("/games/{id}", handlers.GetGameById).Methods("GET")
	router.HandleFunc("/games", handlers.CreateGame).Methods("POST")
	router.HandleFunc("/games/{id}", handlers.UpdateGame).Methods("PUT")
	router.HandleFunc("/games/{id}", handlers.DeleteGame).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
