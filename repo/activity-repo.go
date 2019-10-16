package repo

import (
	"gir-rakshak/models"
	"log"
)

func RepoGetActivity(fromTime int64, toTime int64) []models.Activity {
	rows, err := Db.Query("SELECT id, user_id, description, recorded_time FROM activity where recorded_time between ? and ? order by recorded_time desc", fromTime, toTime)
	if err != nil {
		log.Printf(err.Error())
	}

	defer rows.Close()

	var activities []models.Activity
	for rows.Next() {
		var act models.Activity

		err := rows.Scan(&act.Id, &act.UserId, &act.Description, &act.RecordedTime)
		if err != nil {
			log.Fatal(err)
		}
		activities = append(activities, act)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	// return empty Todo if not found
	return activities
}
