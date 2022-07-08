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

func GetPick(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var params = r.URL.Query()
	var picks models.Pick
	json.Unmarshal(requestBody, &picks)

	result := database.Connector.Where("")

	if params["id"] != nil {
		id, _ := strconv.ParseUint(params["id"][0], 0, 64)
		result.Where("id = ?", id)
	}

	if params["UserID"] != nil {
		UserID, _ := strconv.ParseUint(params["user_id"][0], 0, 64)
		result.Where("user_id = ?", UserID)
	}

	if params["MatchupID"] != nil {
		MatchupID, _ := strconv.ParseUint(params["matchup_id"][0], 0, 64)
		result.Where("matchup_id = ?", MatchupID)
	}

	if params["WinnerID"] != nil {
		WinnerID, _ := strconv.ParseUint(params["winner_id"][0], 0, 64)
		result.Where("winner_id = ?", WinnerID)
	}

	result.Find(&picks)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(picks)
}

func CreatePick(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var pick models.Pick
	var PickInput types.CreatePickInput

	json.Unmarshal(requestBody, &pick)

	err := json.Unmarshal(requestBody, &PickInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.Connector.Create(&PickInput)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(PickInput)
}

func UpdatePick(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var pick models.Pick
	json.Unmarshal(requestBody, &pick)

	database.Connector.Create(pick)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pick)
}

func DeletePick(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var pick models.Pick
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&pick)
	w.WriteHeader(http.StatusNoContent)
}
