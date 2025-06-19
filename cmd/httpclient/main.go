package main

import (
	"fmt"
)

func main() {
	// url := "http://localhost:8080/hello"

	// headers := map[string]string{
	// 	"Content-Type": "application/json",
	// }

	// // jsonData := []byte(`{"name":"ChatGPT","language":"Go"}`)

	// resp, err := DoRequest(MethodGet, url, headers, nil)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Status Code:", resp.StatusCode())
	// fmt.Println("Response Body:", resp.String())
	method := MethodPost

	fmt.Println("HTTP Method:", method)         // 自動で String() が呼ばれる
	fmt.Println("Is GET?", method == MethodGet) // false
}
