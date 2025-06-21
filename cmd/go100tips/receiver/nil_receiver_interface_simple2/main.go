package main

import "fmt"

type myError struct{}

type MyError interface {
	Greet() string
	Error() string
	// Stringer()やFormat()があれば優先されてfmt.Printlnに呼び出される
	// Format(f fmt.State, verb rune)
}

func (e *myError) Greet() string {
	return "hello"
}

func (e *myError) Error() string {
	return "Error!!!"
}

// func (e *myError) Format(f fmt.State, verb rune) {
// 	fmt.Fprintf(f, "aaaa")
// }

func foo() MyError {
	var e *myError = nil
	return e // ここで error interface を返す(nil レシーバ)
}

func main() {
	fmt.Println(foo()) // nil.Error()が呼び出される
}
