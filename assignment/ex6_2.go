package main

import "fmt"

func main() {
	var score int32
	var evaluation string

	fmt.Print("得点を入力してください: ")
	fmt.Scan(&score)

	switch {
	case score >= 90:
		evaluation = "Sランク"
	case score >= 80:
		evaluation = "Aランク"
	case score >= 70:
		evaluation = "Bランク"
	case score >= 50:
		evaluation = "Cランク"
	case score >= 40:
		evaluation = "Dランク"
	default:
		evaluation = "不合格"
	}

	fmt.Printf("評価: %s\n", evaluation)
}
