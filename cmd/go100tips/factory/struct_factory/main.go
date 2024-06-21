package main

import "fmt"

// var _ Pocket = (*Doraemon)(nil)

type Pocket interface {
	getItem() (string, error)
}

type Doraemon struct {
	name string
}

func NewDoraemon() Doraemon {
	return Doraemon{name: "ドラえもん"}
}

// NOTE: おまじないが書かれていたたら、書かなくてもコンパイルエラーにならない
func (d *Doraemon) GetItem() (string, error) {
	return "タケコプター🚁", nil
}

func main() {
	d := NewDoraemon()
	fmt.Println(d.name)
	d.GetItem()
}
