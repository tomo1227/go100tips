package main

import "fmt"

type I interface {
	Do(num int) error
}

type impl struct {
}

type a struct {
	i I
}

func (a a) Do(num int) error {
	fmt.Println(num)
	return nil
}

func (a impl) Do(num int) error {
	fmt.Println(num * 2)
	return nil
}

func main() {
	a := a{
		i: impl{},
	}
	// a.Do(1)
	a.i.Do(1)
}
