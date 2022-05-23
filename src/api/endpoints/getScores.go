package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nuck7/nfl-picks-server/src/database"
	"github.com/nuck7/nfl-picks-server/src/models"
)

func GetScores(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var team models.Team
	json.Unmarshal(requestBody, &team)

	database.Connector.Create(team)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}
