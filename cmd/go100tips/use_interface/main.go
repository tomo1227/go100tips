package main

import "fmt"

type FourDimensionalPocket interface {
	getItem() (string, error)
}

type Doraemon struct {
}

func (d Doraemon) getItem() (string, error) {
	return "暗記パン🍞", nil
}

func NewDoraemon() *Doraemon {
	return &Doraemon{}
}

type Nobita struct {
	pocket FourDimensionalPocket
}

func NewNobita(pocket FourDimensionalPocket) Nobita {
	return Nobita{
		pocket: pocket,
	}
}

func (n Nobita) study() string {
	item, err := n.pocket.getItem()
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
