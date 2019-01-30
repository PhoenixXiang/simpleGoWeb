package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/PhoenixXiang/simpleGoWeb/services"
	"github.com/PhoenixXiang/simpleGoWeb/handles"
)

func main() {


	// 静态文件目录
	handles.Static("/public/", "./public")

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go?charset=utf8")
	services.SetDB(db)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		db.Close()
		fmt.Println(err)
	}
}
