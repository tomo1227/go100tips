package main

import (
	"errors"
	"fmt"
)

type Doraemon struct {
	pocket Pocket
}

func (d Doraemon) GetNewItem() (Item, error) {
	return d.pocket.GetItem()
}

type Pocket interface {
	GetItem() (Item, error)
}

type pocket struct {
	items []Item
}

func (p *pocket) GetItem() (Item, error) {
	if len(p.items) == 0 {
		return Item{}, errors.New("何もない")
	}
	return p.items[0], nil
}

type Item struct {
	name string
}

func main() {
	// pocket 構造体のインスタンスを作成
	p := &pocket{
		items: []Item{
			{name: "タケコプター"},
			{name: "どこでもドア"},
			{name: "タイムマシン"},
		},
	}

	// Doraemon 構造体のインスタンスを作成
	d := Doraemon{
		pocket: p,
	}
	item, err := d.GetNewItem()
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("今日の秘密どうぐ : ", item)
	}
}
