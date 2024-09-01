package main

import (
	"errors"
	"fmt"
)

//lint:ignore U1000 allow unused function
func unwrap3() {
	err := f()
	if err != nil {
		fmt.Println(err)

		unwrappedErr := errors.Unwrap(err)
		if unwrappedErr != nil {
			fmt.Println("Unwrapped Error:", unwrappedErr)
		}
	}
}
