package main

import (
	"log"
	//	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func RepoUserLocation(userId int, fromTime int64, toTime int64) []Location {

	rows, err := db.Query("SELECT latitude, longitude, recorded_time FROM location where user_id = ? and recorded_time between ? and ? order by recorded_time desc", userId, fromTime, toTime)
	if err != nil {
		log.Printf(err.Error())
	}

	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var loc Location

		err := rows.Scan(&loc.Latitude, &loc.Longitude, &loc.RecordedTime)
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

func RepoUploadUserLocation(locations []Location, userId int) (bool, error) {

	sqlStr := "INSERT INTO location(user_id, latitude, longitude, recorded_time, created_date) VALUES "
	vals := []interface{}{}

	for _, row := range locations {
		sqlStr += "(?, ?, ?, ?, ?),"
		vals = append(vals, userId, row.Latitude, row.Longitude, row.RecordedTime, NowAsUnixMilli())
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		log.Printf(err.Error())
	}

	//format all vals at once
	res, err := stmt.Exec(vals...)

	return IsDMLSuccess(res, err)
}
