package main

import (
	"log"
//	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func RepoUserLocation(userId int, fromTime int64, toTime int64) []Location {

	rows, err := db.Query("SELECT latitude, longitude, recorded_time FROM location where user_id = ? and recorded_time between ? and ? order by recorded_time desc", userId, fromTime, toTime)
	if err != nil {
 		panic(err)
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


