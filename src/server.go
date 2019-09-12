package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
    "database/sql"
    "log"
    "crypto/sha1"
    "encoding/hex"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

type Users struct {
    id        int64
    phoneNumber  string
    name  string
    role  string
}

type AddUserResponse struct{ 
    Success bool
    Name  string
    PhoneNumber  string
}

func main() {
    r := mux.NewRouter()
    db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/gir_rakshak?parseTime=true")
    
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    tmpl := template.Must(template.ParseFiles("static/index.html"))

    r.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        user := Users{
            phoneNumber:   r.FormValue("phone-number"),
            name: r.FormValue("name"),
            role: r.FormValue("role"),
        }

        h := sha1.New()
        h.Write([]byte(user.phoneNumber))
        password := hex.EncodeToString(h.Sum(nil))
        currTime := time.Now().Unix()

        result, err := db.Exec(`INSERT INTO user (phone_number, name, password, role, created_date, last_updated_date) VALUES (?, ?, ?, ?, ?, ?)`, user.phoneNumber, user.name, password, user.role, currTime, currTime)
        if err != nil {
            log.Fatal(err)
        }

        id, err := result.LastInsertId()

        up := &user
        up.id = id

        aur := &AddUserResponse{Success: true, Name: user.name, PhoneNumber: user.phoneNumber}

        tmpl.Execute(w, aur)
    })



    r.HandleFunc("/users/{id}/locations/{from-date}/{to-date}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]
        fromDate := vars["from-date"]
        toDate := vars["to-date"]

        fmt.Fprintf(w, "You've requested the user's location: %s from-date %s and to-date %s\n", id, fromDate, toDate)
    }).Methods("GET")


    r.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
        
        rows, err := db.Query(`SELECT id, phone_number, name, role FROM user`)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var users []Users
        for rows.Next() {
            var u Users

            err := rows.Scan(&u.id, &u.phoneNumber, &u.name, &u.role)
            if err != nil {
                log.Fatal(err)
            }
            users = append(users, u)
        }
        if err := rows.Err(); err != nil {
            log.Fatal(err)
        }

        fmt.Fprintf(w, "%#v", users)        

//        fmt.Fprintf(w, "You've requested the user's location: %s from-date %s and to-date %s\n", id, fromDate, toDate)
    }).Methods("GET")


    r.HandleFunc("/users/{name}/{phone-number}/{role}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        name := vars["name"]
        phoneNumber := vars["phone-number"]
        role := vars["role"]
        h := sha1.New()
        h.Write([]byte(phoneNumber))
        password := hex.EncodeToString(h.Sum(nil))
        currTime := time.Now().Unix()

        result, err := db.Exec(`INSERT INTO user (phone_number, name, password, role, created_date, last_updated_date) VALUES (?, ?, ?, ?, ?, ?)`, phoneNumber, name, password, role, currTime, currTime)
        if err != nil {
            log.Fatal(err)
        }

        id, err := result.LastInsertId()

        fmt.Fprintf(w, "You've added a new user with id %s phone %s name and role %s\n", id, phoneNumber, name, role)
    }).Methods("POST")


    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    http.ListenAndServe(":80", r)
}