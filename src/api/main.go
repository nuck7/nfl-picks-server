package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nuck7/nfl-picks-server/src/api/endpoints"
)

func Api() {
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", endpoints.CreatePicks).Methods("POST")
	router.HandleFunc("/get/{id}", endpoints.CreateScores).Methods("GET")
	router.HandleFunc("/create", endpoints.CreateMatchups).Methods("POST")

	router.HandleFunc("/get/{id}", endpoints.UpdateScores).Methods("PATCH")

	router.HandleFunc("/get/{id}", endpoints.GetPicks).Methods("GET")
	router.HandleFunc("/get/{id}", endpoints.GetScores).Methods("GET")
	router.HandleFunc("/get/{id}", endpoints.GetMatchups).Methods("GET")

	router.HandleFunc("/delete/{id}", endpoints.DeletePicks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", router))
}
