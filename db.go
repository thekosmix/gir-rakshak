package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("mysql", "root:root@(127.0.0.1:3306)/gir_rakshak?parseTime=true")
	if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
}