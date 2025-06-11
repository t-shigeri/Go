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

	for name, population := range prefecture {
		fmt.Println("県名：", name, " 人口：", population)
	}

	var key string
	fmt.Print("確認したい県 = ")
	fmt.Scanln(&key)

	population, exists := prefecture[key]
	if exists {
		fmt.Println(key, "は登録されています。人口は", population)
	} else {
		fmt.Println(key, "は登録されていません。")
	}
}
