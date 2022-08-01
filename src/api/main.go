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
	router.HandleFunc("/pick", endpoints.CreatePick).Methods("POST")
	router.HandleFunc("/matchup", endpoints.CreateMatchup).Methods("POST")
	router.HandleFunc("/week", endpoints.CreateWeek).Methods("POST")
	router.HandleFunc("/user", endpoints.CreateUser).Methods("POST")
	router.HandleFunc("/team", endpoints.CreateTeam).Methods("POST")

	router.HandleFunc("/picks", endpoints.GetPick).Methods("GET")
	router.HandleFunc("/matchups", endpoints.GetMatchup).Methods("GET")
	router.HandleFunc("/weekMatchups", endpoints.GetWeekMatchups).Methods("GET")
	router.HandleFunc("/weeks", endpoints.GetWeek).Methods("GET")
	router.HandleFunc("/users", endpoints.GetUser).Methods("GET")
	router.HandleFunc("/teams", endpoints.GetTeam).Methods("GET")
	router.HandleFunc("/scores", endpoints.GetScore).Methods("GET")

	router.HandleFunc("/pick/{id}", endpoints.UpdatePick).Methods("PATCH")
	router.HandleFunc("/matchup/{id}", endpoints.UpdateMatchup).Methods("PATCH")
	router.HandleFunc("/week/{id}", endpoints.UpdateWeek).Methods("PATCH")
	router.HandleFunc("/user/{id}", endpoints.UpdateUser).Methods("PATCH")
	router.HandleFunc("/team/{id}", endpoints.UpdateTeam).Methods("PATCH")

	router.HandleFunc("/pick/{id}", endpoints.DeletePick).Methods("DELETE")
	router.HandleFunc("/matchup/{id}", endpoints.DeleteMatchup).Methods("DELETE")
	router.HandleFunc("/week/{id}", endpoints.DeleteWeek).Methods("DELETE")
	router.HandleFunc("/user/{id}", endpoints.DeleteUser).Methods("DELETE")
	router.HandleFunc("/team/{id}", endpoints.DeleteTeam).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", router))
}
