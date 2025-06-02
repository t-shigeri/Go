package main

import "fmt"

func main() {
	var height, weight float64
	var heightM, bmi, suitableWeight float64

	fmt.Print("身長(cm)を入力してください: ")
	fmt.Scan(&height)

	fmt.Print("体重(kg)を入力してください: ")
	fmt.Scan(&weight)

	heightM = height / 100

	bmi = weight / (heightM * heightM)
	suitableWeight = 22 * (heightM * heightM)

	fmt.Printf("BMI値は %f です\n", bmi)
	fmt.Printf("適正体重は %f kgです\n", suitableWeight)
}
