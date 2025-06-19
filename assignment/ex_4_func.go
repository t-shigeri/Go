package main

import "fmt"

const PAI float64 = 3.14159

func circumCalc(radius float64) float64 {
    return 2 * PAI * radius
}

func areaCircleCalc(radius float64) float64 {
    return PAI * radius * radius
}

func main() {
    var radius float64
    fmt.Print("半径を入力 = ")
    fmt.Scan(&radius)
    circumference := circumCalc(radius)
    area := areaCircleCalc(radius)
    fmt.Printf("円周 = %f¥n", circumference)
    fmt.Printf("面積 = %f¥n", area)
}