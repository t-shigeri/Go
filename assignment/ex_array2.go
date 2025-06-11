package main

import "fmt"

func main() {

	var list = [5]int{1, 2, 3, 4, 5}

	var total int = 0
	var evenNumber int = 0
	var oddNumber int = 0

	for _, value := range list {
		total += value
		if value%2 == 0 {
			evenNumber++
		} else {
			oddNumber++
		}
	}

	fmt.Println("list = ", list)
	fmt.Println("リストの合計 = ", total)
	fmt.Println("偶数値の個数 = ", evenNumber)
	fmt.Println("奇数値の個数 = ", oddNumber)
}
