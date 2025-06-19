package main

import "fmt"

func obesityInfo(height, weight float64) (float64, float64) {
    h := height / 100.0
    bmi := weight / (h * h)
    suitableWeight := 22 * h * h
    return bmi, suitableWeight
}

func main() {
    var height, weight float64
    fmt.Print("身長（ｃｍ単位）を入力 = ")
    fmt.Scan(&height)
    fmt.Print("体重（ｋｇ単位）を入力 = ")
    fmt.Scan(&weight)
    bmi, sutableWeight := obesityInfo(height, weight)
    fmt.Printf("BMI = %f¥n", bmi)
    fmt.Printf("適正体重 = %f¥n", sutableWeight)
}