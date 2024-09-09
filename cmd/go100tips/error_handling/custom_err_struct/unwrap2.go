package main

import (
	"fmt"
)

//lint:ignore U1000 allow unused function
func unwrap2() {
	err := f()
	if err != nil {
		fmt.Println(err)

		u, _ := err.(interface {
			Unwrap() error
		})
		fmt.Println("Unwrapped Error:", u.Unwrap())
	}
}