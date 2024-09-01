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
		return AnpanmanPunchError{Err: err}
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

type AnpanmanPunchError struct {
	Err error
}

func (a AnpanmanPunchError) Error() string {
	return "anpanman_punch failed: " + a.Err.Error()
}

func (a AnpanmanPunchError) Unwrap() error {
	return a.Err
}
