package main

import "fmt"

func main() {
	var data int32

	fmt.Print("整数を入力してください: ")
	fmt.Scan(&data)

	if data%2 == 0 {
		fmt.Println("偶数です")
	} else {
		fmt.Println("奇数です")
	}
}
