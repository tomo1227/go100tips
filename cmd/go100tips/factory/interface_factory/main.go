package main

type Pocket interface {
	getItem() (string, error)
}

type Doraemon struct {
	name string
}

func NewDoraemon() Pocket {
	return &Doraemon{name: "ドラえもん"}
}

func (d *Doraemon) getItem() (string, error) {
	return "タケコプター🚁", nil
}

func main() {
	d := NewDoraemon()
	d.getItem()
	// fmt.Printf(d.name)
}
