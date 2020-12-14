package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "biz"
)

func main() {
	err := initDb()
	if err == nil {
		fmt.Println("connect successfully")
	}
}

func initDb() error {
	connStr :=
		fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr) //驱动名称为：postgres

	if err != nil {
		fmt.Println("Connected failed.", err.Error())
		return err
	}

	//ctx := context.Background()
	//err = db.PingContext(ctx)

	err = db.Ping()
	if err != nil {
		fmt.Println("ping failed.", err.Error())
	}

	// rows, err := db.Query("select * from public.test")

	// if err == nil {
	// 	rows.Close()
	// } else {
	// 	fmt.Println("query failed.", err.Error())
	// }

	return nil
}
