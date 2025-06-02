// omit.go

package main

import "fmt"

func swap(a int, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

func main() {
	var a = 5
	var b = 10

	a, b = swap(a, b)
	fmt.Println(a, b)
}
