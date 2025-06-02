package main

import "fmt"

func main() {
	var data1, data2 int32
	var addAns, subAns, mulAns, modAns int32
	var divAns float32

	fmt.Print("1つ目の整数を入力してください: ")
	fmt.Scan(&data1)
	fmt.Print("2つ目の整数を入力してください: ")
	fmt.Scan(&data2)

	addAns = data1 + data2
	subAns = data1 - data2
	mulAns = data1 * data2
	divAns = float32(data1) / float32(data2)
	modAns = data1 % data2

	fmt.Printf("加算: %d\n", addAns)
	fmt.Printf("減算: %d\n", subAns)
	fmt.Printf("乗算: %d\n", mulAns)
	fmt.Printf("除算: %f\n", divAns)
	fmt.Printf("剰余: %d\n", modAns)
}
