package main

import (
	"fmt"
)

func main() {
	err := f()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func f() error {
	err := anpanman_punch()
	if err != nil {
		return fmt.Errorf("anpanman_punch failed: %v", err)
	}
	return nil
}

func anpanman_punch() error {
	return wetFaceError{}
}

type wetFaceError struct{}

func (w wetFaceError) Error() string {
	return "顔が濡れて元気が出ない😨"
}
