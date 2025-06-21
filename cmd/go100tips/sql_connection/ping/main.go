package main

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
)

func main() {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mock.ExpectPing()

	if err := db.Ping(); err != nil {
		fmt.Println("Ping failed:", err)
	} else {
		fmt.Println("Ping succeeded!")
	}
}
