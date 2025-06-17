package main

import (
	"encoding/json"
	"fmt"
)

func listing() error {
	b := getMessage()
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	fmt.Println("デコード結果:", m)
	// mapのanyをマーシャルするとintでもfloatになる。
	fmt.Printf("%[1]v is %[1]T\n", m["age"])
	return nil
}

func getMessage() []byte {
	return []byte(`{"name": "Taro", "age": 20}`)
}

func main() {
	listing()
}
