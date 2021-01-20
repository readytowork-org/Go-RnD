package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go mysql Tutorial")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Successfully connected to db")
}
