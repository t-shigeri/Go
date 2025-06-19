// mathpkg.go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("数値を入力してください: ")
	fmt.Scan(&x)

	// 各種の結果を表示
	fmt.Printf("%.2fを切り上げた結果は: %.2f\n", x, math.Ceil(x))
	fmt.Printf("%.2fを切り下げた結果は: %.2f\n", x, math.Floor(x))
	fmt.Printf("%.2fを四捨五入した結果は: %.2f\n", x, math.Round(x))
	fmt.Printf("%.2fのe乗は: %.2f\n", x, math.Exp(x))
	fmt.Printf("%.2fの平方根は: %.2f\n", x, math.Sqrt(x))
}
