package main

import (
	"errors"
	"fmt"
)

type Item struct {
	name string
}

type Doraemon struct {
	pocket Pocket
}

func (d Doraemon) GetNewItem() (Item, error) {
	return d.pocket.GetItem()
}

type Pocket struct {
	items []Item
}

func (p *Pocket) GetItem() (Item, error) {
	if len(p.items) == 0 {
		return Item{}, errors.New("何もない")
	}
	return p.items[0], nil
}

func main() {
	d := Doraemon{
		pocket: Pocket{
			items: []Item{
				{name: "タケコプター"}, {name: "どこでもドア"}, {name: "タイムマシン"},
			},
		},
	}
	item, err := d.GetNewItem()
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("今日の秘密どうぐ : ", item)
	}
}
