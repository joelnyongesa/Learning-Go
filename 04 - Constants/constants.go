package main
import ("fmt")

// Constants can be declared both inside and outside a function. This is also an untyped constant
const NAME = "Joel Nyongesa"

// This is a typed constant
const MATERIAL string = "plastic"

// Multiple constants declaration.
const (
	GENDER string = "Male"
	GREETING string = "Good morning"
	B float32 = 3.142
)

func main() {
	// constant declared inside a function
	const PI = 3.14;

	// The values of constants are read-only, and cannot be changed. This will give us an error.
	// PI = 2;

	fmt.Println(PI)
	fmt.Println(NAME)
	fmt.Println(MATERIAL)


	fmt.Println(GENDER, GREETING, B)
}