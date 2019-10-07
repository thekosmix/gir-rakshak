package main

import (
	"log"
//	"fmt"
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
	stmt, err := db.Prepare("update user set is_active=? where id=?")
 	if err != nil {
 		panic(err)
 	}
 	res, err := stmt.Exec(t.IsApproved,t.Id)
 	if err != nil {
 		panic(err)
 	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	if(rowCnt > 0) {
		return t.IsApproved	
	} else {
		return false
	}
	
}

