package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type UserNoTag struct {
	ID    int
	Name  string
	Email string
}

func main() {
	user := User{ID: 1, Name: "太郎"}
	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData)) // {"id":1,"name":"太郎"}

	user2 := UserNoTag{ID: 1, Name: "太郎"}
	jsonData, _ = json.Marshal(user2)
	fmt.Println(string(jsonData)) // {"ID":1,"Name":"太郎","Email":""}
}
