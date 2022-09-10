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

func GetUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var params = r.URL.Query()
	var users []models.User
	db := database.Connector

	json.Unmarshal(requestBody, &users)

	result := db.Where("")

	if params["id"] != nil {
		id, _ := strconv.ParseUint(params["id"][0], 0, 64)
		result.Where("id = ?", id)
	}
	result.Find(&users)

	if params["name"] != nil {
		result.Where("name = ?", params["name"][0])
	}

	if params["email"] != nil {
		result.Where("email = ?", params["email"][0])
	}

	result.Find(&users)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user models.User
	var UserInput types.CreateUserInput
	db := database.Connector

	json.Unmarshal(requestBody, &user)

	err := json.Unmarshal(requestBody, &UserInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user models.User
	db := database.Connector
	json.Unmarshal(requestBody, &user)

	db.Create(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var user models.User
	db := database.Connector

	id, _ := strconv.ParseInt(key, 10, 64)
	db.Where("id = ?", id).Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
