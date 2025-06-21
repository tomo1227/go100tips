package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string {
	return "error occurred"
}

func foo() error {
	var e *MyError = nil
	return e // ここで error interface を返す(nil レシーバ)
}

func main() {
	fmt.Println(foo()) // nil.Error()が呼び出される
}
