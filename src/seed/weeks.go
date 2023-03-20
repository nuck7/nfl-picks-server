package seed

import (
	"time"

	"github.com/nuck7/nfl-picks-server/src/models"
	"github.com/nuck7/nfl-picks-server/src/utils"
)

var Weeks = []models.Week{
	{
		Name:  "Week 1",
		Start: utils.ParseTime(time.RFC3339, "2022-07-07T16:41:46Z"),
		End:   utils.ParseTime(time.RFC3339, "2022-07-14T16:41:46Z"),
	},
	{
		Name:  "Week 2",
		Start: utils.ParseTime(time.RFC3339, "2022-07-15T16:41:46Z"),
		End:   utils.ParseTime(time.RFC3339, "2022-07-22T16:41:46Z"),
	},
}
