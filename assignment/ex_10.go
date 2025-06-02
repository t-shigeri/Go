package main

import "fmt"

func main() {
	var limit int
	var total int = 0
	var mod int

	fmt.Print("上限の整数を入力してください: ")
	fmt.Scan(&limit)

	// 倍数を入力
	fmt.Print("倍数を入力してください: ")
	fmt.Scan(&mod)

	for i := 1; i <= limit; i++ {
		if i%mod == 0 {
			total += i
		}
	}

	fmt.Printf("1から%dまでの%dの倍数の合計は %d です\n", limit, mod, total)
}
