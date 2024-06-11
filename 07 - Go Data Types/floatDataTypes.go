package main

import ("fmt")

func main() {
	// The float data type has two keywords: float32 and float64. The default value is float64
	var x float32 = 123.78
	var y float32 = 3.4e+38
	var z float64 = 1.7e+300

	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)
	fmt.Printf("Type: %T, value: %v\n", z, z)
}