package main

import "fmt"

const PAI = 3.14159

func main() {
	var radius float64

	fmt.Print("円の半径を入力してください: ")
	fmt.Scan(&radius)

	circumference := 2 * PAI * radius
	area := PAI * radius * radius

	fmt.Printf("円周は %f です\n", circumference)
	fmt.Printf("面積は %f です\n", area)
}
