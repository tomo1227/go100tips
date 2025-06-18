package main

import (
	"fmt"

	"github.com/samber/do"
)

type FourDimensionalPocket interface {
	GetTakecopter() (string, error)
	GetDokodemodoa() (string, error)
}

type Bird interface {
	fly() string
}

type Doraemon struct{}

func (d Doraemon) GetTakecopter() (string, error) {
	return "タケコプター🚁", nil
}

func (d Doraemon) GetDokodemodoa() (string, error) {
	return "どこでもドア🚪", nil
}

func NewDoraemon(i *do.Injector) (FourDimensionalPocket, error) {
	return Doraemon{}, nil
}

type Nobita struct {
	pocket FourDimensionalPocket
}

func (n Nobita) fly() string {
	takecopter, err := n.pocket.GetTakecopter()
	if err != nil {
		return "飛べません🐧"
	}
	return fmt.Sprintf("%sを使って飛んでいく!", takecopter)
}

func NewNobita(i *do.Injector) (Bird, error) {
	pocket := do.MustInvoke[FourDimensionalPocket](i)
	return Nobita{pocket: pocket}, nil
}

func main() {
	injector := do.New()

	do.Provide(injector, NewDoraemon)
	do.Provide(injector, NewNobita)

	nobita := do.MustInvoke[Bird](injector)

	fmt.Println(nobita.fly())
}
