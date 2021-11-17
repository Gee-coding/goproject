package db

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
)

func ConnectDB() *sql.DB {

	conn, err := sql.Open("mysql", "testball:testball@tcp(database-cluster-do-user-8554052-0.b.db.ondigitalocean.com:25060)/test_ball")
	if err != nil {
		fmt.Println("DB Connect Init::Err::", err)
		panic(err.Error())
	}

	return conn
}
