package main

import "fmt"

func main() {
	var bottom, height, area float64

	fmt.Print("底辺を入力してください: ")
	fmt.Scan(&bottom)

	fmt.Print("高さを入力してください: ")
	fmt.Scan(&height)

	area = bottom * height / 2

	fmt.Printf("三角形の面積は %f です\n", area)
}
