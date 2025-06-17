package main

import "fmt"

type Doraemon struct{}

func NewDoraemon() Doraemon {
	return Doraemon{}
}

func (d Doraemon) getItem() (string, error) {
	return "暗記パン🍞", nil
}

type Nobita struct {
	doraemon Doraemon
}

func NewNobita(doraemon Doraemon) Nobita {
	return Nobita{
		doraemon: doraemon,
	}
}

func (n Nobita) study() string {
	item, err := n.doraemon.getItem()
	if err != nil {
		return "勉強できないよ〜😭"
	}
	return fmt.Sprintf("%sを使って勉強できた!", item)
}

func main() {
	doraemon := NewDoraemon()
	nobita := NewNobita(doraemon)
	result := nobita.study()
	fmt.Println(result)
}
