package main

import "fmt"

type Person struct {
	name string
	age  byte
}

func main() {
	basicData := [3]Person{
		{"鈴木一朗", 38},
		{"広瀬すず", 22},
		{"土方歳三", 36},
	}

	for _, person := range basicData {
		fmt.Println("氏名：", person.name, " 年齢：", person.age)
	}
}
