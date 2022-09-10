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

func GetTeam(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var params = r.URL.Query()
	var teams []models.Team
	json.Unmarshal(requestBody, &teams)
	db := database.Connector

	result := db.Where("")

	if params["id"] != nil {
		id, _ := strconv.ParseUint(params["id"][0], 0, 64)
		result.Where("id = ?", id)
	}

	if params["name"] != nil {
		result.Where("name = ?", params["name"][0])
	}

	if params["city"] != nil {
		result.Where("city = ?", params["city"][0])
	}

	result.Find(&teams)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(json.Marshal(teams))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var team models.Team
	var TeamInput types.CreateTeamInput
	db := database.Connector

	json.Unmarshal(requestBody, &team)

	err := json.Unmarshal(requestBody, &TeamInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&TeamInput)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(TeamInput)
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var team models.Team
	db := database.Connector

	json.Unmarshal(requestBody, &team)

	db.Create(team)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var team models.Team
	db := database.Connector

	id, _ := strconv.ParseInt(key, 10, 64)
	db.Where("id = ?", id).Delete(&team)
	w.WriteHeader(http.StatusGone)
}
