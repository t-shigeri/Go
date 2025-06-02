// ferror.go

package main

import "fmt"

func main() {
	var x32, y32 float32
	x32 = 7.0
	y32 = 0.3
	fmt.Println("x32*y32=", x32*y32)

	var x64, y64 float64
	x64 = 7.0
	y64 = 0.3
	fmt.Println("x64*y64=", x64*y64)
}
