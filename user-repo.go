package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func RepoGetPassword(phoneNumber string, deviceId string, password string) (User, error) {
	var user User
	row := db.QueryRow("select id, name, role, MD5(RAND()) from user where phone_number = ? and device_id = ? and password = ? and is_active=1", phoneNumber, deviceId, password)
	err := row.Scan(&user.Id, &user.Name, &user.Role, &user.AccessToken)
	if err != nil {
		return user, err
	}
	return user, nil
}

func RepoAllUser() []User {
	rows, err := db.Query(`SELECT id, phone_number, name, role, is_active, created_date FROM user order by last_updated_date desc`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User

		err := rows.Scan(&u.Id, &u.PhoneNumber, &u.Name, &u.Role, &u.IsActive, &u.CreatedDate)
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

func RepoApproveUser(t ApproveUserRequest) (bool, error) {
	stmt, err := db.Prepare("update user set is_active=?, last_updated_date=? where id=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(t.IsApproved, NowAsUnixMilli(), t.Id)

	return IsDMLSuccess(res, err)
}

func RepoRegisterUser(t User) (bool, error) {
	stmt, err := db.Prepare("insert into user(phone_number, name, password, role, device_id, created_date, last_updated_date) values(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(t.PhoneNumber, t.Name, t.Password, "FIELD", t.DeviceId, NowAsUnixMilli(), NowAsUnixMilli())

	return IsDMLSuccess(res, err)
}
