package main;

import ("fmt");

func main () {
	// Creating a variable using the var keyword
	var sayHello string = "This is a variable created using the var keyword.";
	fmt.Println(sayHello);

	// Creating a variable using the := sign
	inferredVar := "This is a variable created using the := sign. Its type is inferred";
	fmt.Println(inferredVar);
}