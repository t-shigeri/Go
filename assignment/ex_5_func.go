package main

import "fmt"

func evenJudg(number int) bool {
    return number%2 == 0
}

func main() {
    var data int
    fmt.Print("整数値を入力 = ")
    fmt.Scan(&data)
    if evenJudg(data) {
        fmt.Printf("%d は偶数¥n", data)
    } else {
        fmt.Printf("%d は奇数¥n", data)
    }
}