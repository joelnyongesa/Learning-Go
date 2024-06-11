// They are used with the Printf() function
package main

import ("fmt")

func main() {
	/*
		%v -> Prints the value in default format
		%#v -> prints the value in Go-syntax format
		%T -> Prints the type of the value
		%% -> Prints the % sign
	*/

	name := "Joel Nyongesa"
	age := 27

	fmt.Printf("My name is %v and I am %v years old\n", name, age)

	fmt.Printf("The type of %v is %T\n", name, name)
	fmt.Printf("The type of %#v is %T\n", name, name)
	
}