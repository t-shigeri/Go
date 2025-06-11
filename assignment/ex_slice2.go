package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var total int = 0
	var evenNumber int = 0
	var oddNumber int = 0

	var upper, lower int

	fmt.Println("スライス前 list = ", list)
	fmt.Print("下限値,上限値 = ")
	fmt.Scanf("%d,%d", &lower, &upper)

	cutOut := list[lower : upper+1]

	for _, value := range cutOut {
		total += value
		if value%2 == 0 {
			evenNumber++
		} else {
			oddNumber++
		}
	}

	fmt.Println("list = ", list)
	fmt.Println("cutOut = ", cutOut)
	fmt.Println("合計 = ", total)
	fmt.Println("偶数値の個数 = ", evenNumber)
	fmt.Println("奇数値の個数 = ", oddNumber)
}
