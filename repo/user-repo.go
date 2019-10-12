package repo

import (
	"gir-rakshak/models"
	"gir-rakshak/utils"
	"log"
	//	_ "github.com/go-sql-driver/mysql"
)

func RepoLoginUser(phoneNumber string, deviceId string, password string) (models.User, error) {
	var user models.User
	row := Db.QueryRow("select id, name, role, MD5(RAND()) from user where phone_number = ? and device_id = ? and password = ? and is_active=1", phoneNumber, deviceId, password)
	err := row.Scan(&user.Id, &user.Name, &user.Role, &user.AccessToken)
	if err != nil {
		return user, err
	}
	return user, nil
}

func RepoAllUser() []models.User {
	rows, err := Db.Query(`SELECT id, phone_number, name, role, is_active, created_date FROM user order by last_updated_date desc`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User

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

func RepoApproveUser(t models.ApproveUserRequest) (bool, error) {
	stmt, err := Db.Prepare("update user set is_active=?, last_updated_date=? where id=?")
	if err != nil {
		log.Printf(err.Error())
	}
	res, err := stmt.Exec(t.IsApproved, utils.NowAsUnixMilli(), t.Id)

	return IsDMLSuccess(res, err)
}

func RepoRegisterUser(t models.User) (bool, error) {
	stmt, err := Db.Prepare("insert into user(phone_number, name, password, role, device_id, created_date, last_updated_date) values(?,?,?,?,?,?,?)")
	if err != nil {
		log.Printf(err.Error())
	}
	res, err := stmt.Exec(t.PhoneNumber, t.Name, t.Password, "FIELD", t.DeviceId, utils.NowAsUnixMilli(), utils.NowAsUnixMilli())

	return IsDMLSuccess(res, err)
}
