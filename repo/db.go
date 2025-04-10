package repo

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "freedb_rakshak_test:GUj!uH9b?kv7U4q@(sql.freedb.tech:3306)/freedb_gir_rakshak?parseTime=true")
	if err != nil {
		log.Printf(err.Error())
	}
	if err := Db.Ping(); err != nil {
		log.Printf(err.Error())
	}
}

func IsDMLSuccess(res sql.Result, err error) (bool, error) {

	if err != nil {
		log.Printf(err.Error())
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Printf(err.Error())
	}

	return rowCnt > 0, err
}

func InitConfig() {
	InitDB()
	InitCache()
}
