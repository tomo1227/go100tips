package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := sql.Open("mysql", "dummy:dummy@tcp(localhost:3306)/dummy")
	if err != nil {
		panic(err)
	}
	fmt.Println("sql.Open succeeded!")
}
