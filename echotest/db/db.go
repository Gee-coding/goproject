package db

import (
	"database/sql"
	"fmt"
)

func ConnectDB() *sql.DB {

	conn, err := sql.Open("mysql", "testball:testball@tcp(localhost:25060)/test_ball")
	if err != nil {
		fmt.Println("DB Connect Init::Err::", err)
		panic(err.Error())
	}

	return conn
}
