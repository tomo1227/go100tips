package main

import (
	"fmt"
)

// Pocketインターフェース
type Pocket interface {
	DoSomething() string
}

// DoraemonはServiceインターフェースを実装する
type Doraemon struct {
	Name string
}

// DoSomethingはMyServiceのメソッド
func (s Doraemon) DoSomething() string {
	return "Hello, " + s.Name
}

// NobitaはServiceを依存性として持つ
type Nobita struct {
	service Pocket
}

// NewClientはClientのコンストラクタで、Serviceを注入する
func NewClient(s Pocket) *Nobita {
	return &Nobita{service: s}
}

// UseServiceはServiceのメソッドを呼び出す
func (c *Nobita) UseService() {
	fmt.Println(c.service.DoSomething())
}

func main() {
	// MyServiceのインスタンスを作成
	myService := Doraemon{Name: "World"}

	// MyServiceのインスタンスをClientに注入
	client := NewClient(myService)

	// ClientがServiceを使用
	client.UseService()
}
