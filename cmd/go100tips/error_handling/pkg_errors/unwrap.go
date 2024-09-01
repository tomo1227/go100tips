package main

import "fmt"

//lint:ignore U1000 allow unused function
func unwrap() {
	err := f()
	if err != nil {
		u, _ := err.(interface {
			Unwrap() error
		})
		fmt.Println("Unwrapped Error:", u.Unwrap())
	}
}
