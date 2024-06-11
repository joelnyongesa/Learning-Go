package main

import ("fmt")

func main() {
	// Integer data types -> Used to store whole numbers without decimals eg: 35, -50, etc
	// This data type has two categories, signed and unsigned integers.

	// Signed integers --> Can store both positive and negative values
	var x int = 500
	var y int =-4000

	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	// Go also has five keywords/types of signed integers: int, int8, int16, int32, and int6

	// Unsigned integers -> declared using one of uint keywords, and can only contain positive values
	var a uint = 500
	var b uint = 4500

	fmt.Printf("Type: %T, value: %v\n", a, a)
	fmt.Printf("Type: %T, value: %v\n", b, b)

	// Types of unsigned integers: uint, uint8, uint16, uint32, uint64
}