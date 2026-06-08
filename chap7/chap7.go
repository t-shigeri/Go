package main

import "fmt"

type Value int

func (v Value) Twice() Value {
	return v * 2
}

func (v Value) Twice() Value {
	return v * 2
}

func main() {
	var n Value = 2
	fmt.Println(n.Twice())
}
