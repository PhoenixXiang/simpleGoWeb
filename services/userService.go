package services

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	Id       int
	Username string
	Password string
	Created  string
}
type UserService struct {
	DB *sql.DB
}

var u *UserService
var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
	u = &UserService{
		DB: db,
	}
}

func QueryUser(username string) *User {
	row := u.DB.QueryRow("SELECT uid, password, created FROM userinfo where username = ?", username)
	var uid int
	var password string
	var created string
	err := row.Scan(&uid, &password, &created)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	u := &User{
		Id:       uid,
		Username: username,
		Password: password,
		Created:  created,
	}
	return u
}

func CreateUser(username, password string) int {
	stmt, err := u.DB.Prepare("INSERT userinfo SET username=?,password=?,created=?")
	if err != nil {
		fmt.Println(err)
		return 0
	}

	res, err := stmt.Exec(username, password, time.Now())
	if err != nil {
		fmt.Println(err)
		return 0
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return int(id)
}
