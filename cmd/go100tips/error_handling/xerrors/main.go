package main

import (
	"fmt"

	"golang.org/x/xerrors"
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
		return xerrors.Errorf("anpanman_punch failed: %w", err)
	}
	return nil
}

func anpanman_punch() error {
	return xerrors.New("顔が濡れて元気が出ない😨")
}
