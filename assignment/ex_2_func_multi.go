package main

import "fmt"

func arithmeticCalc(number1, number2 int) (int, int, int, float64, int) {
    add := number1 + number2
    sub := number1 - number2
    mul := number1 * number2
    div := float64(number1) / float64(number2)
    mod := number1 % number2
    return add, sub, mul, div, mod
}

func main() {
    var data1, data2 int
    var addAns, subAns, mulAns, modAns int
    var divAns float64

    fmt.Print("data1 = ")
    fmt.Scan(&data1)
    fmt.Print("data2 = ")
    fmt.Scan(&data2)

    addAns, subAns, mulAns, divAns, modAns = arithmeticCalc(data1, data2)

    fmt.Printf("%d + %d = %d¥n", data1, data2, addAns)
    fmt.Printf("%d - %d = %d¥n", data1, data2, subAns)
    fmt.Printf("%d * %d = %d¥n", data1, data2, mulAns)
    fmt.Printf("%d / %d = %f¥n", data1, data2, divAns)
    fmt.Printf("%d %% %d = %d¥n", data1, data2, modAns)
}