// change.go

package main

import "fmt"

func main() {
	var n int
	var x float64

	n = 123
	fmt.Println("n=", n)

	// x = n
	x = float64(n)
	fmt.Println("x=", x)
}
