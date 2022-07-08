package endpoints

import (
	"encoding/json"
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
	var matchups models.Matchup
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

func CreateMatchup(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var matchup models.Matchup
	var MatchupInput types.CreateMatchupInput

	json.Unmarshal(requestBody, &matchup)

	err := json.Unmarshal(requestBody, &MatchupInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.Connector.Create(&MatchupInput)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MatchupInput)
}

func UpdateMatchup(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var matchup models.Matchup
	json.Unmarshal(requestBody, &matchup)

	database.Connector.Create(matchup)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(matchup)
}

func DeleteMatchup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var pick models.Pick
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&pick)
	w.WriteHeader(http.StatusNoContent)
}
