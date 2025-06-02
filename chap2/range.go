// range.go

package main

import "fmt"

func main() {
	var x, y int
	x = 3
	{
		var x = 12
		fmt.Println(x)
		var y = 11
		fmt.Println(y)
	}
	y = x
	fmt.Println(y)
}
