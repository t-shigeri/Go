package main

import "fmt"

// 可変長パラメータを受け取り、合計を計算する関数
func variableAddCalc(data ...int) int {
    sum := 0
    for _, value := range data {
        sum += value
    }
    return sum
}

func main() {
    fmt.Println("1 + 2 = ", variableAddCalc(1, 2))
    fmt.Println("10 + 20 + 30 = ", variableAddCalc(10, 20, 30))
}