package main

import "fmt"

func main() {
	prefecture := make(map[string]float64)

	prefecture["青森県"] = 124.9 // (2)
	prefecture["岩手県"] = 122.9 // (3)
	prefecture["秋田県"] = 96.6  // (4)
	prefecture["宮城県"] = 230.6 // (5)
	prefecture["山形県"] = 108.0 // (6)
	prefecture["福島県"] = 184.4 // (7)

	for name, population := range prefecture {
		fmt.Println("県名：", name, " 人口：", population)
	}
}
