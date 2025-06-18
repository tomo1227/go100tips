package main

import "fmt"

type Pocket interface {
	GetTakecopter() (string, error)
	GetDokodemodoa() (string, error)
}

type Doraemon struct {
	name string
}

func NewDoraemon() *Doraemon {
	return &Doraemon{name: "ドラえもん"}
}

func (d *Doraemon) GetTakecopter() (string, error) {
	return "タケコプター🚁", nil
}

func (d *Doraemon) GetDokodemodoa() (string, error) {
	return "どこでもドア🚪", nil
}

type Nobita struct {
	pocket Pocket
}

func NewNobita(pocket Pocket) Nobita {
	return Nobita{
		pocket: pocket,
	}
}

func (n Nobita) fly() string {
	takecopter, err := n.pocket.GetTakecopter()
	if err != nil {
		return "飛べません🐧"
	}
	return fmt.Sprintf("%sを使って飛んでいく!", takecopter)
}

func main() {
	doraemon := NewDoraemon()
	nobita := NewNobita(doraemon)
	result := nobita.fly()
	fmt.Println(result)
}
