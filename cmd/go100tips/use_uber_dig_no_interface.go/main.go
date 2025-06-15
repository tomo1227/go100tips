package main

import (
	"fmt"

	"go.uber.org/dig"
)

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

func NewDoraemon() FourDimensionalPocket {
	return Doraemon{}
}

type Nobita struct {
	pocket FourDimensionalPocket
}

func NewNobita(pocket FourDimensionalPocket) Nobita {
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
	container := dig.New()

	container.Provide(NewDoraemon)
	container.Provide(NewNobita)

	err := container.Invoke(func(b Nobita) {
		fmt.Println(b.fly())
	})

	if err != nil {
		panic(err)
	}
}
