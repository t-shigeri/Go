package main

import "fmt"

func main() {
	prefecture := map[string]float64{
		"青森県": 124.9,
		"岩手県": 122.9,
		"秋田県": 96.6,
		"宮城県": 230.6,
		"山形県": 108.0,
		"福島県": 184.8,
	}

	fmt.Println("削除前")
	for name, population := range prefecture {
		fmt.Println("県名：", name, " 人口：", population)
	}

	var key string
	fmt.Print("削除したい県 = ")
	fmt.Scanln(&key)
	delete(prefecture, key)

	fmt.Println("削除後")
	for name, population := range prefecture {
		fmt.Println("県名：", name, " 人口：", population)
	}
}
