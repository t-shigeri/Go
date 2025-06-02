package main

import "fmt"

func main() {
	var limit int
	var total int = 0

	fmt.Print("上限の整数を入力してください: ")
	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		if i%2 != 0 {
			total += i
		}
	}

	fmt.Printf("1から%dまでの奇数の合計は %d です\n", limit, total)
}
