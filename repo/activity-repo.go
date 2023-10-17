package repo

import (
	"gir-rakshak/models"
	"gir-rakshak/utils"
	"log"
)

func RepoGetActivity(fromTime int64, toTime int64, userId int) []models.Activity {
	var selectFrom string = "SELECT id, user_id, description, recorded_time, lat, lon FROM activity "
	var where string = "where recorded_time between ? and ? "

	if userId != -1 {
		where = where + "and user_id = ? "
	}

	var order string = "order by recorded_time desc"

	var sql string = selectFrom + where + order
	rows, err := Db.Query(sql, fromTime, toTime, userId)
	if err != nil {
		log.Printf(err.Error())
	}

	defer rows.Close()

	var activities []models.Activity
	for rows.Next() {
		var act models.Activity

		err := rows.Scan(&act.Id, &act.UserId, &act.Description, &act.RecordedTime, &act.Lat, &act.Lon)
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

func RepoGetActivityDetail(activityId int64) []models.ActivityDetail {
	rows, err := Db.Query("SELECT user_id, activity_id, description, image_url, lat, lon, recorded_time FROM activity_detail where activity_id = ?", activityId)
	if err != nil {
		log.Printf(err.Error())
	}

	defer rows.Close()

	var activityDetails []models.ActivityDetail
	for rows.Next() {
		var act models.ActivityDetail
		err := rows.Scan(&act.UserId, &act.ActivityId, &act.Description, &act.ImageUrl, &act.Lat, &act.Lon, &act.RecordedTime)
		if err != nil {
			log.Fatal(err)
		}
		activityDetails = append(activityDetails, act)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	// return empty Todo if not found
	return activityDetails
}

func RepoAddActivity(activity models.Activity, userId int) (bool, error) {
	stmt, err := Db.Prepare("insert into activity(user_id, description, recorded_time, lat, lon, created_date) values(?,?,?,?,?,?)")
	if err != nil {
		log.Printf(err.Error())
	}
	res, err := stmt.Exec(userId, activity.Description, activity.RecordedTime, activity.Lat, activity.Lon, utils.NowAsUnixMilli())

	return IsDMLSuccess(res, err)
}

func RepoAddActivityDetail(activityDetail models.ActivityDetail, userId int, activityId int) (bool, error) {
	stmt, err := Db.Prepare("insert into activity_detail(user_id, activity_id, image_url, description, recorded_time, lat, lon, created_date) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Printf(err.Error())
	}
	res, err := stmt.Exec(userId, activityId, activityDetail.ImageUrl, activityDetail.Description, activityDetail.RecordedTime, activityDetail.Lat, activityDetail.Lon, utils.NowAsUnixMilli())

	return IsDMLSuccess(res, err)
}
