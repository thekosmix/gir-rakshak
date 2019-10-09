package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func RepoAllUser() []User {
	rows, err := db.Query(`SELECT id, phone_number, name, role, is_active FROM user order by last_updated_date desc`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User

		err := rows.Scan(&u.Id, &u.PhoneNumber, &u.Name, &u.Role, &u.IsActive)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	// return empty Todo if not found
	return users
}

func RepoApproveUser(t UserApproved) bool {
	stmt, err := db.Prepare("update user set is_active=?, last_updated_date=? where id=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(t.IsApproved, NowAsUnixMilli(), t.Id)

	return IsDMLSuccess(res, err)
}

func RepoRegisterUser(t User) bool {
	stmt, err := db.Prepare("insert into user(phone_number, name, password, role, device_id, created_date, last_updated_date) values(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(t.PhoneNumber, t.Name, t.Password, "FIELD", t.DeviceId, NowAsUnixMilli(), NowAsUnixMilli())

	return IsDMLSuccess(res, err)
}

func IsDMLSuccess(res sql.Result, err error) bool {

	if err != nil {
		panic(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowCnt > 0
}
