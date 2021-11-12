package db

import (
	"database/sql"
	"fmt"
)

func ConnectDB() *sql.DB {

	dbUser := ""
	dbPass := ""
	dbName := ""
	dbURL := ""

	var connectDB string = "" + dbUser + ":" + dbPass + "@tcp(" + dbURL + ")/" + dbName + ""

	conn, err := sql.Open("mysql", connectDB)
	if err != nil {
		fmt.Println("DB Connect Init::Err::", err)
	}

	return conn
}
