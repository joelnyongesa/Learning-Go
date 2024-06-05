package main;

import ("fmt");

// Declaring a variable using the var keyword means we can declare our variable even outside a function.
var age int = 28;

func main() {
	var a string;
	var b int;
	var c bool;


	// It is possible to assign a value to a variable after it is declared
	a = "Joel Nyongesa";
	b = 100;
	c = true;


	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	fmt.Println(age)
}