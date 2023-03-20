package endpoints

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/nuck7/nfl-picks-server/src/database"
// 	"github.com/nuck7/nfl-picks-server/src/types"
// )

// func GetPicks(w http.ResponseWriter, r *http.Request) {
// 	var params = r.URL.Query()
// 	var pickView []types.PickView
// 	var picks []types.Pick
// 	var picksResponse types.PickResponse
// 	db := database.Connector

// 	pickViewResults := db.Table("pickView").Where("")

// 	// weekId, _ := strconv.ParseUint(params["id"][0], 0, 64)
// 	id, _ := strconv.ParseUint(params["id"][0], 0, 64)

// 	pickViewResults.Where("week_id = ?", id)

// 	pickViewResults.Find(&pickView)
// 	if pickViewResults.Error != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	// fmt.Println("picks", picks)
// 	// fmt.Println("picks row count", pickViewResults.RowsAffected)

// 	for i := 0; i < len(picks); i++ {
// 		picks = append(picks, types.Pick{
// 			MatchupID:    picks[i].Matchup_id,
// 			HomeTeamID:   picks[i].Home_team_id,
// 			HomeTeamCity: picks[i].Home_team_city,
// 			HomeTeamName: picks[i].Home_team_name,
// 			AwayTeamID:   picks[i].Away_team_id,
// 			AwayTeamCity: picks[i].Away_team_city,
// 			AwayTeamName: picks[i].Away_team_name,
// 			WeekID:       picks[i].Week_id,
// 		})
// 	}

// 	picksResponse.ID = picks[0].Week_id
// 	picksResponse.Name = picks[0].Week_name
// 	picksResponse.Start = picks[0].Week_start
// 	picksResponse.End = picks[0].Week_end
// 	picksResponse.Picks = picks

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(PickResponse)
// }
