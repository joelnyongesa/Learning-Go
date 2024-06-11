package main

import ("fmt")

func main() {
	// This stores a sequence of characters (text). String values must be surrounded by double quotes
	var txt1 string = "Octavia"
	var txt2 string
	txt3 := "Carbon"

	fmt.Printf("Type: %T, Value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, Value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, Value: %v\n", txt3, txt3)
}