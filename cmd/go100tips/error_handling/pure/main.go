package main

import (
	"fmt"
)

func main() {
	err := f()
	if err != nil {
		fmt.Println(err)
	}
}

func f() error {
	err := anpanman_punch()
	if err != nil {
		return err
	}
	return nil
}

func anpanman_punch() error {
	return wetFaceError{}
}

type wetFaceError struct{}

func (b wetFaceError) Error() string {
	return "顔が濡れて元気が出ない😨"
}
