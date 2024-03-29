package repo

import (
	"gir-rakshak/models"
	"gir-rakshak/utils"
	"log"
)

func RepoUserLocation(userId int, fromTime int64, toTime int64) []models.Location {
	rows, err := Db.Query("SELECT lat, lon, recorded_time FROM location where user_id = ? and recorded_time between ? and ? order by recorded_time desc", userId, fromTime, toTime)
	if err != nil {
		log.Printf(err.Error())
	}

	defer rows.Close()

	var locations []models.Location
	for rows.Next() {
		var loc models.Location

		err := rows.Scan(&loc.Lat, &loc.Lon, &loc.RecordedTime)
		if err != nil {
			log.Fatal(err)
		}
		locations = append(locations, loc)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	// return empty Todo if not found
	return locations
}

func RepoUploadUserLocation(locations []models.Location, userId int) (bool, error) {

	sqlStr := "INSERT INTO location(user_id, lat, lon, recorded_time, created_date) VALUES "
	vals := []interface{}{}

	for _, row := range locations {
		sqlStr += "(?, ?, ?, ?, ?),"
		vals = append(vals, userId, row.Lat, row.Lon, row.RecordedTime, utils.NowAsUnixMilli())
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	stmt, err := Db.Prepare(sqlStr)

	if err != nil {
		log.Printf(err.Error())
	}

	//format all vals at once
	res, err := stmt.Exec(vals...)

	return IsDMLSuccess(res, err)
}
