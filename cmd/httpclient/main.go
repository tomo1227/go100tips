package main

import (
	"fmt"
)

func main() {
	url := "http://localhost:8080/hello"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := DoRequest(MethodGet, url, headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Response Body:", resp.String())
}
