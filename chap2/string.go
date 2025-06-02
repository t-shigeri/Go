// string.go

package main

import (
	"fmt"
)

func main() {
	var a, b string
	a = "ABCdef"
	fmt.Println(string(a[0]))

	b = "123456"
	fmt.Println(string(b[1:3]))
}
