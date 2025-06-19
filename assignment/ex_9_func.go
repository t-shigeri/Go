package main

import "fmt"

func sumOddCalc(limit int) int {
    sum := 0
    for i := 1; i <= limit; i++ {
        if i%2 != 0 {
            sum += i
        }
    }
    return sum
}

func main() {
    var limit int
    fmt.Print("上限の整数を入力 = ")
    fmt.Scan(&limit)
    fmt.Printf("1 から %d までの奇数の合計 = %d¥n", limit, sumOddCalc(limit))
}