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

func GetWeek(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var params = r.URL.Query()
	var weeks models.Week
	json.Unmarshal(requestBody, &weeks)

	result := database.Connector.Where("")

	if params["id"] != nil {
		id, _ := strconv.ParseUint(params["id"][0], 0, 64)
		result.Where("id = ?", id)
	}

	if params["start"] != nil {
		result.Where("start = ?", params["start"][0])
	}

	if params["end"] != nil {
		result.Where("end = ?", params["end"][0])
	}

	result.Find(&weeks)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
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

	database.Connector.Create(&WeekInput)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(WeekInput)
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
