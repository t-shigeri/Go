// iota.go

package main

import "fmt"

func main() {
	const (
		Sun int = iota
		Mon
		Tue
		Wed
		Thu
		Fri
		Sat
	)
	fmt.Println(Sat)
}
