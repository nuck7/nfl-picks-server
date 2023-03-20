package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nuck7/nfl-picks-server/src/database"
	"github.com/nuck7/nfl-picks-server/src/models"
	"github.com/nuck7/nfl-picks-server/src/types"
)

func GetMatchup(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var params = r.URL.Query()
	var matchups []models.Matchup
	json.Unmarshal(requestBody, &matchups)

	result := database.Connector.Where("")

	if params["id"] != nil {
		id, _ := strconv.ParseUint(params["id"][0], 0, 64)
		result.Where("id = ?", id)
	}

	if params["homeTeamID"] != nil {
		homeTeamID, _ := strconv.ParseUint(params["home_team_id"][0], 0, 64)
		result.Where("home_team_id = ?", homeTeamID)
	}

	if params["awayTeamID"] != nil {
		awayTeamID, _ := strconv.ParseUint(params["away_team_id"][0], 0, 64)
		result.Where("away_team_id = ?", awayTeamID)
	}

	if params["weekID"] != nil {
		weekID, _ := strconv.ParseUint(params["week_id"][0], 0, 64)
		result.Where("week_id = ?", weekID)
	}

	result.Find(&matchups)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(matchups)
}

func CreateMatchups(w http.ResponseWriter, r *http.Request) {
	// fmt.Print("CREATE MATCHUPS")

	requestBody, _ := ioutil.ReadAll(r.Body)
	var matchup []models.Matchup
	var MatchupInput []types.CreateMatchupInput
	db := database.Connector

	// fmt.Print("requestBody", requestBody)
	err := json.Unmarshal(requestBody, &MatchupInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.CreateInBatches(&matchup, 20)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(matchup)
}

func UpdateMatchup(w http.ResponseWriter, r *http.Request) {
	var matchup models.Matchup
	requestBody, _ := ioutil.ReadAll(r.Body)
	var matchups types.Matchups

	db := database.Connector
	err := json.Unmarshal([]byte(requestBody), &matchups)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(matchups.Matchups); i++ {
		db.Model(&matchup).Where("id = ?", matchups.Matchups[i].ID).First(&matchup)
		// .Updates(map[string]interface{}{
		// 	"home_team_id": matchups.Matchups[i].HomeTeamID,
		// 	"away_team_id": matchups.Matchups[i].AwayTeamID,
		// })
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(matchup)
}

func DeleteMatchup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	db := database.Connector

	var pick models.Pick
	id, _ := strconv.ParseInt(key, 10, 64)
	db.Where("id = ?", id).Delete(&pick)
	w.WriteHeader(http.StatusNoContent)
}
