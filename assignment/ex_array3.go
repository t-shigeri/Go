package main

import (
	"fmt"
)

func main() {
	var list [5]int
	var total, evenNumber, oddNumber int

	fmt.Println("5つの整数値を入力してください：")
	for i := 0; i < 5; i++ {
		fmt.Printf("%dつ目：", i+1)
		fmt.Scan(&list[i])
	}

	for _, value := range list {
		total += value
		if value%2 == 0 {
			evenNumber++
		} else {
			oddNumber++
		}
	}

	fmt.Println("合計値：", total)
	fmt.Println("偶数の数：", evenNumber)
	fmt.Println("奇数の数：", oddNumber)
}
