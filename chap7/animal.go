package main

import "fmt"

type Animal interface {
	Cry()
}

type Dog struct{}

func (d *Dog) Cry() { fmt.Println("わんわん") }

type Cat struct{}

func (c *Cat) Cry() { fmt.Println("にゃー") }

var animals = []Animal{&Dog{}, &Cat{}}

func Cry(a interface{}) {
	if animal, ok := a.(Animal); ok {
		animal.Cry()
	} else {
		fmt.Println("鳴かないよ")
	}
}

func main() {
	var animals = []Animal{&Dog{}, &Cat{}}
	for _, a := range animals {
		a.Cry()
	}
	Cry(&Dog{}) // わんわん
	Cry(&Cat{}) // にゃー
	Cry("book") // 鳴かないよ
}
