package main

import "fmt"

type Foo struct{}

func (foo *Foo) Bar() string {
	return "bar"
}

func main() {
	var foo *Foo
	fmt.Println(foo)       // nil
	fmt.Println(foo.Bar()) // bar
}
