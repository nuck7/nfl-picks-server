package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nuck7/nfl-picks-server/src/database"
	"github.com/nuck7/nfl-picks-server/src/models"
)

func GetPicks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var picks models.Pick
	database.Connector.First(&picks, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(picks)
}
