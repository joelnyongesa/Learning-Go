package main
import ("fmt")

func main() {
	// Print function prints the arguments with their default format.
	var name, sex string = "Joel", "Male"

	fmt.Print(name)
	fmt.Print(sex)

	// To start a new line:
	fmt.Print(name, "\n")
	fmt.Print(sex, "\n")

	// Also possible when printing out multiple variables using one Print. The "n" here can also be a space
	fmt.Print(name, "\n", sex)

	// If neither arguments are strings, then Print() inserts a space.
	var a int = 10
	var b int = 9

	fmt.Print(a, b)


	// Println is similar to Print(), but adds a whitespace between arguments and a newline at the end
	fmt.Println("Println: ",name, sex)


	// Printf
	var greeting string = "Hello, there";
	var time int = 3;

	fmt.Printf("greeting has a value: %v and type: %T \n", greeting, greeting)
	fmt.Printf("time has a value:%v and type:%T", time, time)
}