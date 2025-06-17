package main

import "fmt"

type FourDimensionalPocket interface {
	GetTakecopter() (string, error)
	GetDokodemodoa() (string, error)
}
type Doraemon struct{}

func (d Doraemon) GetTakecopter() (string, error) {
	return "タケコプター🚁", nil
}

func (d Doraemon) GetDokodemodoa() (string, error) {
	return "どこでもドア🚪", nil
}
func NewDoraemon() Doraemon {
	return Doraemon{}
}

type Nobita struct {
	pocket FourDimensionalPocket
}

type Bird interface {
	Fly() string
}

func NewNobita(pocket FourDimensionalPocket) Bird {
	return Nobita{
		pocket: pocket,
	}
}

func (n Nobita) Fly() string {
	takecopter, err := n.pocket.GetTakecopter()
	if err != nil {
		return "飛べません🐧"
	}
	return fmt.Sprintf("%sを使って飛んでいく!", takecopter)
}

func main() {
	doraemon := NewDoraemon()
	nobita := NewNobita(doraemon)
	result := nobita.Fly()
	fmt.Println(result)
}
