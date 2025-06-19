package main

import "fmt"

func bmiCalc(height, weight float64) float64 {
    h := height / 100.0
    return weight / (h * h)
}

func properWeightCalc(height float64) float64 {
    h := height / 100.0
    return 22 * h * h
}

func main() {
    var height, weight float64
    fmt.Print("身長（ｃｍ単位）を入力 = ")
    fmt.Scan(&height)
    fmt.Print("体重（ｋｇ単位）を入力 = ")
    fmt.Scan(&weight)
    bmi := bmiCalc(height, weight)
    sutableWeight := properWeightCalc(height)
    fmt.Printf("BMI = %f¥n", bmi)
    fmt.Printf("適正体重 = %f¥n", sutableWeight)
}