package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/nuck7/nfl-picks-server/src/database"
	"github.com/nuck7/nfl-picks-server/src/types"
)

func GetWeekMatchups(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var params = r.URL.Query()
	var weekMatchups []types.WeekMatchupView
	var weekMatchupResponse types.WeekMatchupResponse
	var matchups []types.MatchupResponse
	db := database.Connector

	json.Unmarshal(requestBody, &weekMatchups)

	weekMatchupViewResults := db.Table("weekMatchupView").Where("")

	// if params["id"] != nil {
	id, _ := strconv.ParseUint(params["id"][0], 0, 64)
	weekMatchupViewResults.Where("week_id = ?", id)
	// }

	weekMatchupViewResults.Find(&weekMatchups)
	if weekMatchupViewResults.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// fmt.Println("weekMatchup", weekMatchups)
	// fmt.Println("weekMatchup row count", weekMatchupViewResults.RowsAffected)

	for i := 0; i < len(weekMatchups); i++ {
		matchups = append(matchups, types.MatchupResponse{
			ID:           weekMatchups[i].Matchup_id,
			HomeTeamID:   weekMatchups[i].Home_team_id,
			HomeTeamCity: weekMatchups[i].Home_team_city,
			HomeTeamName: weekMatchups[i].Home_team_name,
			AwayTeamID:   weekMatchups[i].Away_team_id,
			AwayTeamCity: weekMatchups[i].Away_team_city,
			AwayTeamName: weekMatchups[i].Away_team_name,
			WeekID:       weekMatchups[i].Week_id,
		})
	}

	weekMatchupResponse.ID = weekMatchups[0].Week_id
	weekMatchupResponse.Name = weekMatchups[0].Week_name
	weekMatchupResponse.Start = weekMatchups[0].Week_start
	weekMatchupResponse.End = weekMatchups[0].Week_end
	weekMatchupResponse.Matchups = matchups

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weekMatchupResponse)
}
