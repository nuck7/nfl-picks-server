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
	"github.com/nuck7/nfl-picks-server/src/utils"
)

func GetWeek(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var params = r.URL.Query()
	var weeks []models.Week
	// var matchups []models.Matchup
	json.Unmarshal(requestBody, &weeks)

	weekResults := database.Connector.Where("")
	includeMatchups := utils.StringToBool(params["includeMatchups"][0]) || false
	matchupResults := database.Connector.Table("matchups")

	if params["id"] != nil {
		id, _ := strconv.ParseUint(params["id"][0], 0, 64)
		weekResults.Where("id = ?", id)
	}

	if params["start"] != nil {
		weekResults.Where("start = ?", params["start"][0])
	}

	if params["end"] != nil {
		weekResults.Where("end = ?", params["end"][0])
	}
	// var getWeekResponse []types.GetWeekResponse
	weekResults.Find(&weeks)
	if weekResults.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if params["id"] != nil && includeMatchups {
		var matchupResponse types.MatchupResponse
		id := utils.StringToUint(params["id"][0])
		// Select("matchups.id, matchups.home_team_id, matchups.away_team_id, matchups.week_id, teams.city, teams.name, teams.id")
		matchupResults.Select("matchups.id, matchups.week_id, at.id as AwayTeamID, CONCAT(at.city, ' ', at.name) as AwayTeamName, ht.id as HomeTeamID, CONCAT(ht.city, ' ', ht.name) as HomeTeamName")
		matchupResults.Joins("join teams at on at.id = matchups.away_team_id").Joins("join teams ht on ht.id = matchups.home_team_id").Where("matchups.week_id = ?", id)
		matchupResults.Find(&matchupResponse)
		fmt.Printf("matchupResponse", matchupResponse)
		if matchupResults.Error != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		for i := 0; i < len(weeks); i++ {
			if weeks[i].ID == uint(id) {
				// weeks[i].Matchups = matchupResponse
				// getWeekResponse[i] = weeks[i]
				// getWeekResponse[i].Matchups = matchupResponse
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weeks)
}

func CreateWeek(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var week models.Week
	var WeekInput types.CreateWeekInput

	json.Unmarshal(requestBody, &week)

	err := json.Unmarshal(requestBody, &WeekInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.Connector.Create(&week)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(week)
}

func UpdateWeek(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var week models.Week
	json.Unmarshal(requestBody, &week)

	database.Connector.Create(week)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(week)
}

func DeleteWeek(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var week models.Week
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&week)
	w.WriteHeader(http.StatusNoContent)
}
