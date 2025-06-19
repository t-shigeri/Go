package main

import "fmt"

func byVal(a int) {
	a = a + 1
}

func byRef(a *int) {
	*a = *a + 1
}5
func main() {
	n := 1
	byVal(n)
	fmt.Println("byValの結果:", n)

	byRef(&n)
	fmt.Println("byRefの結果:", n)
}
