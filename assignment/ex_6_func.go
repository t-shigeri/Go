package main

import "fmt"

func scoreEval(score int) string {
    switch {
    case score >= 90:
        return "Sランク"
    case score >= 80:
        return "Aランク"
    case score >= 70:
        return "Bランク"
    case score >= 50:
        return "Cランク"
    case score >= 40:
        return "Dランク"
    default:
        return "不合格"
    }
}

func main() {
    var score int
    fmt.Print("得点を入力 = ")
    fmt.Scan(&score)
    fmt.Printf("成績の評価は、%s です。¥n", scoreEval(score))
}