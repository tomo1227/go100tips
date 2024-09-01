package main

import (
	"fmt"
)

//lint:ignore U1000 allow unused function
func unwrap1() {
	err := f()
	if err != nil {
		fmt.Println(err)

		if anpanmanErr, ok := err.(AnpanmanPunchError); ok {
			fmt.Println("Unwrapped Error:", anpanmanErr.Err)
		}
	}
}
