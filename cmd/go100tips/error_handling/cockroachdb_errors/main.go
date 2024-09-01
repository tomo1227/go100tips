package main

import (
	"fmt"

	"github.com/cockroachdb/errors"
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
		return errors.Wrap(err, "anpanman_punch failed")
	}
	return nil
}

func anpanman_punch() error {
	return errors.New("顔が濡れて元気が出ない😨")
}
